package types

import (
	"fmt"
)

type IntegrationType string

const (
	IntegrationTypeMcpServer      IntegrationType = "mcp_server"
	IntegrationTypeMcpIntegration IntegrationType = "mcp_integration"
)

func NewIntegrationTypeFromString(s string) (IntegrationType, error) {
	switch s {
	case "mcp_server":
		return IntegrationTypeMcpServer, nil
	case "mcp_integration":
		return IntegrationTypeMcpIntegration, nil
	}
	var t IntegrationType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e IntegrationType) Ptr() *IntegrationType {
	return &e
}
