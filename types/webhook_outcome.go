package types

import (
	"fmt"
)

type WebhookOutcome string

const (
	WebhookOutcomeDelivered              WebhookOutcome = "delivered"
	WebhookOutcomeFailedRetryScheduled   WebhookOutcome = "failed_retry_scheduled"
	WebhookOutcomeFailedExhausted        WebhookOutcome = "failed_exhausted"
	WebhookOutcomeFailedNoRetry          WebhookOutcome = "failed_no_retry"
	WebhookOutcomeFailedRetryError       WebhookOutcome = "failed_retry_error"
	WebhookOutcomeDroppedQueueFull       WebhookOutcome = "dropped_queue_full"
	WebhookOutcomeDroppedPayloadTooLarge WebhookOutcome = "dropped_payload_too_large"
	WebhookOutcomeSkippedDisabled        WebhookOutcome = "skipped_disabled"
	WebhookOutcomeNotSentConfiguration   WebhookOutcome = "not_sent_configuration"
)

func NewWebhookOutcomeFromString(s string) (WebhookOutcome, error) {
	switch s {
	case "delivered":
		return WebhookOutcomeDelivered, nil
	case "failed_retry_scheduled":
		return WebhookOutcomeFailedRetryScheduled, nil
	case "failed_exhausted":
		return WebhookOutcomeFailedExhausted, nil
	case "failed_no_retry":
		return WebhookOutcomeFailedNoRetry, nil
	case "failed_retry_error":
		return WebhookOutcomeFailedRetryError, nil
	case "dropped_queue_full":
		return WebhookOutcomeDroppedQueueFull, nil
	case "dropped_payload_too_large":
		return WebhookOutcomeDroppedPayloadTooLarge, nil
	case "skipped_disabled":
		return WebhookOutcomeSkippedDisabled, nil
	case "not_sent_configuration":
		return WebhookOutcomeNotSentConfiguration, nil
	}
	var t WebhookOutcome
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WebhookOutcome) Ptr() *WebhookOutcome {
	return &e
}
