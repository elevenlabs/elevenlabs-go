package types

import (
	"fmt"
)

// The category of failure.
type MediaGenerationFailedResponseFailureReason string

const (
	MediaGenerationFailedResponseFailureReasonTimeout           MediaGenerationFailedResponseFailureReason = "timeout"
	MediaGenerationFailedResponseFailureReasonModelError        MediaGenerationFailedResponseFailureReason = "model_error"
	MediaGenerationFailedResponseFailureReasonModerated         MediaGenerationFailedResponseFailureReason = "moderated"
	MediaGenerationFailedResponseFailureReasonInvalidParameters MediaGenerationFailedResponseFailureReason = "invalid_parameters"
	MediaGenerationFailedResponseFailureReasonDependencyFailed  MediaGenerationFailedResponseFailureReason = "dependency_failed"
	MediaGenerationFailedResponseFailureReasonChargingFailed    MediaGenerationFailedResponseFailureReason = "charging_failed"
	MediaGenerationFailedResponseFailureReasonInternalError     MediaGenerationFailedResponseFailureReason = "internal_error"
)

func NewMediaGenerationFailedResponseFailureReasonFromString(s string) (MediaGenerationFailedResponseFailureReason, error) {
	switch s {
	case "timeout":
		return MediaGenerationFailedResponseFailureReasonTimeout, nil
	case "model_error":
		return MediaGenerationFailedResponseFailureReasonModelError, nil
	case "moderated":
		return MediaGenerationFailedResponseFailureReasonModerated, nil
	case "invalid_parameters":
		return MediaGenerationFailedResponseFailureReasonInvalidParameters, nil
	case "dependency_failed":
		return MediaGenerationFailedResponseFailureReasonDependencyFailed, nil
	case "charging_failed":
		return MediaGenerationFailedResponseFailureReasonChargingFailed, nil
	case "internal_error":
		return MediaGenerationFailedResponseFailureReasonInternalError, nil
	}
	var t MediaGenerationFailedResponseFailureReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MediaGenerationFailedResponseFailureReason) Ptr() *MediaGenerationFailedResponseFailureReason {
	return &e
}
