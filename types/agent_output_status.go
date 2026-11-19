package types

import (
	"fmt"
)

type AgentOutputStatus string

const (
	AgentOutputStatusInProgress     AgentOutputStatus = "in_progress"
	AgentOutputStatusCompleted      AgentOutputStatus = "completed"
	AgentOutputStatusRequiresAction AgentOutputStatus = "requires_action"
	AgentOutputStatusFailed         AgentOutputStatus = "failed"
	AgentOutputStatusTimedOut       AgentOutputStatus = "timed_out"
)

func NewAgentOutputStatusFromString(s string) (AgentOutputStatus, error) {
	switch s {
	case "in_progress":
		return AgentOutputStatusInProgress, nil
	case "completed":
		return AgentOutputStatusCompleted, nil
	case "requires_action":
		return AgentOutputStatusRequiresAction, nil
	case "failed":
		return AgentOutputStatusFailed, nil
	case "timed_out":
		return AgentOutputStatusTimedOut, nil
	}
	var t AgentOutputStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentOutputStatus) Ptr() *AgentOutputStatus {
	return &e
}
