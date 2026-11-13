package types

import (
	"fmt"
)

type AlertingWebhookMethod string

const (
	AlertingWebhookMethodPost AlertingWebhookMethod = "POST"
	AlertingWebhookMethodPut  AlertingWebhookMethod = "PUT"
)

func NewAlertingWebhookMethodFromString(s string) (AlertingWebhookMethod, error) {
	switch s {
	case "POST":
		return AlertingWebhookMethodPost, nil
	case "PUT":
		return AlertingWebhookMethodPut, nil
	}
	var t AlertingWebhookMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AlertingWebhookMethod) Ptr() *AlertingWebhookMethod {
	return &e
}
