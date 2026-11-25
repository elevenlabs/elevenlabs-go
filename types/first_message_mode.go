package types

import (
	"fmt"
)

// How the receptionist's greeting plays.
type FirstMessageMode string

const (
	FirstMessageModeInterruptible   FirstMessageMode = "interruptible"
	FirstMessageModeUninterruptible FirstMessageMode = "uninterruptible"
	FirstMessageModeRecordingNotice FirstMessageMode = "recording_notice"
)

func NewFirstMessageModeFromString(s string) (FirstMessageMode, error) {
	switch s {
	case "interruptible":
		return FirstMessageModeInterruptible, nil
	case "uninterruptible":
		return FirstMessageModeUninterruptible, nil
	case "recording_notice":
		return FirstMessageModeRecordingNotice, nil
	}
	var t FirstMessageMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e FirstMessageMode) Ptr() *FirstMessageMode {
	return &e
}
