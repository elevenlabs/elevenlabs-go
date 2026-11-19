package types

import (
	"fmt"
)

type AgentToolResponseStatus string

const (
	AgentToolResponseStatusSuccess AgentToolResponseStatus = "success"
	AgentToolResponseStatusError   AgentToolResponseStatus = "error"
	AgentToolResponseStatusBlocked AgentToolResponseStatus = "blocked"
	AgentToolResponseStatusSkipped AgentToolResponseStatus = "skipped"
)

func NewAgentToolResponseStatusFromString(s string) (AgentToolResponseStatus, error) {
	switch s {
	case "success":
		return AgentToolResponseStatusSuccess, nil
	case "error":
		return AgentToolResponseStatusError, nil
	case "blocked":
		return AgentToolResponseStatusBlocked, nil
	case "skipped":
		return AgentToolResponseStatusSkipped, nil
	}
	var t AgentToolResponseStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentToolResponseStatus) Ptr() *AgentToolResponseStatus {
	return &e
}
