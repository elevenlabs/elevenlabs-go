package types

import (
	"fmt"
)

type ConversationAgentResponseAgentResponseEventResponseType string

const (
	ConversationAgentResponseAgentResponseEventResponseTypeText  ConversationAgentResponseAgentResponseEventResponseType = "text"
	ConversationAgentResponseAgentResponseEventResponseTypeVoice ConversationAgentResponseAgentResponseEventResponseType = "voice"
)

func NewConversationAgentResponseAgentResponseEventResponseTypeFromString(s string) (ConversationAgentResponseAgentResponseEventResponseType, error) {
	switch s {
	case "text":
		return ConversationAgentResponseAgentResponseEventResponseTypeText, nil
	case "voice":
		return ConversationAgentResponseAgentResponseEventResponseTypeVoice, nil
	}
	var t ConversationAgentResponseAgentResponseEventResponseType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationAgentResponseAgentResponseEventResponseType) Ptr() *ConversationAgentResponseAgentResponseEventResponseType {
	return &e
}
