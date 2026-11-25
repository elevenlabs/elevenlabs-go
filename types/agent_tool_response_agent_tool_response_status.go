package types

import (
	"fmt"
)

// Tool call status derived from execution flags.
type AgentToolResponseAgentToolResponseStatus string

const (
	AgentToolResponseAgentToolResponseStatusSuccess AgentToolResponseAgentToolResponseStatus = "success"
	AgentToolResponseAgentToolResponseStatusError   AgentToolResponseAgentToolResponseStatus = "error"
	AgentToolResponseAgentToolResponseStatusBlocked AgentToolResponseAgentToolResponseStatus = "blocked"
	AgentToolResponseAgentToolResponseStatusSkipped AgentToolResponseAgentToolResponseStatus = "skipped"
)

func NewAgentToolResponseAgentToolResponseStatusFromString(s string) (AgentToolResponseAgentToolResponseStatus, error) {
	switch s {
	case "success":
		return AgentToolResponseAgentToolResponseStatusSuccess, nil
	case "error":
		return AgentToolResponseAgentToolResponseStatusError, nil
	case "blocked":
		return AgentToolResponseAgentToolResponseStatusBlocked, nil
	case "skipped":
		return AgentToolResponseAgentToolResponseStatusSkipped, nil
	}
	var t AgentToolResponseAgentToolResponseStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentToolResponseAgentToolResponseStatus) Ptr() *AgentToolResponseAgentToolResponseStatus {
	return &e
}
