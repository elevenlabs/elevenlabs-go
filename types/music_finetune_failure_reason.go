package types

import (
	"fmt"
)

type MusicFinetuneFailureReason string

const (
	MusicFinetuneFailureReasonAudioProcessingFailed MusicFinetuneFailureReason = "audio_processing_failed"
	MusicFinetuneFailureReasonCopyrightViolation    MusicFinetuneFailureReason = "copyright_violation"
	MusicFinetuneFailureReasonTrainingFailed        MusicFinetuneFailureReason = "training_failed"
)

func NewMusicFinetuneFailureReasonFromString(s string) (MusicFinetuneFailureReason, error) {
	switch s {
	case "audio_processing_failed":
		return MusicFinetuneFailureReasonAudioProcessingFailed, nil
	case "copyright_violation":
		return MusicFinetuneFailureReasonCopyrightViolation, nil
	case "training_failed":
		return MusicFinetuneFailureReasonTrainingFailed, nil
	}
	var t MusicFinetuneFailureReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MusicFinetuneFailureReason) Ptr() *MusicFinetuneFailureReason {
	return &e
}
