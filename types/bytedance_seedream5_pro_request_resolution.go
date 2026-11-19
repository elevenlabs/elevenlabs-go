package types

import (
	"fmt"
)

// The resolution of the output image.
type BytedanceSeedream5ProRequestResolution string

const (
	BytedanceSeedream5ProRequestResolutionOneK BytedanceSeedream5ProRequestResolution = "1K"
	BytedanceSeedream5ProRequestResolutionTwoK BytedanceSeedream5ProRequestResolution = "2K"
)

func NewBytedanceSeedream5ProRequestResolutionFromString(s string) (BytedanceSeedream5ProRequestResolution, error) {
	switch s {
	case "1K":
		return BytedanceSeedream5ProRequestResolutionOneK, nil
	case "2K":
		return BytedanceSeedream5ProRequestResolutionTwoK, nil
	}
	var t BytedanceSeedream5ProRequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BytedanceSeedream5ProRequestResolution) Ptr() *BytedanceSeedream5ProRequestResolution {
	return &e
}
