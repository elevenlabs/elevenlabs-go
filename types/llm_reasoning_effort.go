package types

import (
	"fmt"
)

type LlmReasoningEffort string

const (
	LlmReasoningEffortNone    LlmReasoningEffort = "none"
	LlmReasoningEffortMinimal LlmReasoningEffort = "minimal"
	LlmReasoningEffortLow     LlmReasoningEffort = "low"
	LlmReasoningEffortMedium  LlmReasoningEffort = "medium"
	LlmReasoningEffortHigh    LlmReasoningEffort = "high"
	LlmReasoningEffortXhigh   LlmReasoningEffort = "xhigh"
	LlmReasoningEffortMax     LlmReasoningEffort = "max"
)

func NewLlmReasoningEffortFromString(s string) (LlmReasoningEffort, error) {
	switch s {
	case "none":
		return LlmReasoningEffortNone, nil
	case "minimal":
		return LlmReasoningEffortMinimal, nil
	case "low":
		return LlmReasoningEffortLow, nil
	case "medium":
		return LlmReasoningEffortMedium, nil
	case "high":
		return LlmReasoningEffortHigh, nil
	case "xhigh":
		return LlmReasoningEffortXhigh, nil
	case "max":
		return LlmReasoningEffortMax, nil
	}
	var t LlmReasoningEffort
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e LlmReasoningEffort) Ptr() *LlmReasoningEffort {
	return &e
}
