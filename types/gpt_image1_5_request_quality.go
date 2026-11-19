package types

import (
	"fmt"
)

// The quality of the output image.
type GptImage15RequestQuality string

const (
	GptImage15RequestQualityLow    GptImage15RequestQuality = "low"
	GptImage15RequestQualityMedium GptImage15RequestQuality = "medium"
	GptImage15RequestQualityHigh   GptImage15RequestQuality = "high"
)

func NewGptImage15RequestQualityFromString(s string) (GptImage15RequestQuality, error) {
	switch s {
	case "low":
		return GptImage15RequestQualityLow, nil
	case "medium":
		return GptImage15RequestQualityMedium, nil
	case "high":
		return GptImage15RequestQualityHigh, nil
	}
	var t GptImage15RequestQuality
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage15RequestQuality) Ptr() *GptImage15RequestQuality {
	return &e
}
