package types

import (
	"fmt"
)

type TemplateArrayOutputFailureReason string

const (
	TemplateArrayOutputFailureReasonTimeout           TemplateArrayOutputFailureReason = "timeout"
	TemplateArrayOutputFailureReasonModelError        TemplateArrayOutputFailureReason = "model_error"
	TemplateArrayOutputFailureReasonModerated         TemplateArrayOutputFailureReason = "moderated"
	TemplateArrayOutputFailureReasonInvalidParameters TemplateArrayOutputFailureReason = "invalid_parameters"
	TemplateArrayOutputFailureReasonDependencyFailed  TemplateArrayOutputFailureReason = "dependency_failed"
	TemplateArrayOutputFailureReasonChargingFailed    TemplateArrayOutputFailureReason = "charging_failed"
	TemplateArrayOutputFailureReasonInternalError     TemplateArrayOutputFailureReason = "internal_error"
)

func NewTemplateArrayOutputFailureReasonFromString(s string) (TemplateArrayOutputFailureReason, error) {
	switch s {
	case "timeout":
		return TemplateArrayOutputFailureReasonTimeout, nil
	case "model_error":
		return TemplateArrayOutputFailureReasonModelError, nil
	case "moderated":
		return TemplateArrayOutputFailureReasonModerated, nil
	case "invalid_parameters":
		return TemplateArrayOutputFailureReasonInvalidParameters, nil
	case "dependency_failed":
		return TemplateArrayOutputFailureReasonDependencyFailed, nil
	case "charging_failed":
		return TemplateArrayOutputFailureReasonChargingFailed, nil
	case "internal_error":
		return TemplateArrayOutputFailureReasonInternalError, nil
	}
	var t TemplateArrayOutputFailureReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TemplateArrayOutputFailureReason) Ptr() *TemplateArrayOutputFailureReason {
	return &e
}
