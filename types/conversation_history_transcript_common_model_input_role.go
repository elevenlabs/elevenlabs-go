package types

import (
	"fmt"
)

type ConversationHistoryTranscriptCommonModelInputRole string

const (
	ConversationHistoryTranscriptCommonModelInputRoleUser  ConversationHistoryTranscriptCommonModelInputRole = "user"
	ConversationHistoryTranscriptCommonModelInputRoleAgent ConversationHistoryTranscriptCommonModelInputRole = "agent"
)

func NewConversationHistoryTranscriptCommonModelInputRoleFromString(s string) (ConversationHistoryTranscriptCommonModelInputRole, error) {
	switch s {
	case "user":
		return ConversationHistoryTranscriptCommonModelInputRoleUser, nil
	case "agent":
		return ConversationHistoryTranscriptCommonModelInputRoleAgent, nil
	}
	var t ConversationHistoryTranscriptCommonModelInputRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationHistoryTranscriptCommonModelInputRole) Ptr() *ConversationHistoryTranscriptCommonModelInputRole {
	return &e
}
