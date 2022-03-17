package constlabels

const (
	ContentKey = "content_key"

	//HttpMethod          = "http_method"
	HttpUrl             = "http_url"
	HttpApmTraceType    = "trace_type"
	HttpApmTraceId      = "trace_id"
	HttpRequestPayload  = "request_payload"
	HttpResponsePayload = "response_payload"
	//HttpStatusCode      = "http_status_code"

	//DnsId     = "dns_id"
	//DnsDomain = "dns_domain"
	//DnsRcode  = "dns_rcode"
	//DnsIp     = "dns_ip"

	Sql        = "sql"
	SqlErrCode = "sql_error_code"
	SqlErrMsg  = "sql_error_msg"

	RedisErrMsg = "redis_error_msg"

	KafkaApi           = "kafka_api"
	KafkaVersion       = "kafka_version"
	KafkaCorrelationId = "kafka_id"
	KafkaTopic         = "kafka_topic"
	KafkaPartition     = "kafka_partition"
	KafkaErrorCode     = "kafka_error_code"


	// ***** ARMS Metrics 2.0 *****
	// -- HTTP dimensions --

	HttpMethod = "httpMethod"
	HttpStatusCode = "HttpStatusCode"

	// -- RPC dimensions --

	RpcService = "rpcService"
	RpcMethod = "rpcMethod"

	// -- Database dimensions --

	DbStatement = "dbStatement"
	DbName = "dbName"
	DbTableName = "dbTableName"
	DbOperation = "dbOperation"

	// -- Message dimensions --
	MsgTopic = "msgTopic"
	MsgProtocol = "msgProtocol"
	MsgProtocolVersion = "msgProtocolVersion"
	MsgConsumerGroup = "msgConsumerGroup"

	// -- DNS dimensions --
	DnsId     = "dnsId"
	DnsDomain = "reqDomain"
	DnsRcode  = "dnsRetCode"
	DnsIp     = "resp"
)
