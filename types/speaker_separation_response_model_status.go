package types

import (
	"fmt"
)

// The status of the speaker separation.
type SpeakerSeparationResponseModelStatus string

const (
	SpeakerSeparationResponseModelStatusNotStarted SpeakerSeparationResponseModelStatus = "not_started"
	SpeakerSeparationResponseModelStatusPending    SpeakerSeparationResponseModelStatus = "pending"
	SpeakerSeparationResponseModelStatusCompleted  SpeakerSeparationResponseModelStatus = "completed"
	SpeakerSeparationResponseModelStatusFailed     SpeakerSeparationResponseModelStatus = "failed"
)

func NewSpeakerSeparationResponseModelStatusFromString(s string) (SpeakerSeparationResponseModelStatus, error) {
	switch s {
	case "not_started":
		return SpeakerSeparationResponseModelStatusNotStarted, nil
	case "pending":
		return SpeakerSeparationResponseModelStatusPending, nil
	case "completed":
		return SpeakerSeparationResponseModelStatusCompleted, nil
	case "failed":
		return SpeakerSeparationResponseModelStatusFailed, nil
	}
	var t SpeakerSeparationResponseModelStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SpeakerSeparationResponseModelStatus) Ptr() *SpeakerSeparationResponseModelStatus {
	return &e
}
