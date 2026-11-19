package types

import (
	"fmt"
)

type ImageAnalysisStatus string

const (
	ImageAnalysisStatusProcessing ImageAnalysisStatus = "processing"
	ImageAnalysisStatusCompleted  ImageAnalysisStatus = "completed"
	ImageAnalysisStatusFailed     ImageAnalysisStatus = "failed"
)

func NewImageAnalysisStatusFromString(s string) (ImageAnalysisStatus, error) {
	switch s {
	case "processing":
		return ImageAnalysisStatusProcessing, nil
	case "completed":
		return ImageAnalysisStatusCompleted, nil
	case "failed":
		return ImageAnalysisStatusFailed, nil
	}
	var t ImageAnalysisStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ImageAnalysisStatus) Ptr() *ImageAnalysisStatus {
	return &e
}
