package types

import (
	"fmt"
)

// The aspect ratio of the output image.
type GptImage1RequestAspectRatio string

const (
	GptImage1RequestAspectRatioOne1   GptImage1RequestAspectRatio = "1:1"
	GptImage1RequestAspectRatioThree2 GptImage1RequestAspectRatio = "3:2"
	GptImage1RequestAspectRatioTwo3   GptImage1RequestAspectRatio = "2:3"
)

func NewGptImage1RequestAspectRatioFromString(s string) (GptImage1RequestAspectRatio, error) {
	switch s {
	case "1:1":
		return GptImage1RequestAspectRatioOne1, nil
	case "3:2":
		return GptImage1RequestAspectRatioThree2, nil
	case "2:3":
		return GptImage1RequestAspectRatioTwo3, nil
	}
	var t GptImage1RequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage1RequestAspectRatio) Ptr() *GptImage1RequestAspectRatio {
	return &e
}
