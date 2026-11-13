package types

import (
	"fmt"
)

type TranscriptPlatformEvent string

const (
	TranscriptPlatformEventUserTurnTimeout     TranscriptPlatformEvent = "user_turn_timeout"
	TranscriptPlatformEventSkipTurnWaitElapsed TranscriptPlatformEvent = "skip_turn_wait_elapsed"
)

func NewTranscriptPlatformEventFromString(s string) (TranscriptPlatformEvent, error) {
	switch s {
	case "user_turn_timeout":
		return TranscriptPlatformEventUserTurnTimeout, nil
	case "skip_turn_wait_elapsed":
		return TranscriptPlatformEventSkipTurnWaitElapsed, nil
	}
	var t TranscriptPlatformEvent
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TranscriptPlatformEvent) Ptr() *TranscriptPlatformEvent {
	return &e
}
