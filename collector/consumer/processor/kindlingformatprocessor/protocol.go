package kindlingformatprocessor

import (
	"github.com/Kindling-project/kindling/collector/model/constlabels"
	"strconv"
)

type ProtocolType string

const (
	http  = "http"
	http2 = "http2"
	grpc  = "grpc"
	dubbo = "dubbo"
	dns   = "dns"
	kafka = "kafka"
	mysql = "mysql"
)

func fillSpecialProtocolLabels(g *gauges, protocol ProtocolType) {
	switch protocol {
	case kafka:
		fillKafkaMetricProtocolLabel(g)
	default:
		// Do nothing
	}
}

func fillSpanProtocolLabels(g *gauges, protocol ProtocolType) {
	switch protocol {
	case http:
		fillSpanHttpProtocolLabel(g)
	case dns:
		fillSpanDNSProtocolLabel(g)
	case mysql:
		fillSpanMysqlProtocolLabel(g)
	}
}

func fillSpanMysqlProtocolLabel(g *gauges) {
	g.targetLabels.AddStringValue("mysql.sql", g.Labels.GetStringValue(constlabels.Sql))
	g.targetLabels.AddStringValue("mysql.error_code", strconv.FormatInt(g.Labels.GetIntValue(constlabels.SqlErrCode), 10))
	g.targetLabels.AddStringValue("mysql.error_msg", g.Labels.GetStringValue(constlabels.SqlErrMsg))
}

func fillSpanDNSProtocolLabel(g *gauges) {
	g.targetLabels.AddStringValue("dns.domain", g.Labels.GetStringValue(constlabels.DnsDomain))
	g.targetLabels.AddStringValue("dns.rcode", strconv.FormatInt(g.Labels.GetIntValue(constlabels.DnsRcode), 10))
}

func fillCommonProtocolLabelsV2(g *gauges, protocol ProtocolType, needProtocolDetail bool) {
	switch protocol {
	case http:
		if needProtocolDetail {
			fillEntityHttpProtocolLabelV2(g)
		} else {
			fillTopologyHttpProtocolLabel(g)
		}
	case dns:
		if needProtocolDetail {
			fillEntityDnsProtocolLabel(g)
		} else {
			fillTopologyDnsProtocolLabel(g)
		}
	case kafka:
		if needProtocolDetail {
			fillEntityKafkaProtocolLabel(g)
		} else {
			fillTopologyKafkaProtocolLabel(g)
		}
	case mysql:
		if needProtocolDetail {
			fillEntityMysqlProtocolLabel(g)
		} else {
			fillTopologyMysqlProtocolLabel(g)
		}
	case grpc:
		if needProtocolDetail {
			fillEntityHttpProtocolLabelV2(g)
		} else {
			fillTopologyHttpProtocolLabel(g)
		}
	default:
		// Do nothing ?
	}
}

func fillCommonProtocolLabels(g *gauges, protocol ProtocolType, isServer bool) {
	switch protocol {
	case http:
		if isServer {
			fillEntityHttpProtocolLabel(g)
		} else {
			fillTopologyHttpProtocolLabel(g)
		}
	case dns:
		if isServer {
			fillEntityDnsProtocolLabel(g)
		} else {
			fillTopologyDnsProtocolLabel(g)
		}
	case kafka:
		if isServer {
			fillEntityKafkaProtocolLabel(g)
		} else {
			fillTopologyKafkaProtocolLabel(g)
		}
	case mysql:
		if isServer {
			fillEntityMysqlProtocolLabel(g)
		} else {
			fillTopologyMysqlProtocolLabel(g)
		}
	case grpc:
		if isServer {
			fillEntityHttpProtocolLabel(g)
		} else {
			fillTopologyHttpProtocolLabel(g)
		}
	default:
		// Do nothing ?
	}
}

func fillEntityHttpProtocolLabelV2(g *gauges) {
	g.targetLabels.AddStringValue(constlabels.HttpMethod, g.Labels.GetStringValue(constlabels.ContentKey))
	g.targetLabels.AddStringValue(constlabels.ArmsHttpStatusCode, strconv.FormatInt(g.Labels.GetIntValue(constlabels.HttpStatusCode) / 100, 10) + "xx")
}

func fillEntityHttpProtocolLabel(g *gauges) {
	g.targetLabels.AddStringValue(constlabels.RequestContent, g.Labels.GetStringValue(constlabels.ContentKey))
	g.targetLabels.AddStringValue(constlabels.ResponseContent, strconv.FormatInt(g.Labels.GetIntValue(constlabels.HttpStatusCode), 10))
}

