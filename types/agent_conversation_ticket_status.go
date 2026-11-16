package types

import (
	"fmt"
)

type AgentConversationTicketStatus string

const (
	AgentConversationTicketStatusOpen       AgentConversationTicketStatus = "open"
	AgentConversationTicketStatusInProgress AgentConversationTicketStatus = "in_progress"
	AgentConversationTicketStatusResolved   AgentConversationTicketStatus = "resolved"
	AgentConversationTicketStatusCancelled  AgentConversationTicketStatus = "cancelled"
	AgentConversationTicketStatusMerged     AgentConversationTicketStatus = "merged"
)

func NewAgentConversationTicketStatusFromString(s string) (AgentConversationTicketStatus, error) {
	switch s {
	case "open":
		return AgentConversationTicketStatusOpen, nil
	case "in_progress":
		return AgentConversationTicketStatusInProgress, nil
	case "resolved":
		return AgentConversationTicketStatusResolved, nil
	case "cancelled":
		return AgentConversationTicketStatusCancelled, nil
	case "merged":
		return AgentConversationTicketStatusMerged, nil
	}
	var t AgentConversationTicketStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentConversationTicketStatus) Ptr() *AgentConversationTicketStatus {
	return &e
}
