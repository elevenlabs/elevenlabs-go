package types

import (
	"fmt"
)

type McpTransport string

const (
	McpTransportSse            McpTransport = "SSE"
	McpTransportStreamableHttp McpTransport = "STREAMABLE_HTTP"
)

func NewMcpTransportFromString(s string) (McpTransport, error) {
	switch s {
	case "SSE":
		return McpTransportSse, nil
	case "STREAMABLE_HTTP":
		return McpTransportStreamableHttp, nil
	}
	var t McpTransport
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e McpTransport) Ptr() *McpTransport {
	return &e
}
