package types

import (
	"fmt"
)

// The resolution of the output video.
type CreatifyAuroraRequestResolution string

const (
	CreatifyAuroraRequestResolutionFour80P  CreatifyAuroraRequestResolution = "480p"
	CreatifyAuroraRequestResolutionSeven20P CreatifyAuroraRequestResolution = "720p"
)

func NewCreatifyAuroraRequestResolutionFromString(s string) (CreatifyAuroraRequestResolution, error) {
	switch s {
	case "480p":
		return CreatifyAuroraRequestResolutionFour80P, nil
	case "720p":
		return CreatifyAuroraRequestResolutionSeven20P, nil
	}
	var t CreatifyAuroraRequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreatifyAuroraRequestResolution) Ptr() *CreatifyAuroraRequestResolution {
	return &e
}
