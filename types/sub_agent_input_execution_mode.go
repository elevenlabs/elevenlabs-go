package types

import (
	"fmt"
)

// Whether to wait for the reply, or keep going and get it later in the conversation.
type SubAgentInputExecutionMode string

const (
	SubAgentInputExecutionModeImmediate SubAgentInputExecutionMode = "immediate"
	SubAgentInputExecutionModeAsync     SubAgentInputExecutionMode = "async"
)

func NewSubAgentInputExecutionModeFromString(s string) (SubAgentInputExecutionMode, error) {
	switch s {
	case "immediate":
		return SubAgentInputExecutionModeImmediate, nil
	case "async":
		return SubAgentInputExecutionModeAsync, nil
	}
	var t SubAgentInputExecutionMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SubAgentInputExecutionMode) Ptr() *SubAgentInputExecutionMode {
	return &e
}
