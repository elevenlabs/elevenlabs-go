package types

import (
	"fmt"
)

type DubbingStatusStatus string

const (
	DubbingStatusStatusStarted DubbingStatusStatus = "started"
	DubbingStatusStatusStopped DubbingStatusStatus = "stopped"
)

func NewDubbingStatusStatusFromString(s string) (DubbingStatusStatus, error) {
	switch s {
	case "started":
		return DubbingStatusStatusStarted, nil
	case "stopped":
		return DubbingStatusStatusStopped, nil
	}
	var t DubbingStatusStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e DubbingStatusStatus) Ptr() *DubbingStatusStatus {
	return &e
}
