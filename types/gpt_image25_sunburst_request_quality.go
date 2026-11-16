package types

import (
	"fmt"
)

// The quality of the output image.
type GptImage25SunburstRequestQuality string

const (
	GptImage25SunburstRequestQualityLow    GptImage25SunburstRequestQuality = "low"
	GptImage25SunburstRequestQualityMedium GptImage25SunburstRequestQuality = "medium"
	GptImage25SunburstRequestQualityHigh   GptImage25SunburstRequestQuality = "high"
	GptImage25SunburstRequestQualityXhigh  GptImage25SunburstRequestQuality = "xhigh"
	GptImage25SunburstRequestQualityMax    GptImage25SunburstRequestQuality = "max"
)

func NewGptImage25SunburstRequestQualityFromString(s string) (GptImage25SunburstRequestQuality, error) {
	switch s {
	case "low":
		return GptImage25SunburstRequestQualityLow, nil
	case "medium":
		return GptImage25SunburstRequestQualityMedium, nil
	case "high":
		return GptImage25SunburstRequestQualityHigh, nil
	case "xhigh":
		return GptImage25SunburstRequestQualityXhigh, nil
	case "max":
		return GptImage25SunburstRequestQualityMax, nil
	}
	var t GptImage25SunburstRequestQuality
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage25SunburstRequestQuality) Ptr() *GptImage25SunburstRequestQuality {
	return &e
}
