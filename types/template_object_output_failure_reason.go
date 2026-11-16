package types

import (
	"fmt"
)

type TemplateObjectOutputFailureReason string

const (
	TemplateObjectOutputFailureReasonTimeout           TemplateObjectOutputFailureReason = "timeout"
	TemplateObjectOutputFailureReasonModelError        TemplateObjectOutputFailureReason = "model_error"
	TemplateObjectOutputFailureReasonModerated         TemplateObjectOutputFailureReason = "moderated"
	TemplateObjectOutputFailureReasonInvalidParameters TemplateObjectOutputFailureReason = "invalid_parameters"
	TemplateObjectOutputFailureReasonDependencyFailed  TemplateObjectOutputFailureReason = "dependency_failed"
	TemplateObjectOutputFailureReasonChargingFailed    TemplateObjectOutputFailureReason = "charging_failed"
	TemplateObjectOutputFailureReasonInternalError     TemplateObjectOutputFailureReason = "internal_error"
)

func NewTemplateObjectOutputFailureReasonFromString(s string) (TemplateObjectOutputFailureReason, error) {
	switch s {
	case "timeout":
		return TemplateObjectOutputFailureReasonTimeout, nil
	case "model_error":
		return TemplateObjectOutputFailureReasonModelError, nil
	case "moderated":
		return TemplateObjectOutputFailureReasonModerated, nil
	case "invalid_parameters":
		return TemplateObjectOutputFailureReasonInvalidParameters, nil
	case "dependency_failed":
		return TemplateObjectOutputFailureReasonDependencyFailed, nil
	case "charging_failed":
		return TemplateObjectOutputFailureReasonChargingFailed, nil
	case "internal_error":
		return TemplateObjectOutputFailureReasonInternalError, nil
	}
	var t TemplateObjectOutputFailureReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TemplateObjectOutputFailureReason) Ptr() *TemplateObjectOutputFailureReason {
	return &e
}
