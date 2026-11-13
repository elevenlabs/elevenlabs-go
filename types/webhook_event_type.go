package types

import (
	"fmt"
)

type WebhookEventType string

const (
	WebhookEventTypeTranscript                WebhookEventType = "transcript"
	WebhookEventTypeAudio                     WebhookEventType = "audio"
	WebhookEventTypeCallInitiationFailure     WebhookEventType = "call_initiation_failure"
	WebhookEventTypeAnsweringMachineDetection WebhookEventType = "answering_machine_detection"
	WebhookEventTypeUnredactedTranscript      WebhookEventType = "unredacted_transcript"
	WebhookEventTypeUnredactedAudio           WebhookEventType = "unredacted_audio"
)

func NewWebhookEventTypeFromString(s string) (WebhookEventType, error) {
	switch s {
	case "transcript":
		return WebhookEventTypeTranscript, nil
	case "audio":
		return WebhookEventTypeAudio, nil
	case "call_initiation_failure":
		return WebhookEventTypeCallInitiationFailure, nil
	case "answering_machine_detection":
		return WebhookEventTypeAnsweringMachineDetection, nil
	case "unredacted_transcript":
		return WebhookEventTypeUnredactedTranscript, nil
	case "unredacted_audio":
		return WebhookEventTypeUnredactedAudio, nil
	}
	var t WebhookEventType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WebhookEventType) Ptr() *WebhookEventType {
	return &e
}
