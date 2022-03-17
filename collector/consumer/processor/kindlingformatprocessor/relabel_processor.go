package kindlingformatprocessor

import (
	"github.com/Kindling-project/kindling/collector/component"
	"github.com/Kindling-project/kindling/collector/consumer"
	"github.com/Kindling-project/kindling/collector/consumer/processor"
	"github.com/Kindling-project/kindling/collector/model"
	"github.com/Kindling-project/kindling/collector/model/constlabels"
	"go.uber.org/multierr"
	"math/rand"
)

const ProcessorName = "kindlingformatprocessor"

// RelabelProcessor generates new model.GaugeGroup according to the documentation.
type RelabelProcessor struct {
	cfg          *Config
	nextConsumer consumer.Consumer
	telemetry    *component.TelemetryTools
}

func NewRelabelProcessor(cfg interface{}, telemetry *component.TelemetryTools, nextConsumer consumer.Consumer) processor.Processor {
	processorCfg := cfg.(*Config)
	if processorCfg.SamplingRate == nil {
		processorCfg.SamplingRate = &SampleConfig{
			NormalData: 0,
			SlowData:   100,
			ErrorData:  100,
		}
	}
	return &RelabelProcessor{
		cfg:          processorCfg,
		nextConsumer: nextConsumer,
		telemetry:    telemetry,
	}
}

func (r *RelabelProcessor) Consume(gaugeGroup *model.GaugeGroup) error {
	common := newGauges(gaugeGroup)

	var traceErr error = nil
	var spanErr error = nil

	if r.cfg.NeedTraceAsMetric && common.isSlowOrError() {
		// Trace as Metric
		trace := newGauges(gaugeGroup)
		traceErr = r.nextConsumer.Consume(trace.Process(r.cfg, TraceName, TopologyTraceInstanceInfo,
			TopologyTraceK8sInfo, SrcContainerInfo, DstContainerInfo, ServiceProtocolInfo, TraceStatusInfo))
	}
	if r.cfg.NeedTraceAsResourceSpan {
		var isSample = false
		randSeed := rand.Intn(100)
		if common.isSlowOrError() {
			if (randSeed < r.cfg.SamplingRate.SlowData) && gaugeGroup.Labels.GetBoolValue(constlabels.IsSlow) {
				isSample = true
			}
			if (randSeed < r.cfg.SamplingRate.ErrorData) && gaugeGroup.Labels.GetBoolValue(constlabels.IsError) {
				isSample = true
			}
		} else {
			if randSeed < r.cfg.SamplingRate.NormalData {
				isSample = true
			}
		}
		if isSample {
			// Trace As Span
			span := newGauges(gaugeGroup)
			spanErr = r.nextConsumer.Consume(span.Process(r.cfg, SpanName, traceSpanInstanceInfo,
				TopologyTraceK8sInfo, traceSpanContainerInfo, SpanProtocolInfo, traceSpanValuesToLabel))
		}
	}

	// The data when the field is Error is true and the error Type is 2, do not generate metric
	errorType := gaugeGroup.Labels.GetIntValue(constlabels.ErrorType)
	if errorType == constlabels.ConnectFail || errorType == constlabels.NoResponse {
		return traceErr
	}

	// 处理metrics，主要生成detail指标和topology指标
	if gaugeGroup.Labels.GetBoolValue(constlabels.IsServer) {
		// Do not emit detail protocol metric at this version
		//protocol := newGauges(gaugeGroup)
		//protocolErr := r.nextConsumer.Consume(protocol.Process(r.cfg, ProtocolDetailMetricName, ServiceInstanceInfo, ServiceK8sInfo, ProtocolDetailInfo))
		// @qianlu.kk detail metrics for server, modify the labels and metric names.
		metricErr := r.nextConsumer.Consume(common.Process(r.cfg, DetailMetricName, ServiceInstanceInfo,
			ServiceK8sInfo, ServiceProtocolInfo))
		var metricErr2 error
		if r.cfg.StoreExternalSrcIP {
			srcNamespace := gaugeGroup.Labels.GetStringValue(constlabels.SrcNamespace)
			if srcNamespace == constlabels.ExternalClusterNamespace {
				// Use data from server-side to generate a topology metric only when the namespace is EXTERNAL.
				externalGaugeGroup := newGauges(gaugeGroup)
				// Here we have to modify the field "IsServer" to generate the topology metric.
				//externalGaugeGroup.Labels.AddBoolValue(constlabels.IsServer, false)
				// @qianlu.kk 服务端拓扑，srcIp是不靠谱的 topo, keep it, just modify the metric name.
				metricErr2 = r.nextConsumer.Consume(externalGaugeGroup.Process(r.cfg, TopologyMetricName,
					TopologyInstanceInfo, TopologyK8sInfo, DstContainerInfo, TopologyProtocolInfo))
				// In case of using the original data later, we reset the field "IsServer".
				//externalGaugeGroup.Labels.AddBoolValue(constlabels.IsServer, true)
			}
		}
		return multierr.Combine(traceErr, spanErr, metricErr, metricErr2)
	} else {
		// topo metrics for client, keep it, just modify the metric name.
		metricErr := r.nextConsumer.Consume(common.Process(r.cfg, TopologyMetricName, TopologyInstanceInfo,
			TopologyK8sInfo, SrcContainerInfo, DstContainerInfo, TopologyProtocolInfo))
		// TODO @qianlu.kk detail metrics for client, modify the labels and metric names.
		metricErr2 := r.nextConsumer.Consume(common.Process(r.cfg, DetailMetricName, ServiceInstanceInfo, ServiceK8sInfo,
			ServiceProtocolInfo))
		return multierr.Combine(traceErr, metricErr, metricErr2)
	}
}
