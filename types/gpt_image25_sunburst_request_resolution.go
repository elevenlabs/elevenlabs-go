package types

import (
	"fmt"
)

// The resolution of the output image.
type GptImage25SunburstRequestResolution string

const (
	GptImage25SunburstRequestResolutionOneK  GptImage25SunburstRequestResolution = "1K"
	GptImage25SunburstRequestResolutionTwoK  GptImage25SunburstRequestResolution = "2K"
	GptImage25SunburstRequestResolutionFourK GptImage25SunburstRequestResolution = "4K"
)

func NewGptImage25SunburstRequestResolutionFromString(s string) (GptImage25SunburstRequestResolution, error) {
	switch s {
	case "1K":
		return GptImage25SunburstRequestResolutionOneK, nil
	case "2K":
		return GptImage25SunburstRequestResolutionTwoK, nil
	case "4K":
		return GptImage25SunburstRequestResolutionFourK, nil
	}
	var t GptImage25SunburstRequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage25SunburstRequestResolution) Ptr() *GptImage25SunburstRequestResolution {
	return &e
}
