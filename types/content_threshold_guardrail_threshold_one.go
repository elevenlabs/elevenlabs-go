package types

import (
	"fmt"
)

type ContentThresholdGuardrailThresholdOne string

const (
	ContentThresholdGuardrailThresholdOneLow    ContentThresholdGuardrailThresholdOne = "low"
	ContentThresholdGuardrailThresholdOneMedium ContentThresholdGuardrailThresholdOne = "medium"
	ContentThresholdGuardrailThresholdOneHigh   ContentThresholdGuardrailThresholdOne = "high"
)

func NewContentThresholdGuardrailThresholdOneFromString(s string) (ContentThresholdGuardrailThresholdOne, error) {
	switch s {
	case "low":
		return ContentThresholdGuardrailThresholdOneLow, nil
	case "medium":
		return ContentThresholdGuardrailThresholdOneMedium, nil
	case "high":
		return ContentThresholdGuardrailThresholdOneHigh, nil
	}
	var t ContentThresholdGuardrailThresholdOne
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ContentThresholdGuardrailThresholdOne) Ptr() *ContentThresholdGuardrailThresholdOne {
	return &e
}
