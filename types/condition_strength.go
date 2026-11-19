package types

import (
	"fmt"
)

type ConditionStrength string

const (
	ConditionStrengthLow    ConditionStrength = "low"
	ConditionStrengthMedium ConditionStrength = "medium"
	ConditionStrengthHigh   ConditionStrength = "high"
	ConditionStrengthXhigh  ConditionStrength = "xhigh"
)

func NewConditionStrengthFromString(s string) (ConditionStrength, error) {
	switch s {
	case "low":
		return ConditionStrengthLow, nil
	case "medium":
		return ConditionStrengthMedium, nil
	case "high":
		return ConditionStrengthHigh, nil
	case "xhigh":
		return ConditionStrengthXhigh, nil
	}
	var t ConditionStrength
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ConditionStrength) Ptr() *ConditionStrength {
	return &e
}
