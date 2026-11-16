package types

import (
	"fmt"
)

type TemplateImageOutputFailureReason string

const (
	TemplateImageOutputFailureReasonTimeout           TemplateImageOutputFailureReason = "timeout"
	TemplateImageOutputFailureReasonModelError        TemplateImageOutputFailureReason = "model_error"
	TemplateImageOutputFailureReasonModerated         TemplateImageOutputFailureReason = "moderated"
	TemplateImageOutputFailureReasonInvalidParameters TemplateImageOutputFailureReason = "invalid_parameters"
	TemplateImageOutputFailureReasonDependencyFailed  TemplateImageOutputFailureReason = "dependency_failed"
	TemplateImageOutputFailureReasonChargingFailed    TemplateImageOutputFailureReason = "charging_failed"
	TemplateImageOutputFailureReasonInternalError     TemplateImageOutputFailureReason = "internal_error"
)

func NewTemplateImageOutputFailureReasonFromString(s string) (TemplateImageOutputFailureReason, error) {
	switch s {
	case "timeout":
		return TemplateImageOutputFailureReasonTimeout, nil
	case "model_error":
		return TemplateImageOutputFailureReasonModelError, nil
	case "moderated":
		return TemplateImageOutputFailureReasonModerated, nil
	case "invalid_parameters":
		return TemplateImageOutputFailureReasonInvalidParameters, nil
	case "dependency_failed":
		return TemplateImageOutputFailureReasonDependencyFailed, nil
	case "charging_failed":
		return TemplateImageOutputFailureReasonChargingFailed, nil
	case "internal_error":
		return TemplateImageOutputFailureReasonInternalError, nil
	}
	var t TemplateImageOutputFailureReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TemplateImageOutputFailureReason) Ptr() *TemplateImageOutputFailureReason {
	return &e
}
