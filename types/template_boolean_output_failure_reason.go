package types

import (
	"fmt"
)

type TemplateBooleanOutputFailureReason string

const (
	TemplateBooleanOutputFailureReasonTimeout           TemplateBooleanOutputFailureReason = "timeout"
	TemplateBooleanOutputFailureReasonModelError        TemplateBooleanOutputFailureReason = "model_error"
	TemplateBooleanOutputFailureReasonModerated         TemplateBooleanOutputFailureReason = "moderated"
	TemplateBooleanOutputFailureReasonInvalidParameters TemplateBooleanOutputFailureReason = "invalid_parameters"
	TemplateBooleanOutputFailureReasonDependencyFailed  TemplateBooleanOutputFailureReason = "dependency_failed"
	TemplateBooleanOutputFailureReasonChargingFailed    TemplateBooleanOutputFailureReason = "charging_failed"
	TemplateBooleanOutputFailureReasonInternalError     TemplateBooleanOutputFailureReason = "internal_error"
)

func NewTemplateBooleanOutputFailureReasonFromString(s string) (TemplateBooleanOutputFailureReason, error) {
	switch s {
	case "timeout":
		return TemplateBooleanOutputFailureReasonTimeout, nil
	case "model_error":
		return TemplateBooleanOutputFailureReasonModelError, nil
	case "moderated":
		return TemplateBooleanOutputFailureReasonModerated, nil
	case "invalid_parameters":
		return TemplateBooleanOutputFailureReasonInvalidParameters, nil
	case "dependency_failed":
		return TemplateBooleanOutputFailureReasonDependencyFailed, nil
	case "charging_failed":
		return TemplateBooleanOutputFailureReasonChargingFailed, nil
	case "internal_error":
		return TemplateBooleanOutputFailureReasonInternalError, nil
	}
	var t TemplateBooleanOutputFailureReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TemplateBooleanOutputFailureReason) Ptr() *TemplateBooleanOutputFailureReason {
	return &e
}
