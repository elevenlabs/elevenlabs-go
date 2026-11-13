package types

import (
	"fmt"
)

type ToolType string

const (
	ToolTypeSystem                ToolType = "system"
	ToolTypeWebhook               ToolType = "webhook"
	ToolTypeClient                ToolType = "client"
	ToolTypeMcp                   ToolType = "mcp"
	ToolTypeWorkflow              ToolType = "workflow"
	ToolTypeApiIntegrationWebhook ToolType = "api_integration_webhook"
	ToolTypeSmb                   ToolType = "smb"
)

func NewToolTypeFromString(s string) (ToolType, error) {
	switch s {
	case "system":
		return ToolTypeSystem, nil
	case "webhook":
		return ToolTypeWebhook, nil
	case "client":
		return ToolTypeClient, nil
	case "mcp":
		return ToolTypeMcp, nil
	case "workflow":
		return ToolTypeWorkflow, nil
	case "api_integration_webhook":
		return ToolTypeApiIntegrationWebhook, nil
	case "smb":
		return ToolTypeSmb, nil
	}
	var t ToolType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ToolType) Ptr() *ToolType {
	return &e
}
