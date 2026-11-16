package types

import (
	"fmt"
)

// The aspect ratio of the output image.
type GptImage15RequestAspectRatio string

const (
	GptImage15RequestAspectRatioOne1   GptImage15RequestAspectRatio = "1:1"
	GptImage15RequestAspectRatioThree2 GptImage15RequestAspectRatio = "3:2"
	GptImage15RequestAspectRatioTwo3   GptImage15RequestAspectRatio = "2:3"
)

func NewGptImage15RequestAspectRatioFromString(s string) (GptImage15RequestAspectRatio, error) {
	switch s {
	case "1:1":
		return GptImage15RequestAspectRatioOne1, nil
	case "3:2":
		return GptImage15RequestAspectRatioThree2, nil
	case "2:3":
		return GptImage15RequestAspectRatioTwo3, nil
	}
	var t GptImage15RequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage15RequestAspectRatio) Ptr() *GptImage15RequestAspectRatio {
	return &e
}
