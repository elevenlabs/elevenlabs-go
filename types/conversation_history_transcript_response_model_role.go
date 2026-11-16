package types

import (
	"fmt"
)

type ConversationHistoryTranscriptResponseModelRole string

const (
	ConversationHistoryTranscriptResponseModelRoleUser  ConversationHistoryTranscriptResponseModelRole = "user"
	ConversationHistoryTranscriptResponseModelRoleAgent ConversationHistoryTranscriptResponseModelRole = "agent"
)

func NewConversationHistoryTranscriptResponseModelRoleFromString(s string) (ConversationHistoryTranscriptResponseModelRole, error) {
	switch s {
	case "user":
		return ConversationHistoryTranscriptResponseModelRoleUser, nil
	case "agent":
		return ConversationHistoryTranscriptResponseModelRoleAgent, nil
	}
	var t ConversationHistoryTranscriptResponseModelRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationHistoryTranscriptResponseModelRole) Ptr() *ConversationHistoryTranscriptResponseModelRole {
	return &e
}
