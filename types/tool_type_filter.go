package types

import (
	"fmt"
)

type ToolTypeFilter string

const (
	ToolTypeFilterWebhook               ToolTypeFilter = "webhook"
	ToolTypeFilterClient                ToolTypeFilter = "client"
	ToolTypeFilterApiIntegrationWebhook ToolTypeFilter = "api_integration_webhook"
)

func NewToolTypeFilterFromString(s string) (ToolTypeFilter, error) {
	switch s {
	case "webhook":
		return ToolTypeFilterWebhook, nil
	case "client":
		return ToolTypeFilterClient, nil
	case "api_integration_webhook":
		return ToolTypeFilterApiIntegrationWebhook, nil
	}
	var t ToolTypeFilter
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ToolTypeFilter) Ptr() *ToolTypeFilter {
	return &e
}
