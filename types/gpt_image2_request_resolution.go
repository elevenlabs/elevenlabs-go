package types

import (
	"fmt"
)

// The resolution of the output image.
type GptImage2RequestResolution string

const (
	GptImage2RequestResolutionOneK  GptImage2RequestResolution = "1K"
	GptImage2RequestResolutionTwoK  GptImage2RequestResolution = "2K"
	GptImage2RequestResolutionFourK GptImage2RequestResolution = "4K"
)

func NewGptImage2RequestResolutionFromString(s string) (GptImage2RequestResolution, error) {
	switch s {
	case "1K":
		return GptImage2RequestResolutionOneK, nil
	case "2K":
		return GptImage2RequestResolutionTwoK, nil
	case "4K":
		return GptImage2RequestResolutionFourK, nil
	}
	var t GptImage2RequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage2RequestResolution) Ptr() *GptImage2RequestResolution {
	return &e
}
