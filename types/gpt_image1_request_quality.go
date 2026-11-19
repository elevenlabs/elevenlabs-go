package types

import (
	"fmt"
)

// The quality of the output image.
type GptImage1RequestQuality string

const (
	GptImage1RequestQualityLow    GptImage1RequestQuality = "low"
	GptImage1RequestQualityMedium GptImage1RequestQuality = "medium"
	GptImage1RequestQualityHigh   GptImage1RequestQuality = "high"
)

func NewGptImage1RequestQualityFromString(s string) (GptImage1RequestQuality, error) {
	switch s {
	case "low":
		return GptImage1RequestQualityLow, nil
	case "medium":
		return GptImage1RequestQualityMedium, nil
	case "high":
		return GptImage1RequestQualityHigh, nil
	}
	var t GptImage1RequestQuality
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage1RequestQuality) Ptr() *GptImage1RequestQuality {
	return &e
}
