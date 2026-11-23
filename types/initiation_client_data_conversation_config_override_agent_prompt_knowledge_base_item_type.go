package types

import (
	"fmt"
)

// The type of the knowledge base
type InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemType string

const (
	InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemTypeFile   InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemType = "file"
	InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemTypeUrl    InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemType = "url"
	InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemTypeText   InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemType = "text"
	InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemTypeFolder InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemType = "folder"
)

func NewInitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemTypeFromString(s string) (InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemType, error) {
	switch s {
	case "file":
		return InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemTypeFile, nil
	case "url":
		return InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemTypeUrl, nil
	case "text":
		return InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemTypeText, nil
	case "folder":
		return InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemTypeFolder, nil
	}
	var t InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemType) Ptr() *InitiationClientDataConversationConfigOverrideAgentPromptKnowledgeBaseItemType {
	return &e
}
