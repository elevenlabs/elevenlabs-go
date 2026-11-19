package types

import (
	"fmt"
)

// The HTTP method to use for the webhook
type WebhookToolApiSchemaConfigInputMethod string

const (
	WebhookToolApiSchemaConfigInputMethodGet    WebhookToolApiSchemaConfigInputMethod = "GET"
	WebhookToolApiSchemaConfigInputMethodPost   WebhookToolApiSchemaConfigInputMethod = "POST"
	WebhookToolApiSchemaConfigInputMethodPut    WebhookToolApiSchemaConfigInputMethod = "PUT"
	WebhookToolApiSchemaConfigInputMethodPatch  WebhookToolApiSchemaConfigInputMethod = "PATCH"
	WebhookToolApiSchemaConfigInputMethodDelete WebhookToolApiSchemaConfigInputMethod = "DELETE"
)

func NewWebhookToolApiSchemaConfigInputMethodFromString(s string) (WebhookToolApiSchemaConfigInputMethod, error) {
	switch s {
	case "GET":
		return WebhookToolApiSchemaConfigInputMethodGet, nil
	case "POST":
		return WebhookToolApiSchemaConfigInputMethodPost, nil
	case "PUT":
		return WebhookToolApiSchemaConfigInputMethodPut, nil
	case "PATCH":
		return WebhookToolApiSchemaConfigInputMethodPatch, nil
	case "DELETE":
		return WebhookToolApiSchemaConfigInputMethodDelete, nil
	}
	var t WebhookToolApiSchemaConfigInputMethod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WebhookToolApiSchemaConfigInputMethod) Ptr() *WebhookToolApiSchemaConfigInputMethod {
	return &e
}
