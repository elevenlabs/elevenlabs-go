package types

import (
	"fmt"
)

// The usage mode of the knowledge base
type InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageMode string

const (
	InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageModePrompt InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageMode = "prompt"
	InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageModeAuto   InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageMode = "auto"
)

func NewInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageModeFromString(s string) (InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageMode, error) {
	switch s {
	case "prompt":
		return InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageModePrompt, nil
	case "auto":
		return InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageModeAuto, nil
	}
	var t InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageMode) Ptr() *InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemUsageMode {
	return &e
}
