package types

import (
	"fmt"
)

type TemplateAudioOutputFailureReason string

const (
	TemplateAudioOutputFailureReasonTimeout           TemplateAudioOutputFailureReason = "timeout"
	TemplateAudioOutputFailureReasonModelError        TemplateAudioOutputFailureReason = "model_error"
	TemplateAudioOutputFailureReasonModerated         TemplateAudioOutputFailureReason = "moderated"
	TemplateAudioOutputFailureReasonInvalidParameters TemplateAudioOutputFailureReason = "invalid_parameters"
	TemplateAudioOutputFailureReasonDependencyFailed  TemplateAudioOutputFailureReason = "dependency_failed"
	TemplateAudioOutputFailureReasonChargingFailed    TemplateAudioOutputFailureReason = "charging_failed"
	TemplateAudioOutputFailureReasonInternalError     TemplateAudioOutputFailureReason = "internal_error"
)

func NewTemplateAudioOutputFailureReasonFromString(s string) (TemplateAudioOutputFailureReason, error) {
	switch s {
	case "timeout":
		return TemplateAudioOutputFailureReasonTimeout, nil
	case "model_error":
		return TemplateAudioOutputFailureReasonModelError, nil
	case "moderated":
		return TemplateAudioOutputFailureReasonModerated, nil
	case "invalid_parameters":
		return TemplateAudioOutputFailureReasonInvalidParameters, nil
	case "dependency_failed":
		return TemplateAudioOutputFailureReasonDependencyFailed, nil
	case "charging_failed":
		return TemplateAudioOutputFailureReasonChargingFailed, nil
	case "internal_error":
		return TemplateAudioOutputFailureReasonInternalError, nil
	}
	var t TemplateAudioOutputFailureReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TemplateAudioOutputFailureReason) Ptr() *TemplateAudioOutputFailureReason {
	return &e
}
