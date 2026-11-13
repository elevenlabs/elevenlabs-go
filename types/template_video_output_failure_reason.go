package types

import (
	"fmt"
)

type TemplateVideoOutputFailureReason string

const (
	TemplateVideoOutputFailureReasonTimeout           TemplateVideoOutputFailureReason = "timeout"
	TemplateVideoOutputFailureReasonModelError        TemplateVideoOutputFailureReason = "model_error"
	TemplateVideoOutputFailureReasonModerated         TemplateVideoOutputFailureReason = "moderated"
	TemplateVideoOutputFailureReasonInvalidParameters TemplateVideoOutputFailureReason = "invalid_parameters"
	TemplateVideoOutputFailureReasonDependencyFailed  TemplateVideoOutputFailureReason = "dependency_failed"
	TemplateVideoOutputFailureReasonChargingFailed    TemplateVideoOutputFailureReason = "charging_failed"
	TemplateVideoOutputFailureReasonInternalError     TemplateVideoOutputFailureReason = "internal_error"
)

func NewTemplateVideoOutputFailureReasonFromString(s string) (TemplateVideoOutputFailureReason, error) {
	switch s {
	case "timeout":
		return TemplateVideoOutputFailureReasonTimeout, nil
	case "model_error":
		return TemplateVideoOutputFailureReasonModelError, nil
	case "moderated":
		return TemplateVideoOutputFailureReasonModerated, nil
	case "invalid_parameters":
		return TemplateVideoOutputFailureReasonInvalidParameters, nil
	case "dependency_failed":
		return TemplateVideoOutputFailureReasonDependencyFailed, nil
	case "charging_failed":
		return TemplateVideoOutputFailureReasonChargingFailed, nil
	case "internal_error":
		return TemplateVideoOutputFailureReasonInternalError, nil
	}
	var t TemplateVideoOutputFailureReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TemplateVideoOutputFailureReason) Ptr() *TemplateVideoOutputFailureReason {
	return &e
}
