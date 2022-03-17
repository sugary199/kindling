package constlabels

import "github.com/Kindling-project/kindling/collector/model/constvalues"

// key1: originName key2: isServer
var metricNameDictionary = map[string]map[bool]string{
	constvalues.RequestIo:        {true: EntityRequestIoMetric, false: TopologyRequestIoMetric},
	constvalues.ResponseIo:       {true: EntityResponseIoMetric, false: TopologyResponseIoMetric},
	constvalues.RequestTotalTime: {true: EntityRequestLatencyMetric, false: TopologyRequestLatencyMetric},
}

var armsMetricNameDictionary = map[string]map[bool]string{
	constvalues.RequestIo:        {true: EntityRequestIoMetric, false: TopologyRequestIoMetric},
	constvalues.ResponseIo:       {true: EntityResponseIoMetric, false: TopologyResponseIoMetric},
	constvalues.RequestTotalTime: {true: ArmsEntityDurationMetric, false: ArmsTopologyRequestLatencyMetric},
}

var protocol2type = map[string]string {
	"http": "http",
	"mysql": "db",
	"redis": "db",
	"kafka": "message",
	"NOSUPPORT": "unknown",
	"dns": "dns",
	"grpc": "rpc",
}

const (
	TopologyRequestIoMetric      = "request_bytes_total"
	TopologyResponseIoMetric     = "response_bytes_total"
	TopologyRequestLatencyMetric = "duration_nanoseconds"

	EntityRequestIoMetric      = "receive_bytes_total"
	EntityResponseIoMetric     = "send_bytes_total"
	EntityRequestLatencyMetric = "duration_nanoseconds"

	// ======= ARMS Metric2.0 =======
	ArmsEntityDurationMetric = "duration_ms"
	ArmsTopologyRequestLatencyMetric = "duration_ms"
	// ==============================
)

const (
	NPMPrefixKindling = "kindling"

	ArmsPrefix = "arms"

	EntityPrefix   = "entity"
	TopologyPrefix = "topology"
)

func ToKindlingTraceAsMetricName() string {
	return NPMPrefixKindling + "_trace_request_" + TopologyRequestLatencyMetric
}

func ToType(protocol string) string {
	if _type, ok := protocol2type[protocol]; ok {
		return _type
	}
	return ""
}

func ToTopologyMetricName(origName string) string {
	if names, ok := armsMetricNameDictionary[origName]; !ok {
		return ""
	} else {
		return getArmsPrefix(false) + "request_" + names[false]
	}
}

func ToDetailMetricName(origName string) string {
	if names, ok := armsMetricNameDictionary[origName]; !ok {
		return ""
	} else {
		return getArmsPrefix(true) + "request_" + names[true]
	}
}

func ToKindlingMetricName(origName string, isServer bool) string {
	if names, ok := metricNameDictionary[origName]; !ok {
		return ""
	} else {
		return getKindlingPrefix(isServer) + "request_" + names[isServer]
	}
}

//ToKindlingDetailMetricName For ServerDetail Metric
func ToKindlingDetailMetricName(origName string, protocol string) string {
	if names, ok := metricNameDictionary[origName]; !ok {
		return ""
	} else {
		return getKindlingPrefix(true) + protocol + "_" + names[true]
	}
}

func getArmsPrefix(isDetail bool) string {
	if isDetail {
		return ArmsPrefix + "_" + EntityPrefix + "_"
	} else {
		return ArmsPrefix + "_" + TopologyPrefix + "_"
	}
}

func getKindlingPrefix(isServer bool) string {
	var kindMark string
	if isServer {
		kindMark = EntityPrefix
	} else {
		kindMark = TopologyPrefix
	}
	return NPMPrefixKindling + "_" + kindMark + "_"
}
