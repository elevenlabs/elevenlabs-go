package types

import (
	"fmt"
)

type AgentConversationTicketSource string

const (
	AgentConversationTicketSourceQa     AgentConversationTicketSource = "qa"
	AgentConversationTicketSourceAgent  AgentConversationTicketSource = "agent"
	AgentConversationTicketSourceManual AgentConversationTicketSource = "manual"
)

func NewAgentConversationTicketSourceFromString(s string) (AgentConversationTicketSource, error) {
	switch s {
	case "qa":
		return AgentConversationTicketSourceQa, nil
	case "agent":
		return AgentConversationTicketSourceAgent, nil
	case "manual":
		return AgentConversationTicketSourceManual, nil
	}
	var t AgentConversationTicketSource
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentConversationTicketSource) Ptr() *AgentConversationTicketSource {
	return &e
}
