package types

import (
	"fmt"
)

type MusicFinetuneStatus string

const (
	MusicFinetuneStatusPending    MusicFinetuneStatus = "pending"
	MusicFinetuneStatusInProgress MusicFinetuneStatus = "in_progress"
	MusicFinetuneStatusCompleted  MusicFinetuneStatus = "completed"
	MusicFinetuneStatusFailed     MusicFinetuneStatus = "failed"
	MusicFinetuneStatusBlocked    MusicFinetuneStatus = "blocked"
)

func NewMusicFinetuneStatusFromString(s string) (MusicFinetuneStatus, error) {
	switch s {
	case "pending":
		return MusicFinetuneStatusPending, nil
	case "in_progress":
		return MusicFinetuneStatusInProgress, nil
	case "completed":
		return MusicFinetuneStatusCompleted, nil
	case "failed":
		return MusicFinetuneStatusFailed, nil
	case "blocked":
		return MusicFinetuneStatusBlocked, nil
	}
	var t MusicFinetuneStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e MusicFinetuneStatus) Ptr() *MusicFinetuneStatus {
	return &e
}