func fillTopologyHttpProtocolLabel(g *gauges) {
	g.targetLabels.AddStringValue(constlabels.StatusCode, strconv.FormatInt(g.Labels.GetIntValue(constlabels.HttpStatusCode) / 100, 10) + "xx")
}

func fillSpanHttpProtocolLabel(g *gauges) {
	g.targetLabels.AddStringValue("http.method", g.Labels.GetStringValue(constlabels.HttpMethod))
	g.targetLabels.AddStringValue("http.endpoint", g.Labels.GetStringValue(constlabels.HttpUrl))
	g.targetLabels.AddIntValue("http.status_code", g.Labels.GetIntValue(constlabels.HttpStatusCode))
	g.targetLabels.AddStringValue("http.trace_id", g.Labels.GetStringValue(constlabels.HttpApmTraceId))
	g.targetLabels.AddStringValue("http.trace_type", g.Labels.GetStringValue(constlabels.HttpApmTraceType))
	g.targetLabels.AddStringValue("http.request_headers", g.Labels.GetStringValue(constlabels.HttpRequestPayload))
	g.targetLabels.AddStringValue("http.request_body", "")
	g.targetLabels.AddStringValue("http.response_headers", g.Labels.GetStringValue(constlabels.HttpResponsePayload))
	g.targetLabels.AddStringValue("http.response_body", "")
}

func fillEntityDnsProtocolLabel(g *gauges) {
	g.targetLabels.AddStringValue(constlabels.RequestContent, g.Labels.GetStringValue(constlabels.DnsDomain))
	g.targetLabels.AddStringValue(constlabels.ResponseContent, strconv.FormatInt(g.Labels.GetIntValue(constlabels.DnsRcode), 10))
}

func fillTopologyDnsProtocolLabel(g *gauges) {
	g.targetLabels.AddStringValue(constlabels.StatusCode, strconv.FormatInt(g.Labels.GetIntValue(constlabels.DnsRcode), 10))
}

func fillEntityRpcProtocolLabelV2(g *gauges) {
	g.targetLabels.AddStringValue(constlabels.RpcMethod, g.Labels.GetStringValue(constlabels.RpcMethod))
	g.targetLabels.AddStringValue(constlabels.RpcService, g.Labels.GetStringValue(constlabels.RpcMethod))
}

func fillEntityKafkaProtocolLabelV2(g *gauges) {
	g.targetLabels.AddStringValue(constlabels.RequestContent, g.Labels.GetStringValue(constlabels.KafkaTopic))
	g.targetLabels.AddStringValue(constlabels.ResponseContent, g.Labels.GetStringValue(constlabels.STR_EMPTY))
}

func fillEntityKafkaProtocolLabel(g *gauges) {
	g.targetLabels.AddStringValue(constlabels.RequestContent, g.Labels.GetStringValue(constlabels.KafkaTopic))
	g.targetLabels.AddStringValue(constlabels.ResponseContent, g.Labels.GetStringValue(constlabels.STR_EMPTY))
}

func fillTopologyKafkaProtocolLabel(g *gauges) {
	g.targetLabels.AddStringValue(constlabels.StatusCode, g.Labels.GetStringValue(constlabels.STR_EMPTY))
}

func fillEntityDatabaseProtocolLabelV2(g *gauges) {
	g.targetLabels.AddStringValue(constlabels.DbStatement, g.Labels.GetStringValue(constlabels.ContentKey))
	g.targetLabels.AddStringValue(constlabels.StatusCode, strconv.FormatInt(g.Labels.GetIntValue(constlabels.SqlErrCode), 10))
}

func fillEntityMysqlProtocolLabel(g *gauges) {
	g.targetLabels.AddStringValue(constlabels.RequestContent, g.Labels.GetStringValue(constlabels.ContentKey))
	g.targetLabels.AddStringValue(constlabels.ResponseContent, strconv.FormatInt(g.Labels.GetIntValue(constlabels.SqlErrCode), 10))
}

func fillTopologyMysqlProtocolLabel(g *gauges) {
	g.targetLabels.AddStringValue(constlabels.StatusCode, strconv.FormatInt(g.Labels.GetIntValue(constlabels.SqlErrCode), 10))
}

func fillKafkaMetricProtocolLabel(g *gauges) {
	// TODO Missing Information Element
	g.targetLabels.AddStringValue(constlabels.Topic, g.Labels.GetStringValue(constlabels.KafkaTopic))
	//g.targetLabels.AddStringValue(constlabels.Operation,g.Labels.GetStringValue())
	//g.targetLabels.AddStringValue(constlabels.ConsumerId, g.Labels.GetStringValue())
}
