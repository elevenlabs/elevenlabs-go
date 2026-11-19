package types

import (
	"fmt"
)

// The resolution of the output video.
type BytedanceSeedance25RequestResolution string

const (
	BytedanceSeedance25RequestResolutionFour80P  BytedanceSeedance25RequestResolution = "480p"
	BytedanceSeedance25RequestResolutionSeven20P BytedanceSeedance25RequestResolution = "720p"
	BytedanceSeedance25RequestResolutionOne080P  BytedanceSeedance25RequestResolution = "1080p"
	BytedanceSeedance25RequestResolutionFourk    BytedanceSeedance25RequestResolution = "4k"
)

func NewBytedanceSeedance25RequestResolutionFromString(s string) (BytedanceSeedance25RequestResolution, error) {
	switch s {
	case "480p":
		return BytedanceSeedance25RequestResolutionFour80P, nil
	case "720p":
		return BytedanceSeedance25RequestResolutionSeven20P, nil
	case "1080p":
		return BytedanceSeedance25RequestResolutionOne080P, nil
	case "4k":
		return BytedanceSeedance25RequestResolutionFourk, nil
	}
	var t BytedanceSeedance25RequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BytedanceSeedance25RequestResolution) Ptr() *BytedanceSeedance25RequestResolution {
	return &e
}
