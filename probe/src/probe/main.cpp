#include <cstdio>
#include <iostream>
#include <cstdlib>
#include "sinsp.h"
#include "src/probe/utils/termination_handler.h"
#include <unistd.h>
#include <sys/un.h>
#include "src/probe/converter/sysdig_converter.h"
#include "src/probe/publisher/publisher.h"
#include "src/probe/converter/kindling_event.pb.h"
#include "driver/driver_config.h"
#include "src/stirling/stirling.h"
#include "src/common/base/base.h"
#include <src/stirling/stirling.h>
#include "src/probe/converter/uprobe_converter.h"
#include <src/stirling/utils/linux_headers.h>
#include <src/common/metrics/metrics.h>

#include "src/probe/catch_sig.h"

#include <prometheus/exposer.h>
#include <prometheus/registry.h>

DEFINE_bool(enable_stirling, false, "If true, pixie stirling module is enabled.");
DEFINE_int32(port, 9112, "The port to export prometheus metrics.");

#define KINDLING_PROBE_VERSION "v0.1-2021-1221"
void do_inspect(sinsp *inspector, sinsp_evt_formatter *formatter, int pid, int syscallOn, publisher* pub) {
    int32_t res;
    sinsp_evt *ev;
    string line;
    int filter_out_pid_event = -1;
    int is_syscall_out = 0;
    converter *sysdigConverter = new sysdig_converter(inspector);
    while (true) {
        res = inspector->next(&ev);
        if (res == SCAP_TIMEOUT) {
            continue;
        } else if (res != SCAP_SUCCESS) {
            cerr << "res = " << res << endl;
            break;
        }

        if (!inspector->is_debug_enabled() &&
            ev->get_category() & EC_INTERNAL) {
            continue;
        }
        if (ev->get_thread_info() == nullptr) {
            continue;
        }
        if (ev->get_thread_info()->m_ptid == (__int64_t) pid || ev->get_thread_info()->m_pid == (__int64_t) pid) {
            continue;
        }

        if (ev->get_thread_info()->m_comm == "sshd" || ev->get_type() == PPME_SCHEDSWITCH_6_E || ev->get_type() == PPME_SCHEDSWITCH_6_X || ev->get_thread_info()->m_comm == "containerd" || ev->get_thread_info()->m_comm == "dockerd" || ev->get_thread_info()->m_comm == "container-shim") {
            continue;
        }
        pub->consume_sysdig_event(ev, pid, sysdigConverter);
        if (is_syscall_out == 1 && filter_out_pid_event == ev->get_thread_info()->m_pid && formatter->tostring(ev, &line)) {
            cout<< line << endl;
        }

//        if (formatter->tostring(ev, &line)) {
//            cout << line << endl;
//        }
    }
}

void get_capture_statistics(sinsp* inspector) {
	scap_stats s;
	while(1) {
		inspector->get_capture_stats(&s);
		printf("seen by driver: %" PRIu64 "\n", s.n_evts);
		if(s.n_drops != 0){
            printf("Number of dropped events: %" PRIu64 "\n", s.n_drops);
		}
		if(s.n_drops_buffer != 0){
            printf("Number of dropped events caused by full buffer: %" PRIu64 "\n", s.n_drops_buffer);
		}
		if(s.n_drops_pf != 0){
            printf("Number of dropped events caused by invalid memory access: %" PRIu64 "\n", s.n_drops_pf);
		}
		if(s.n_drops_bug != 0){
            printf("Number of dropped events caused by an invalid condition in the kernel instrumentation: %" PRIu64 "\n", s.n_drops_bug);
		}
		if(s.n_preemptions != 0){
            printf("Number of preemptions: %" PRIu64 "\n", s.n_preemptions);
		}
		if(s.n_suppressed != 0){
            printf("Number of events skipped due to the tid being in a set of suppressed tids: %" PRIu64 "\n", s.n_suppressed);
		}
		if(s.n_tids_suppressed != 0){
            printf("Number of threads currently being suppressed: %" PRIu64 "\n", s.n_tids_suppressed);
		}
		sleep(10);
	}
}

