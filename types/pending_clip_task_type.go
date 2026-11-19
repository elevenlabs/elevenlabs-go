package types

import (
	"fmt"
)

type PendingClipTaskType string

const (
	PendingClipTaskTypePreprocessing   PendingClipTaskType = "preprocessing"
	PendingClipTaskTypeSpeechImport    PendingClipTaskType = "speech_import"
	PendingClipTaskTypeDubbing         PendingClipTaskType = "dubbing"
	PendingClipTaskTypeVideoToMusic    PendingClipTaskType = "video_to_music"
	PendingClipTaskTypeMediaGeneration PendingClipTaskType = "media_generation"
)

func NewPendingClipTaskTypeFromString(s string) (PendingClipTaskType, error) {
	switch s {
	case "preprocessing":
		return PendingClipTaskTypePreprocessing, nil
	case "speech_import":
		return PendingClipTaskTypeSpeechImport, nil
	case "dubbing":
		return PendingClipTaskTypeDubbing, nil
	case "video_to_music":
		return PendingClipTaskTypeVideoToMusic, nil
	case "media_generation":
		return PendingClipTaskTypeMediaGeneration, nil
	}
	var t PendingClipTaskType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e PendingClipTaskType) Ptr() *PendingClipTaskType {
	return &e
}
