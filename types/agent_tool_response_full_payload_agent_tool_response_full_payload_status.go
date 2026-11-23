package types

import (
	"fmt"
)

// Tool call status derived from execution flags.
type AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatus string

const (
	AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatusSuccess AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatus = "success"
	AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatusError   AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatus = "error"
	AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatusBlocked AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatus = "blocked"
	AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatusSkipped AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatus = "skipped"
)

func NewAgentToolResponseFullPayloadAgentToolResponseFullPayloadStatusFromString(s string) (AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatus, error) {
	switch s {
	case "success":
		return AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatusSuccess, nil
	case "error":
		return AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatusError, nil
	case "blocked":
		return AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatusBlocked, nil
	case "skipped":
		return AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatusSkipped, nil
	}
	var t AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatus) Ptr() *AgentToolResponseFullPayloadAgentToolResponseFullPayloadStatus {
	return &e
}
