package types

import (
	"fmt"
)

// Supported MCP server transport types.
type McpServerTransport string

const (
	McpServerTransportSse            McpServerTransport = "SSE"
	McpServerTransportStreamableHttp McpServerTransport = "STREAMABLE_HTTP"
)

func NewMcpServerTransportFromString(s string) (McpServerTransport, error) {
	switch s {
	case "SSE":
		return McpServerTransportSse, nil
	case "STREAMABLE_HTTP":
		return McpServerTransportStreamableHttp, nil
	}
	var t McpServerTransport
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e McpServerTransport) Ptr() *McpServerTransport {
	return &e
}
