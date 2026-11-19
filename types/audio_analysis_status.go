package types

import (
	"fmt"
)

type AudioAnalysisStatus string

const (
	AudioAnalysisStatusProcessing AudioAnalysisStatus = "processing"
	AudioAnalysisStatusCompleted  AudioAnalysisStatus = "completed"
	AudioAnalysisStatusFailed     AudioAnalysisStatus = "failed"
)

func NewAudioAnalysisStatusFromString(s string) (AudioAnalysisStatus, error) {
	switch s {
	case "processing":
		return AudioAnalysisStatusProcessing, nil
	case "completed":
		return AudioAnalysisStatusCompleted, nil
	case "failed":
		return AudioAnalysisStatusFailed, nil
	}
	var t AudioAnalysisStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AudioAnalysisStatus) Ptr() *AudioAnalysisStatus {
	return &e
}
