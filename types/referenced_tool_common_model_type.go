package types

import (
	"fmt"
)

// The type of the tool
type ReferencedToolCommonModelType string

const (
	ReferencedToolCommonModelTypeSystem                ReferencedToolCommonModelType = "system"
	ReferencedToolCommonModelTypeWebhook               ReferencedToolCommonModelType = "webhook"
	ReferencedToolCommonModelTypeClient                ReferencedToolCommonModelType = "client"
	ReferencedToolCommonModelTypeWorkflow              ReferencedToolCommonModelType = "workflow"
	ReferencedToolCommonModelTypeApiIntegrationWebhook ReferencedToolCommonModelType = "api_integration_webhook"
	ReferencedToolCommonModelTypeMcp                   ReferencedToolCommonModelType = "mcp"
	ReferencedToolCommonModelTypeCode                  ReferencedToolCommonModelType = "code"
)

func NewReferencedToolCommonModelTypeFromString(s string) (ReferencedToolCommonModelType, error) {
	switch s {
	case "system":
		return ReferencedToolCommonModelTypeSystem, nil
	case "webhook":
		return ReferencedToolCommonModelTypeWebhook, nil
	case "client":
		return ReferencedToolCommonModelTypeClient, nil
	case "workflow":
		return ReferencedToolCommonModelTypeWorkflow, nil
	case "api_integration_webhook":
		return ReferencedToolCommonModelTypeApiIntegrationWebhook, nil
	case "mcp":
		return ReferencedToolCommonModelTypeMcp, nil
	case "code":
		return ReferencedToolCommonModelTypeCode, nil
	}
	var t ReferencedToolCommonModelType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReferencedToolCommonModelType) Ptr() *ReferencedToolCommonModelType {
	return &e
}
