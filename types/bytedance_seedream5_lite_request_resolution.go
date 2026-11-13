package types

import (
	"fmt"
)

// The resolution of the output image.
type BytedanceSeedream5LiteRequestResolution string

const (
	BytedanceSeedream5LiteRequestResolutionTwoK   BytedanceSeedream5LiteRequestResolution = "2K"
	BytedanceSeedream5LiteRequestResolutionThreeK BytedanceSeedream5LiteRequestResolution = "3K"
)

func NewBytedanceSeedream5LiteRequestResolutionFromString(s string) (BytedanceSeedream5LiteRequestResolution, error) {
	switch s {
	case "2K":
		return BytedanceSeedream5LiteRequestResolutionTwoK, nil
	case "3K":
		return BytedanceSeedream5LiteRequestResolutionThreeK, nil
	}
	var t BytedanceSeedream5LiteRequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BytedanceSeedream5LiteRequestResolution) Ptr() *BytedanceSeedream5LiteRequestResolution {
	return &e
}
