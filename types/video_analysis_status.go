package types

import (
	"fmt"
)

type VideoAnalysisStatus string

const (
	VideoAnalysisStatusProcessing VideoAnalysisStatus = "processing"
	VideoAnalysisStatusCompleted  VideoAnalysisStatus = "completed"
	VideoAnalysisStatusFailed     VideoAnalysisStatus = "failed"
)

func NewVideoAnalysisStatusFromString(s string) (VideoAnalysisStatus, error) {
	switch s {
	case "processing":
		return VideoAnalysisStatusProcessing, nil
	case "completed":
		return VideoAnalysisStatusCompleted, nil
	case "failed":
		return VideoAnalysisStatusFailed, nil
	}
	var t VideoAnalysisStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e VideoAnalysisStatus) Ptr() *VideoAnalysisStatus {
	return &e
}
