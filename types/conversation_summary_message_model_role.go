package types

import (
	"fmt"
)

type ConversationSummaryMessageModelRole string

const (
	ConversationSummaryMessageModelRoleUser  ConversationSummaryMessageModelRole = "user"
	ConversationSummaryMessageModelRoleAgent ConversationSummaryMessageModelRole = "agent"
)

func NewConversationSummaryMessageModelRoleFromString(s string) (ConversationSummaryMessageModelRole, error) {
	switch s {
	case "user":
		return ConversationSummaryMessageModelRoleUser, nil
	case "agent":
		return ConversationSummaryMessageModelRoleAgent, nil
	}
	var t ConversationSummaryMessageModelRole
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConversationSummaryMessageModelRole) Ptr() *ConversationSummaryMessageModelRole {
	return &e
}
