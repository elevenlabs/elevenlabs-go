package types

import (
	"fmt"
)

// The quality of the output image.
type GptImage25FlareRequestQuality string

const (
	GptImage25FlareRequestQualityLow    GptImage25FlareRequestQuality = "low"
	GptImage25FlareRequestQualityMedium GptImage25FlareRequestQuality = "medium"
	GptImage25FlareRequestQualityHigh   GptImage25FlareRequestQuality = "high"
	GptImage25FlareRequestQualityXhigh  GptImage25FlareRequestQuality = "xhigh"
	GptImage25FlareRequestQualityMax    GptImage25FlareRequestQuality = "max"
)

func NewGptImage25FlareRequestQualityFromString(s string) (GptImage25FlareRequestQuality, error) {
	switch s {
	case "low":
		return GptImage25FlareRequestQualityLow, nil
	case "medium":
		return GptImage25FlareRequestQualityMedium, nil
	case "high":
		return GptImage25FlareRequestQualityHigh, nil
	case "xhigh":
		return GptImage25FlareRequestQualityXhigh, nil
	case "max":
		return GptImage25FlareRequestQualityMax, nil
	}
	var t GptImage25FlareRequestQuality
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage25FlareRequestQuality) Ptr() *GptImage25FlareRequestQuality {
	return &e
}
