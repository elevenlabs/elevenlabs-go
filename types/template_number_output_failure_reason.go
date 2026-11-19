package types

import (
	"fmt"
)

type TemplateNumberOutputFailureReason string

const (
	TemplateNumberOutputFailureReasonTimeout           TemplateNumberOutputFailureReason = "timeout"
	TemplateNumberOutputFailureReasonModelError        TemplateNumberOutputFailureReason = "model_error"
	TemplateNumberOutputFailureReasonModerated         TemplateNumberOutputFailureReason = "moderated"
	TemplateNumberOutputFailureReasonInvalidParameters TemplateNumberOutputFailureReason = "invalid_parameters"
	TemplateNumberOutputFailureReasonDependencyFailed  TemplateNumberOutputFailureReason = "dependency_failed"
	TemplateNumberOutputFailureReasonChargingFailed    TemplateNumberOutputFailureReason = "charging_failed"
	TemplateNumberOutputFailureReasonInternalError     TemplateNumberOutputFailureReason = "internal_error"
)

func NewTemplateNumberOutputFailureReasonFromString(s string) (TemplateNumberOutputFailureReason, error) {
	switch s {
	case "timeout":
		return TemplateNumberOutputFailureReasonTimeout, nil
	case "model_error":
		return TemplateNumberOutputFailureReasonModelError, nil
	case "moderated":
		return TemplateNumberOutputFailureReasonModerated, nil
	case "invalid_parameters":
		return TemplateNumberOutputFailureReasonInvalidParameters, nil
	case "dependency_failed":
		return TemplateNumberOutputFailureReasonDependencyFailed, nil
	case "charging_failed":
		return TemplateNumberOutputFailureReasonChargingFailed, nil
	case "internal_error":
		return TemplateNumberOutputFailureReasonInternalError, nil
	}
	var t TemplateNumberOutputFailureReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TemplateNumberOutputFailureReason) Ptr() *TemplateNumberOutputFailureReason {
	return &e
}
