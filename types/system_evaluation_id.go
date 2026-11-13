package types

import (
	"fmt"
)

type SystemEvaluationId string

const (
	SystemEvaluationIdSystemEvalCriteriaSentiment   SystemEvaluationId = "__system_eval_criteria_sentiment"
	SystemEvaluationIdSystemEvalCriteriaFrustration SystemEvaluationId = "__system_eval_criteria_frustration"
)

func NewSystemEvaluationIdFromString(s string) (SystemEvaluationId, error) {
	switch s {
	case "__system_eval_criteria_sentiment":
		return SystemEvaluationIdSystemEvalCriteriaSentiment, nil
	case "__system_eval_criteria_frustration":
		return SystemEvaluationIdSystemEvalCriteriaFrustration, nil
	}
	var t SystemEvaluationId
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SystemEvaluationId) Ptr() *SystemEvaluationId {
	return &e
}
