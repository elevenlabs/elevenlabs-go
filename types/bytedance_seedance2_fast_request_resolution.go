package types

import (
	"fmt"
)

// The resolution of the output video.
type BytedanceSeedance2FastRequestResolution string

const (
	BytedanceSeedance2FastRequestResolutionFour80P  BytedanceSeedance2FastRequestResolution = "480p"
	BytedanceSeedance2FastRequestResolutionSeven20P BytedanceSeedance2FastRequestResolution = "720p"
)

func NewBytedanceSeedance2FastRequestResolutionFromString(s string) (BytedanceSeedance2FastRequestResolution, error) {
	switch s {
	case "480p":
		return BytedanceSeedance2FastRequestResolutionFour80P, nil
	case "720p":
		return BytedanceSeedance2FastRequestResolutionSeven20P, nil
	}
	var t BytedanceSeedance2FastRequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BytedanceSeedance2FastRequestResolution) Ptr() *BytedanceSeedance2FastRequestResolution {
	return &e
}
