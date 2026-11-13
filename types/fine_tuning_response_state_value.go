package types

import (
	"fmt"
)

type FineTuningResponseStateValue string

const (
	FineTuningResponseStateValueNotStarted FineTuningResponseStateValue = "not_started"
	FineTuningResponseStateValueQueued     FineTuningResponseStateValue = "queued"
	FineTuningResponseStateValueFineTuning FineTuningResponseStateValue = "fine_tuning"
	FineTuningResponseStateValueFineTuned  FineTuningResponseStateValue = "fine_tuned"
	FineTuningResponseStateValueFailed     FineTuningResponseStateValue = "failed"
	FineTuningResponseStateValueDelayed    FineTuningResponseStateValue = "delayed"
)

func NewFineTuningResponseStateValueFromString(s string) (FineTuningResponseStateValue, error) {
	switch s {
	case "not_started":
		return FineTuningResponseStateValueNotStarted, nil
	case "queued":
		return FineTuningResponseStateValueQueued, nil
	case "fine_tuning":
		return FineTuningResponseStateValueFineTuning, nil
	case "fine_tuned":
		return FineTuningResponseStateValueFineTuned, nil
	case "failed":
		return FineTuningResponseStateValueFailed, nil
	case "delayed":
		return FineTuningResponseStateValueDelayed, nil
	}
	var t FineTuningResponseStateValue
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e FineTuningResponseStateValue) Ptr() *FineTuningResponseStateValue {
	return &e
}
