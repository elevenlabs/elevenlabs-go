package types

import (
	"fmt"
)

// The resolution of the output video.
type BytedanceSeedance2MiniRequestResolution string

const (
	BytedanceSeedance2MiniRequestResolutionFour80P  BytedanceSeedance2MiniRequestResolution = "480p"
	BytedanceSeedance2MiniRequestResolutionSeven20P BytedanceSeedance2MiniRequestResolution = "720p"
)

func NewBytedanceSeedance2MiniRequestResolutionFromString(s string) (BytedanceSeedance2MiniRequestResolution, error) {
	switch s {
	case "480p":
		return BytedanceSeedance2MiniRequestResolutionFour80P, nil
	case "720p":
		return BytedanceSeedance2MiniRequestResolutionSeven20P, nil
	}
	var t BytedanceSeedance2MiniRequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BytedanceSeedance2MiniRequestResolution) Ptr() *BytedanceSeedance2MiniRequestResolution {
	return &e
}