int main(int argc, char** argv) {
    px::EnvironmentGuard env_guard(&argc, argv);
//  absl::flat_hash_set<std::string_view>& source_names = ;
//  auto sr = px::stirling::CreateProdSourceRegistry();
// init prometheus
    // create an http server running on port 8080
    LOG(INFO) << "init prometheus ... ";
    prometheus::Exposer exposer{"0.0.0.0:" + std::to_string(FLAGS_port)};
    std::shared_ptr s_registry = std::shared_ptr<prometheus::Registry>(&(GetMetricsRegistry()));
    exposer.RegisterCollectable(s_registry);
    LOG(INFO) << "metrics registry succesfully registerd!";


    int pid = getpid();
    int is_syscall_on = 0;
    sinsp *inspector = nullptr;
    bool bpf = false;
    string bpf_probe;
    string output_format;
    output_format = "*%evt.num %evt.outputtime %evt.cpu %container.name (%container.id) %proc.name (%thread.tid:%thread.vtid) %evt.dir %evt.type %evt.info";

    cout << "Start kindling probe...\n";
    cout << "KINDLING_PROBE_VERSION: " << KINDLING_PROBE_VERSION << endl;
    TerminationHandler::InstallSignalHandlers();
    try {
        inspector = new sinsp();
        inspector->set_hostname_and_port_resolution_mode(false);
        sinsp_evt_formatter formatter(inspector, output_format);
        inspector->set_snaplen(80);

        const char *probe = scap_get_bpf_probe_from_env();
        if (probe) {
            bpf = true;
            bpf_probe = probe;
        }

        bool open_success = true;

        try {
            inspector->open("");
            inspector->clear_eventmask();
        }
        catch (const sinsp_exception &e) {
            open_success = false;
        }

        //
        // Starting the live capture failed, try to load the driver with
        // modprobe.
        //
        if (!open_success) {
            if (bpf) {
                if (bpf_probe.empty()) {
                    fprintf(stderr, "Unable to locate the BPF probe\n");
                }
            } else {
                if (system("modprobe " PROBE_NAME " > /dev/null 2> /dev/null")) {
                    fprintf(stderr, "Unable to load the driver\n");
                }
            }

            inspector->open("");
        }
        thread catch_signal(sigsetup);
		thread stat(get_capture_statistics, inspector);

		uprobe_converter* uconv = new uprobe_converter();
      publisher *pub = new publisher(inspector, uconv);

		auto kernel_version = px::stirling::utils::GetKernelVersion().ValueOrDie();
        LOG(INFO) << absl::Substitute("kernel version is $0.$1.$2", kernel_version.version, kernel_version.major_rev, kernel_version.minor_rev);
	std::unique_ptr<px::stirling::Stirling> stirling_;
        if (!FLAGS_enable_stirling) {
            LOG(WARNING) << "stirling module is set to disable, add --enable_stirling to enable ... ";
        } else {
            // check kernel version
            bool init_stirling = true;
            if ((kernel_version.version == 4 && kernel_version.major_rev < 14) || kernel_version.version < 4) {
                init_stirling = false;
                LOG(WARNING) << absl::Substitute("kernel version is $0.$1.$2, do not init stirling ... ", kernel_version.version, kernel_version.major_rev, kernel_version.minor_rev);
            }

            if (init_stirling) {
                // init bcc & stirling
                LOG(INFO) << "begin to init stirling ...";
                auto stirling = px::stirling::Stirling::Create(px::stirling::CreateSourceRegistry(px::stirling::GetSourceNamesForGroup(px::stirling::SourceConnectorGroup::kTracers))
                                                                       .ConsumeValueOrDie());
                stirling->RegisterDataPushCallback(std::bind(&publisher::consume_uprobe_data, pub,
                                                             std::placeholders::_1, std::placeholders::_2,
                                                             std::placeholders::_3));
                TerminationHandler::set_stirling(stirling.get());
                auto status = stirling->RunAsThread();
                LOG(INFO) << absl::Substitute("stirling begin to run core, status:$0", status.ok());
                stirling_ = std::move(stirling);
            }
        }

        TerminationHandler::set_sinsp(inspector);
        thread inspect(do_inspect, inspector, &formatter, pid, is_syscall_on, pub);
        pub->start();

        inspector->close();
    }
    catch (const exception &e) {
        fprintf(stderr, "kindling probe init err: %s", e.what());
        return 1;
    }
    delete inspector;
    return 0;
}






