package types

import (
	"fmt"
)

type EvaluationSuccessResult string

const (
	EvaluationSuccessResultSuccess EvaluationSuccessResult = "success"
	EvaluationSuccessResultFailure EvaluationSuccessResult = "failure"
	EvaluationSuccessResultUnknown EvaluationSuccessResult = "unknown"
)

func NewEvaluationSuccessResultFromString(s string) (EvaluationSuccessResult, error) {
	switch s {
	case "success":
		return EvaluationSuccessResultSuccess, nil
	case "failure":
		return EvaluationSuccessResultFailure, nil
	case "unknown":
		return EvaluationSuccessResultUnknown, nil
	}
	var t EvaluationSuccessResult
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e EvaluationSuccessResult) Ptr() *EvaluationSuccessResult {
	return &e
}
