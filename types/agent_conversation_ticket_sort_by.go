package types

import (
	"fmt"
)

type AgentConversationTicketSortBy string

const (
	AgentConversationTicketSortByCreatedAt    AgentConversationTicketSortBy = "created_at"
	AgentConversationTicketSortByPriority     AgentConversationTicketSortBy = "priority"
	AgentConversationTicketSortByLastReported AgentConversationTicketSortBy = "last_reported"
)

func NewAgentConversationTicketSortByFromString(s string) (AgentConversationTicketSortBy, error) {
	switch s {
	case "created_at":
		return AgentConversationTicketSortByCreatedAt, nil
	case "priority":
		return AgentConversationTicketSortByPriority, nil
	case "last_reported":
		return AgentConversationTicketSortByLastReported, nil
	}
	var t AgentConversationTicketSortBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentConversationTicketSortBy) Ptr() *AgentConversationTicketSortBy {
	return &e
}
