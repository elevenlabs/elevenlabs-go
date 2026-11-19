package types

import (
	"fmt"
)

// Lifecycle status: `queued` (waiting on the project to be ready, or on a worker), `processing` while it is being dubbed, `completed` once its output is available, `stale` when the transcript changed after the output was produced, or `failed`.
type DubbingLanguageResponseStatus string

const (
	DubbingLanguageResponseStatusQueued     DubbingLanguageResponseStatus = "queued"
	DubbingLanguageResponseStatusProcessing DubbingLanguageResponseStatus = "processing"
	DubbingLanguageResponseStatusCompleted  DubbingLanguageResponseStatus = "completed"
	DubbingLanguageResponseStatusStale      DubbingLanguageResponseStatus = "stale"
	DubbingLanguageResponseStatusFailed     DubbingLanguageResponseStatus = "failed"
)

func NewDubbingLanguageResponseStatusFromString(s string) (DubbingLanguageResponseStatus, error) {
	switch s {
	case "queued":
		return DubbingLanguageResponseStatusQueued, nil
	case "processing":
		return DubbingLanguageResponseStatusProcessing, nil
	case "completed":
		return DubbingLanguageResponseStatusCompleted, nil
	case "stale":
		return DubbingLanguageResponseStatusStale, nil
	case "failed":
		return DubbingLanguageResponseStatusFailed, nil
	}
	var t DubbingLanguageResponseStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DubbingLanguageResponseStatus) Ptr() *DubbingLanguageResponseStatus {
	return &e
}
