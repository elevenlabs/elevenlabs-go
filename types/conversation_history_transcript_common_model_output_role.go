package types

import (
	"fmt"
)

type ConversationHistoryTranscriptCommonModelOutputRole string

const (
	ConversationHistoryTranscriptCommonModelOutputRoleUser  ConversationHistoryTranscriptCommonModelOutputRole = "user"
	ConversationHistoryTranscriptCommonModelOutputRoleAgent ConversationHistoryTranscriptCommonModelOutputRole = "agent"
)

func NewConversationHistoryTranscriptCommonModelOutputRoleFromString(s string) (ConversationHistoryTranscriptCommonModelOutputRole, error) {
	switch s {
	case "user":
		return ConversationHistoryTranscriptCommonModelOutputRoleUser, nil
	case "agent":
		return ConversationHistoryTranscriptCommonModelOutputRoleAgent, nil
	}
	var t ConversationHistoryTranscriptCommonModelOutputRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationHistoryTranscriptCommonModelOutputRole) Ptr() *ConversationHistoryTranscriptCommonModelOutputRole {
	return &e
}
