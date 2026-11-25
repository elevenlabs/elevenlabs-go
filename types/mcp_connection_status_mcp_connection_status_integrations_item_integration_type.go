package types

import (
	"fmt"
)

type McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationType string

const (
	McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationTypeMcpServer      McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationType = "mcp_server"
	McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationTypeMcpIntegration McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationType = "mcp_integration"
)

func NewMcpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationTypeFromString(s string) (McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationType, error) {
	switch s {
	case "mcp_server":
		return McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationTypeMcpServer, nil
	case "mcp_integration":
		return McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationTypeMcpIntegration, nil
	}
	var t McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationType) Ptr() *McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationType {
	return &e
}
