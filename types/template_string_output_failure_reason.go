package types

import (
	"fmt"
)

type TemplateStringOutputFailureReason string

const (
	TemplateStringOutputFailureReasonTimeout           TemplateStringOutputFailureReason = "timeout"
	TemplateStringOutputFailureReasonModelError        TemplateStringOutputFailureReason = "model_error"
	TemplateStringOutputFailureReasonModerated         TemplateStringOutputFailureReason = "moderated"
	TemplateStringOutputFailureReasonInvalidParameters TemplateStringOutputFailureReason = "invalid_parameters"
	TemplateStringOutputFailureReasonDependencyFailed  TemplateStringOutputFailureReason = "dependency_failed"
	TemplateStringOutputFailureReasonChargingFailed    TemplateStringOutputFailureReason = "charging_failed"
	TemplateStringOutputFailureReasonInternalError     TemplateStringOutputFailureReason = "internal_error"
)

func NewTemplateStringOutputFailureReasonFromString(s string) (TemplateStringOutputFailureReason, error) {
	switch s {
	case "timeout":
		return TemplateStringOutputFailureReasonTimeout, nil
	case "model_error":
		return TemplateStringOutputFailureReasonModelError, nil
	case "moderated":
		return TemplateStringOutputFailureReasonModerated, nil
	case "invalid_parameters":
		return TemplateStringOutputFailureReasonInvalidParameters, nil
	case "dependency_failed":
		return TemplateStringOutputFailureReasonDependencyFailed, nil
	case "charging_failed":
		return TemplateStringOutputFailureReasonChargingFailed, nil
	case "internal_error":
		return TemplateStringOutputFailureReasonInternalError, nil
	}
	var t TemplateStringOutputFailureReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TemplateStringOutputFailureReason) Ptr() *TemplateStringOutputFailureReason {
	return &e
}
