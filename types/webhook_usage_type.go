package types

import (
	"fmt"
)

type WebhookUsageType string

const (
	WebhookUsageTypeConvAiAgentSettings        WebhookUsageType = "ConvAI Agent Settings"
	WebhookUsageTypeConvAiSettings             WebhookUsageType = "ConvAI Settings"
	WebhookUsageTypeVoiceLibraryRemovalNotices WebhookUsageType = "Voice Library Removal Notices"
	WebhookUsageTypeSpeechToText               WebhookUsageType = "Speech to Text"
	WebhookUsageTypeAgentQaEvaluations         WebhookUsageType = "Agent QA Evaluations"
	WebhookUsageTypeConvAiAlerting             WebhookUsageType = "ConvAI Alerting"
	WebhookUsageTypeFlows                      WebhookUsageType = "Flows"
	WebhookUsageTypeDubbing                    WebhookUsageType = "Dubbing"
)

func NewWebhookUsageTypeFromString(s string) (WebhookUsageType, error) {
	switch s {
	case "ConvAI Agent Settings":
		return WebhookUsageTypeConvAiAgentSettings, nil
	case "ConvAI Settings":
		return WebhookUsageTypeConvAiSettings, nil
	case "Voice Library Removal Notices":
		return WebhookUsageTypeVoiceLibraryRemovalNotices, nil
	case "Speech to Text":
		return WebhookUsageTypeSpeechToText, nil
	case "Agent QA Evaluations":
		return WebhookUsageTypeAgentQaEvaluations, nil
	case "ConvAI Alerting":
		return WebhookUsageTypeConvAiAlerting, nil
	case "Flows":
		return WebhookUsageTypeFlows, nil
	case "Dubbing":
		return WebhookUsageTypeDubbing, nil
	}
	var t WebhookUsageType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WebhookUsageType) Ptr() *WebhookUsageType {
	return &e
}
