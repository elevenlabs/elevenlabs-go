package types

import (
	"fmt"
)

// Content type for the request body. Only applies to POST/PUT/PATCH requests.
type WebhookToolApiSchemaConfigOutputContentType string

const (
	WebhookToolApiSchemaConfigOutputContentTypeApplicationJson               WebhookToolApiSchemaConfigOutputContentType = "application/json"
	WebhookToolApiSchemaConfigOutputContentTypeApplicationXWwwFormUrlencoded WebhookToolApiSchemaConfigOutputContentType = "application/x-www-form-urlencoded"
)

func NewWebhookToolApiSchemaConfigOutputContentTypeFromString(s string) (WebhookToolApiSchemaConfigOutputContentType, error) {
	switch s {
	case "application/json":
		return WebhookToolApiSchemaConfigOutputContentTypeApplicationJson, nil
	case "application/x-www-form-urlencoded":
		return WebhookToolApiSchemaConfigOutputContentTypeApplicationXWwwFormUrlencoded, nil
	}
	var t WebhookToolApiSchemaConfigOutputContentType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WebhookToolApiSchemaConfigOutputContentType) Ptr() *WebhookToolApiSchemaConfigOutputContentType {
	return &e
}
