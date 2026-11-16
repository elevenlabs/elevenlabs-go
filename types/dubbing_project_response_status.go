package types

import (
	"fmt"
)

// Lifecycle status of the project: `queued` before the source is picked up, `preparing` while it is transcribed, `ready` once transcription is done and language targets can start, or `failed`. A project is never reported as `processing` — that value belongs to language targets.
type DubbingProjectResponseStatus string

const (
	DubbingProjectResponseStatusQueued     DubbingProjectResponseStatus = "queued"
	DubbingProjectResponseStatusPreparing  DubbingProjectResponseStatus = "preparing"
	DubbingProjectResponseStatusProcessing DubbingProjectResponseStatus = "processing"
	DubbingProjectResponseStatusReady      DubbingProjectResponseStatus = "ready"
	DubbingProjectResponseStatusFailed     DubbingProjectResponseStatus = "failed"
)

func NewDubbingProjectResponseStatusFromString(s string) (DubbingProjectResponseStatus, error) {
	switch s {
	case "queued":
		return DubbingProjectResponseStatusQueued, nil
	case "preparing":
		return DubbingProjectResponseStatusPreparing, nil
	case "processing":
		return DubbingProjectResponseStatusProcessing, nil
	case "ready":
		return DubbingProjectResponseStatusReady, nil
	case "failed":
		return DubbingProjectResponseStatusFailed, nil
	}
	var t DubbingProjectResponseStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DubbingProjectResponseStatus) Ptr() *DubbingProjectResponseStatus {
	return &e
}
