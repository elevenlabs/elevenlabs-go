package types

import (
	"fmt"
)

type WebhookTranscriptFormat string

const (
	WebhookTranscriptFormatJson          WebhookTranscriptFormat = "json"
	WebhookTranscriptFormatOpentelemetry WebhookTranscriptFormat = "opentelemetry"
)

func NewWebhookTranscriptFormatFromString(s string) (WebhookTranscriptFormat, error) {
	switch s {
	case "json":
		return WebhookTranscriptFormatJson, nil
	case "opentelemetry":
		return WebhookTranscriptFormatOpentelemetry, nil
	}
	var t WebhookTranscriptFormat
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WebhookTranscriptFormat) Ptr() *WebhookTranscriptFormat {
	return &e
}
