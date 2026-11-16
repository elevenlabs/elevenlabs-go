package types

import (
	"fmt"
)

// The resolution of the output video.
type Veo31FastRequestResolution string

const (
	Veo31FastRequestResolutionSeven20P Veo31FastRequestResolution = "720p"
	Veo31FastRequestResolutionOne080P  Veo31FastRequestResolution = "1080p"
	Veo31FastRequestResolutionFourK    Veo31FastRequestResolution = "4K"
)

func NewVeo31FastRequestResolutionFromString(s string) (Veo31FastRequestResolution, error) {
	switch s {
	case "720p":
		return Veo31FastRequestResolutionSeven20P, nil
	case "1080p":
		return Veo31FastRequestResolutionOne080P, nil
	case "4K":
		return Veo31FastRequestResolutionFourK, nil
	}
	var t Veo31FastRequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Veo31FastRequestResolution) Ptr() *Veo31FastRequestResolution {
	return &e
}
