package types

import (
	"fmt"
)

// Content type for the request body. Only applies to POST/PUT/PATCH requests.
type WebhookToolApiSchemaConfigInputContentType string

const (
	WebhookToolApiSchemaConfigInputContentTypeApplicationJson               WebhookToolApiSchemaConfigInputContentType = "application/json"
	WebhookToolApiSchemaConfigInputContentTypeApplicationXWwwFormUrlencoded WebhookToolApiSchemaConfigInputContentType = "application/x-www-form-urlencoded"
)

func NewWebhookToolApiSchemaConfigInputContentTypeFromString(s string) (WebhookToolApiSchemaConfigInputContentType, error) {
	switch s {
	case "application/json":
		return WebhookToolApiSchemaConfigInputContentTypeApplicationJson, nil
	case "application/x-www-form-urlencoded":
		return WebhookToolApiSchemaConfigInputContentTypeApplicationXWwwFormUrlencoded, nil
	}
	var t WebhookToolApiSchemaConfigInputContentType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WebhookToolApiSchemaConfigInputContentType) Ptr() *WebhookToolApiSchemaConfigInputContentType {
	return &e
}
