package types

import (
	"fmt"
)

// “BURSTY“ clients stream ahead of wall time, so decisions use the audio clock.
type AudioInputPacing string

const (
	AudioInputPacingRealtime AudioInputPacing = "realtime"
	AudioInputPacingBursty   AudioInputPacing = "bursty"
)

func NewAudioInputPacingFromString(s string) (AudioInputPacing, error) {
	switch s {
	case "realtime":
		return AudioInputPacingRealtime, nil
	case "bursty":
		return AudioInputPacingBursty, nil
	}
	var t AudioInputPacing
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AudioInputPacing) Ptr() *AudioInputPacing {
	return &e
}
