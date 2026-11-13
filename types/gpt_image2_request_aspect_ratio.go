package types

import (
	"fmt"
)

// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
type GptImage2RequestAspectRatio string

const (
	GptImage2RequestAspectRatioAuto   GptImage2RequestAspectRatio = "auto"
	GptImage2RequestAspectRatioOne1   GptImage2RequestAspectRatio = "1:1"
	GptImage2RequestAspectRatioFour5  GptImage2RequestAspectRatio = "4:5"
	GptImage2RequestAspectRatioFive4  GptImage2RequestAspectRatio = "5:4"
	GptImage2RequestAspectRatioThree4 GptImage2RequestAspectRatio = "3:4"
	GptImage2RequestAspectRatioFour3  GptImage2RequestAspectRatio = "4:3"
	GptImage2RequestAspectRatioTwo3   GptImage2RequestAspectRatio = "2:3"
	GptImage2RequestAspectRatioThree2 GptImage2RequestAspectRatio = "3:2"
	GptImage2RequestAspectRatioOne2   GptImage2RequestAspectRatio = "1:2"
	GptImage2RequestAspectRatioTwo1   GptImage2RequestAspectRatio = "2:1"
	GptImage2RequestAspectRatioNine16 GptImage2RequestAspectRatio = "9:16"
	GptImage2RequestAspectRatioOne69  GptImage2RequestAspectRatio = "16:9"
	GptImage2RequestAspectRatioTwo19  GptImage2RequestAspectRatio = "21:9"
	GptImage2RequestAspectRatioOne3   GptImage2RequestAspectRatio = "1:3"
	GptImage2RequestAspectRatioThree1 GptImage2RequestAspectRatio = "3:1"
)

func NewGptImage2RequestAspectRatioFromString(s string) (GptImage2RequestAspectRatio, error) {
	switch s {
	case "auto":
		return GptImage2RequestAspectRatioAuto, nil
	case "1:1":
		return GptImage2RequestAspectRatioOne1, nil
	case "4:5":
		return GptImage2RequestAspectRatioFour5, nil
	case "5:4":
		return GptImage2RequestAspectRatioFive4, nil
	case "3:4":
		return GptImage2RequestAspectRatioThree4, nil
	case "4:3":
		return GptImage2RequestAspectRatioFour3, nil
	case "2:3":
		return GptImage2RequestAspectRatioTwo3, nil
	case "3:2":
		return GptImage2RequestAspectRatioThree2, nil
	case "1:2":
		return GptImage2RequestAspectRatioOne2, nil
	case "2:1":
		return GptImage2RequestAspectRatioTwo1, nil
	case "9:16":
		return GptImage2RequestAspectRatioNine16, nil
	case "16:9":
		return GptImage2RequestAspectRatioOne69, nil
	case "21:9":
		return GptImage2RequestAspectRatioTwo19, nil
	case "1:3":
		return GptImage2RequestAspectRatioOne3, nil
	case "3:1":
		return GptImage2RequestAspectRatioThree1, nil
	}
	var t GptImage2RequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage2RequestAspectRatio) Ptr() *GptImage2RequestAspectRatio {
	return &e
}
