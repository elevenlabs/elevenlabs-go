package types

import (
	"fmt"
)

type GuardrailExecutionMode string

const (
	GuardrailExecutionModeStreaming GuardrailExecutionMode = "streaming"
	GuardrailExecutionModeBlocking  GuardrailExecutionMode = "blocking"
)

func NewGuardrailExecutionModeFromString(s string) (GuardrailExecutionMode, error) {
	switch s {
	case "streaming":
		return GuardrailExecutionModeStreaming, nil
	case "blocking":
		return GuardrailExecutionModeBlocking, nil
	}
	var t GuardrailExecutionMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GuardrailExecutionMode) Ptr() *GuardrailExecutionMode {
	return &e
}
