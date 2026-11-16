package types

import (
	"fmt"
)

// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
type GptImage25FlareRequestAspectRatio string

const (
	GptImage25FlareRequestAspectRatioAuto   GptImage25FlareRequestAspectRatio = "auto"
	GptImage25FlareRequestAspectRatioOne1   GptImage25FlareRequestAspectRatio = "1:1"
	GptImage25FlareRequestAspectRatioFour5  GptImage25FlareRequestAspectRatio = "4:5"
	GptImage25FlareRequestAspectRatioFive4  GptImage25FlareRequestAspectRatio = "5:4"
	GptImage25FlareRequestAspectRatioThree4 GptImage25FlareRequestAspectRatio = "3:4"
	GptImage25FlareRequestAspectRatioFour3  GptImage25FlareRequestAspectRatio = "4:3"
	GptImage25FlareRequestAspectRatioTwo3   GptImage25FlareRequestAspectRatio = "2:3"
	GptImage25FlareRequestAspectRatioThree2 GptImage25FlareRequestAspectRatio = "3:2"
	GptImage25FlareRequestAspectRatioOne2   GptImage25FlareRequestAspectRatio = "1:2"
	GptImage25FlareRequestAspectRatioTwo1   GptImage25FlareRequestAspectRatio = "2:1"
	GptImage25FlareRequestAspectRatioNine16 GptImage25FlareRequestAspectRatio = "9:16"
	GptImage25FlareRequestAspectRatioOne69  GptImage25FlareRequestAspectRatio = "16:9"
	GptImage25FlareRequestAspectRatioTwo19  GptImage25FlareRequestAspectRatio = "21:9"
	GptImage25FlareRequestAspectRatioOne3   GptImage25FlareRequestAspectRatio = "1:3"
	GptImage25FlareRequestAspectRatioThree1 GptImage25FlareRequestAspectRatio = "3:1"
)

func NewGptImage25FlareRequestAspectRatioFromString(s string) (GptImage25FlareRequestAspectRatio, error) {
	switch s {
	case "auto":
		return GptImage25FlareRequestAspectRatioAuto, nil
	case "1:1":
		return GptImage25FlareRequestAspectRatioOne1, nil
	case "4:5":
		return GptImage25FlareRequestAspectRatioFour5, nil
	case "5:4":
		return GptImage25FlareRequestAspectRatioFive4, nil
	case "3:4":
		return GptImage25FlareRequestAspectRatioThree4, nil
	case "4:3":
		return GptImage25FlareRequestAspectRatioFour3, nil
	case "2:3":
		return GptImage25FlareRequestAspectRatioTwo3, nil
	case "3:2":
		return GptImage25FlareRequestAspectRatioThree2, nil
	case "1:2":
		return GptImage25FlareRequestAspectRatioOne2, nil
	case "2:1":
		return GptImage25FlareRequestAspectRatioTwo1, nil
	case "9:16":
		return GptImage25FlareRequestAspectRatioNine16, nil
	case "16:9":
		return GptImage25FlareRequestAspectRatioOne69, nil
	case "21:9":
		return GptImage25FlareRequestAspectRatioTwo19, nil
	case "1:3":
		return GptImage25FlareRequestAspectRatioOne3, nil
	case "3:1":
		return GptImage25FlareRequestAspectRatioThree1, nil
	}
	var t GptImage25FlareRequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage25FlareRequestAspectRatio) Ptr() *GptImage25FlareRequestAspectRatio {
	return &e
}
