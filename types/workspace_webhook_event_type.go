package types

import (
	"fmt"
)

type WorkspaceWebhookEventType string

const (
	WorkspaceWebhookEventTypeVoiceLibraryRemovalNotice WorkspaceWebhookEventType = "voice_library_removal_notice"
	WorkspaceWebhookEventTypeSpeechToText              WorkspaceWebhookEventType = "speech_to_text"
	WorkspaceWebhookEventTypeAgentQa                   WorkspaceWebhookEventType = "agent_qa"
	WorkspaceWebhookEventTypeFlows                     WorkspaceWebhookEventType = "flows"
)

func NewWorkspaceWebhookEventTypeFromString(s string) (WorkspaceWebhookEventType, error) {
	switch s {
	case "voice_library_removal_notice":
		return WorkspaceWebhookEventTypeVoiceLibraryRemovalNotice, nil
	case "speech_to_text":
		return WorkspaceWebhookEventTypeSpeechToText, nil
	case "agent_qa":
		return WorkspaceWebhookEventTypeAgentQa, nil
	case "flows":
		return WorkspaceWebhookEventTypeFlows, nil
	}
	var t WorkspaceWebhookEventType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WorkspaceWebhookEventType) Ptr() *WorkspaceWebhookEventType {
	return &e
}
