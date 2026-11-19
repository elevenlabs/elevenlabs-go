package types

import (
	"fmt"
)

// The quality of the output image.
type GptImage2RequestQuality string

const (
	GptImage2RequestQualityLow    GptImage2RequestQuality = "low"
	GptImage2RequestQualityMedium GptImage2RequestQuality = "medium"
	GptImage2RequestQualityHigh   GptImage2RequestQuality = "high"
)

func NewGptImage2RequestQualityFromString(s string) (GptImage2RequestQuality, error) {
	switch s {
	case "low":
		return GptImage2RequestQualityLow, nil
	case "medium":
		return GptImage2RequestQualityMedium, nil
	case "high":
		return GptImage2RequestQualityHigh, nil
	}
	var t GptImage2RequestQuality
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage2RequestQuality) Ptr() *GptImage2RequestQuality {
	return &e
}
