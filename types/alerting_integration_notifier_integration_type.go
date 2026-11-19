package types

import (
	"fmt"
)

// Integration to deliver to. The server treats an omitted value as “pagerduty“.
type AlertingIntegrationNotifierIntegrationType string

const (
	AlertingIntegrationNotifierIntegrationTypePagerduty AlertingIntegrationNotifierIntegrationType = "pagerduty"
	AlertingIntegrationNotifierIntegrationTypeSlack     AlertingIntegrationNotifierIntegrationType = "slack"
)

func NewAlertingIntegrationNotifierIntegrationTypeFromString(s string) (AlertingIntegrationNotifierIntegrationType, error) {
	switch s {
	case "pagerduty":
		return AlertingIntegrationNotifierIntegrationTypePagerduty, nil
	case "slack":
		return AlertingIntegrationNotifierIntegrationTypeSlack, nil
	}
	var t AlertingIntegrationNotifierIntegrationType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AlertingIntegrationNotifierIntegrationType) Ptr() *AlertingIntegrationNotifierIntegrationType {
	return &e
}
