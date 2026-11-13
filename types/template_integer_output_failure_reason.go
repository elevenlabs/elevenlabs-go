package types

import (
	"fmt"
)

type TemplateIntegerOutputFailureReason string

const (
	TemplateIntegerOutputFailureReasonTimeout           TemplateIntegerOutputFailureReason = "timeout"
	TemplateIntegerOutputFailureReasonModelError        TemplateIntegerOutputFailureReason = "model_error"
	TemplateIntegerOutputFailureReasonModerated         TemplateIntegerOutputFailureReason = "moderated"
	TemplateIntegerOutputFailureReasonInvalidParameters TemplateIntegerOutputFailureReason = "invalid_parameters"
	TemplateIntegerOutputFailureReasonDependencyFailed  TemplateIntegerOutputFailureReason = "dependency_failed"
	TemplateIntegerOutputFailureReasonChargingFailed    TemplateIntegerOutputFailureReason = "charging_failed"
	TemplateIntegerOutputFailureReasonInternalError     TemplateIntegerOutputFailureReason = "internal_error"
)

func NewTemplateIntegerOutputFailureReasonFromString(s string) (TemplateIntegerOutputFailureReason, error) {
	switch s {
	case "timeout":
		return TemplateIntegerOutputFailureReasonTimeout, nil
	case "model_error":
		return TemplateIntegerOutputFailureReasonModelError, nil
	case "moderated":
		return TemplateIntegerOutputFailureReasonModerated, nil
	case "invalid_parameters":
		return TemplateIntegerOutputFailureReasonInvalidParameters, nil
	case "dependency_failed":
		return TemplateIntegerOutputFailureReasonDependencyFailed, nil
	case "charging_failed":
		return TemplateIntegerOutputFailureReasonChargingFailed, nil
	case "internal_error":
		return TemplateIntegerOutputFailureReasonInternalError, nil
	}
	var t TemplateIntegerOutputFailureReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TemplateIntegerOutputFailureReason) Ptr() *TemplateIntegerOutputFailureReason {
	return &e
}
