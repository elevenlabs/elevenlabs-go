package types

import (
	"fmt"
)

// Whether to wait for the reply, or keep going and get it later in the conversation.
type SubAgentOutputExecutionMode string

const (
	SubAgentOutputExecutionModeImmediate SubAgentOutputExecutionMode = "immediate"
	SubAgentOutputExecutionModeAsync     SubAgentOutputExecutionMode = "async"
)

func NewSubAgentOutputExecutionModeFromString(s string) (SubAgentOutputExecutionMode, error) {
	switch s {
	case "immediate":
		return SubAgentOutputExecutionModeImmediate, nil
	case "async":
		return SubAgentOutputExecutionModeAsync, nil
	}
	var t SubAgentOutputExecutionMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SubAgentOutputExecutionMode) Ptr() *SubAgentOutputExecutionMode {
	return &e
}
