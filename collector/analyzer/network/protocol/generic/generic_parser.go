package generic

import (
	"github.com/Kindling-project/kindling/collector/analyzer/network/protocol"
)

func NewGenericParser() *protocol.ProtocolParser {
	requestParser := protocol.CreatePkgParser(fastfailGeneric(), parseGeneric())
	responseParser := protocol.CreatePkgParser(fastfailGeneric(), parseGeneric())

	return protocol.NewProtocolParser(protocol.NOSUPPORT, requestParser, responseParser, nil)
}

func fastfailGeneric() protocol.FastFailFn {
	return func(message *protocol.PayloadMessage) bool {
		return false
	}
}

func parseGeneric() protocol.ParsePkgFn {
	return func(message *protocol.PayloadMessage) (bool, bool) {
		return true, true
	}
}
