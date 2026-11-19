package types

import (
	"fmt"
)

type EvaluationResultFilter string

const (
	EvaluationResultFilterSuccess EvaluationResultFilter = "success"
	EvaluationResultFilterFailure EvaluationResultFilter = "failure"
	EvaluationResultFilterUnknown EvaluationResultFilter = "unknown"
	EvaluationResultFilterError   EvaluationResultFilter = "error"
)

func NewEvaluationResultFilterFromString(s string) (EvaluationResultFilter, error) {
	switch s {
	case "success":
		return EvaluationResultFilterSuccess, nil
	case "failure":
		return EvaluationResultFilterFailure, nil
	case "unknown":
		return EvaluationResultFilterUnknown, nil
	case "error":
		return EvaluationResultFilterError, nil
	}
	var t EvaluationResultFilter
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EvaluationResultFilter) Ptr() *EvaluationResultFilter {
	return &e
}
