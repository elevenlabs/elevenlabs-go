package types

import (
	"fmt"
)

// How strictly the receptionist follows a rule.
type RulePriority string

const (
	RulePriorityStrict     RulePriority = "strict"
	RulePrioritySuggestion RulePriority = "suggestion"
)

func NewRulePriorityFromString(s string) (RulePriority, error) {
	switch s {
	case "strict":
		return RulePriorityStrict, nil
	case "suggestion":
		return RulePrioritySuggestion, nil
	}
	var t RulePriority
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e RulePriority) Ptr() *RulePriority {
	return &e
}
