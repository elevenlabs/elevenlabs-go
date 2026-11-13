package types

import (
	"fmt"
)

type AlertingIntegrationNotifierResponseIntegrationType string

const (
	AlertingIntegrationNotifierResponseIntegrationTypePagerduty AlertingIntegrationNotifierResponseIntegrationType = "pagerduty"
	AlertingIntegrationNotifierResponseIntegrationTypeSlack     AlertingIntegrationNotifierResponseIntegrationType = "slack"
)

func NewAlertingIntegrationNotifierResponseIntegrationTypeFromString(s string) (AlertingIntegrationNotifierResponseIntegrationType, error) {
	switch s {
	case "pagerduty":
		return AlertingIntegrationNotifierResponseIntegrationTypePagerduty, nil
	case "slack":
		return AlertingIntegrationNotifierResponseIntegrationTypeSlack, nil
	}
	var t AlertingIntegrationNotifierResponseIntegrationType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AlertingIntegrationNotifierResponseIntegrationType) Ptr() *AlertingIntegrationNotifierResponseIntegrationType {
	return &e
}
