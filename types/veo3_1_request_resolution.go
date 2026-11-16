package types

import (
	"fmt"
)

// The resolution of the output video.
type Veo31RequestResolution string

const (
	Veo31RequestResolutionSeven20P Veo31RequestResolution = "720p"
	Veo31RequestResolutionOne080P  Veo31RequestResolution = "1080p"
	Veo31RequestResolutionFourK    Veo31RequestResolution = "4K"
)

func NewVeo31RequestResolutionFromString(s string) (Veo31RequestResolution, error) {
	switch s {
	case "720p":
		return Veo31RequestResolutionSeven20P, nil
	case "1080p":
		return Veo31RequestResolutionOne080P, nil
	case "4K":
		return Veo31RequestResolutionFourK, nil
	}
	var t Veo31RequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Veo31RequestResolution) Ptr() *Veo31RequestResolution {
	return &e
}
