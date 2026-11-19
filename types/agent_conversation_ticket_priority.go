package types

import (
	"fmt"
)

type AgentConversationTicketPriority string

const (
	AgentConversationTicketPriorityLow    AgentConversationTicketPriority = "low"
	AgentConversationTicketPriorityMedium AgentConversationTicketPriority = "medium"
	AgentConversationTicketPriorityHigh   AgentConversationTicketPriority = "high"
	AgentConversationTicketPriorityUrgent AgentConversationTicketPriority = "urgent"
)

func NewAgentConversationTicketPriorityFromString(s string) (AgentConversationTicketPriority, error) {
	switch s {
	case "low":
		return AgentConversationTicketPriorityLow, nil
	case "medium":
		return AgentConversationTicketPriorityMedium, nil
	case "high":
		return AgentConversationTicketPriorityHigh, nil
	case "urgent":
		return AgentConversationTicketPriorityUrgent, nil
	}
	var t AgentConversationTicketPriority
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentConversationTicketPriority) Ptr() *AgentConversationTicketPriority {
	return &e
}
