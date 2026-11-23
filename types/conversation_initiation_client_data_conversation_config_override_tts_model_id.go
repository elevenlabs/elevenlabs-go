package types

import (
	"fmt"
)

// The model to use for TTS
type ConversationInitiationClientDataConversationConfigOverrideTtsModelId string

const (
	ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenTurboV2          ConversationInitiationClientDataConversationConfigOverrideTtsModelId = "eleven_turbo_v2"
	ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenTurboV25         ConversationInitiationClientDataConversationConfigOverrideTtsModelId = "eleven_turbo_v2_5"
	ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenFlashV2          ConversationInitiationClientDataConversationConfigOverrideTtsModelId = "eleven_flash_v2"
	ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenFlashV25         ConversationInitiationClientDataConversationConfigOverrideTtsModelId = "eleven_flash_v2_5"
	ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenMultilingualV2   ConversationInitiationClientDataConversationConfigOverrideTtsModelId = "eleven_multilingual_v2"
	ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenV3Conversational ConversationInitiationClientDataConversationConfigOverrideTtsModelId = "eleven_v3_conversational"
	ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenV4               ConversationInitiationClientDataConversationConfigOverrideTtsModelId = "eleven_v4"
	ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenV4Turbo          ConversationInitiationClientDataConversationConfigOverrideTtsModelId = "eleven_v4_turbo"
)

func NewConversationInitiationClientDataConversationConfigOverrideTtsModelIdFromString(s string) (ConversationInitiationClientDataConversationConfigOverrideTtsModelId, error) {
	switch s {
	case "eleven_turbo_v2":
		return ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenTurboV2, nil
	case "eleven_turbo_v2_5":
		return ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenTurboV25, nil
	case "eleven_flash_v2":
		return ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenFlashV2, nil
	case "eleven_flash_v2_5":
		return ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenFlashV25, nil
	case "eleven_multilingual_v2":
		return ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenMultilingualV2, nil
	case "eleven_v3_conversational":
		return ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenV3Conversational, nil
	case "eleven_v4":
		return ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenV4, nil
	case "eleven_v4_turbo":
		return ConversationInitiationClientDataConversationConfigOverrideTtsModelIdElevenV4Turbo, nil
	}
	var t ConversationInitiationClientDataConversationConfigOverrideTtsModelId
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationInitiationClientDataConversationConfigOverrideTtsModelId) Ptr() *ConversationInitiationClientDataConversationConfigOverrideTtsModelId {
	return &e
}
