package types

import (
	"fmt"
)

// The resolution of the output image.
type GptImage25FlareRequestResolution string

const (
	GptImage25FlareRequestResolutionOneK  GptImage25FlareRequestResolution = "1K"
	GptImage25FlareRequestResolutionTwoK  GptImage25FlareRequestResolution = "2K"
	GptImage25FlareRequestResolutionFourK GptImage25FlareRequestResolution = "4K"
)

func NewGptImage25FlareRequestResolutionFromString(s string) (GptImage25FlareRequestResolution, error) {
	switch s {
	case "1K":
		return GptImage25FlareRequestResolutionOneK, nil
	case "2K":
		return GptImage25FlareRequestResolutionTwoK, nil
	case "4K":
		return GptImage25FlareRequestResolutionFourK, nil
	}
	var t GptImage25FlareRequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage25FlareRequestResolution) Ptr() *GptImage25FlareRequestResolution {
	return &e
}
