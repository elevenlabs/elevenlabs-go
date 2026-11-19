package types

import (
	"fmt"
)

// The resolution of the output video.
type BytedanceSeedance2RequestResolution string

const (
	BytedanceSeedance2RequestResolutionFour80P  BytedanceSeedance2RequestResolution = "480p"
	BytedanceSeedance2RequestResolutionSeven20P BytedanceSeedance2RequestResolution = "720p"
	BytedanceSeedance2RequestResolutionOne080P  BytedanceSeedance2RequestResolution = "1080p"
	BytedanceSeedance2RequestResolutionFourk    BytedanceSeedance2RequestResolution = "4k"
)

func NewBytedanceSeedance2RequestResolutionFromString(s string) (BytedanceSeedance2RequestResolution, error) {
	switch s {
	case "480p":
		return BytedanceSeedance2RequestResolutionFour80P, nil
	case "720p":
		return BytedanceSeedance2RequestResolutionSeven20P, nil
	case "1080p":
		return BytedanceSeedance2RequestResolutionOne080P, nil
	case "4k":
		return BytedanceSeedance2RequestResolutionFourk, nil
	}
	var t BytedanceSeedance2RequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BytedanceSeedance2RequestResolution) Ptr() *BytedanceSeedance2RequestResolution {
	return &e
}
