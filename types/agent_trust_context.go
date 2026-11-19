package types

import (
	"fmt"
)

// The trust context in which the agent operates.
//
// UNKNOWN: not yet classified (existing agents created before this feature).
// LOW: serves untrusted external participants (e.g. customer support, sales) —
// outputs should be vetted and tool access scoped.
// HIGH: serves the owner (e.g. personal assistant) — full tool access is appropriate.
type AgentTrustContext string

const (
	AgentTrustContextUnknown AgentTrustContext = "unknown"
	AgentTrustContextLow     AgentTrustContext = "low"
	AgentTrustContextHigh    AgentTrustContext = "high"
)

func NewAgentTrustContextFromString(s string) (AgentTrustContext, error) {
	switch s {
	case "unknown":
		return AgentTrustContextUnknown, nil
	case "low":
		return AgentTrustContextLow, nil
	case "high":
		return AgentTrustContextHigh, nil
	}
	var t AgentTrustContext
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AgentTrustContext) Ptr() *AgentTrustContext {
	return &e
}
