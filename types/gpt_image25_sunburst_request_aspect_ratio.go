package types

import (
	"fmt"
)

// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
type GptImage25SunburstRequestAspectRatio string

const (
	GptImage25SunburstRequestAspectRatioAuto   GptImage25SunburstRequestAspectRatio = "auto"
	GptImage25SunburstRequestAspectRatioOne1   GptImage25SunburstRequestAspectRatio = "1:1"
	GptImage25SunburstRequestAspectRatioFour5  GptImage25SunburstRequestAspectRatio = "4:5"
	GptImage25SunburstRequestAspectRatioFive4  GptImage25SunburstRequestAspectRatio = "5:4"
	GptImage25SunburstRequestAspectRatioThree4 GptImage25SunburstRequestAspectRatio = "3:4"
	GptImage25SunburstRequestAspectRatioFour3  GptImage25SunburstRequestAspectRatio = "4:3"
	GptImage25SunburstRequestAspectRatioTwo3   GptImage25SunburstRequestAspectRatio = "2:3"
	GptImage25SunburstRequestAspectRatioThree2 GptImage25SunburstRequestAspectRatio = "3:2"
	GptImage25SunburstRequestAspectRatioOne2   GptImage25SunburstRequestAspectRatio = "1:2"
	GptImage25SunburstRequestAspectRatioTwo1   GptImage25SunburstRequestAspectRatio = "2:1"
	GptImage25SunburstRequestAspectRatioNine16 GptImage25SunburstRequestAspectRatio = "9:16"
	GptImage25SunburstRequestAspectRatioOne69  GptImage25SunburstRequestAspectRatio = "16:9"
	GptImage25SunburstRequestAspectRatioTwo19  GptImage25SunburstRequestAspectRatio = "21:9"
	GptImage25SunburstRequestAspectRatioOne3   GptImage25SunburstRequestAspectRatio = "1:3"
	GptImage25SunburstRequestAspectRatioThree1 GptImage25SunburstRequestAspectRatio = "3:1"
)

func NewGptImage25SunburstRequestAspectRatioFromString(s string) (GptImage25SunburstRequestAspectRatio, error) {
	switch s {
	case "auto":
		return GptImage25SunburstRequestAspectRatioAuto, nil
	case "1:1":
		return GptImage25SunburstRequestAspectRatioOne1, nil
	case "4:5":
		return GptImage25SunburstRequestAspectRatioFour5, nil
	case "5:4":
		return GptImage25SunburstRequestAspectRatioFive4, nil
	case "3:4":
		return GptImage25SunburstRequestAspectRatioThree4, nil
	case "4:3":
		return GptImage25SunburstRequestAspectRatioFour3, nil
	case "2:3":
		return GptImage25SunburstRequestAspectRatioTwo3, nil
	case "3:2":
		return GptImage25SunburstRequestAspectRatioThree2, nil
	case "1:2":
		return GptImage25SunburstRequestAspectRatioOne2, nil
	case "2:1":
		return GptImage25SunburstRequestAspectRatioTwo1, nil
	case "9:16":
		return GptImage25SunburstRequestAspectRatioNine16, nil
	case "16:9":
		return GptImage25SunburstRequestAspectRatioOne69, nil
	case "21:9":
		return GptImage25SunburstRequestAspectRatioTwo19, nil
	case "1:3":
		return GptImage25SunburstRequestAspectRatioOne3, nil
	case "3:1":
		return GptImage25SunburstRequestAspectRatioThree1, nil
	}
	var t GptImage25SunburstRequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GptImage25SunburstRequestAspectRatio) Ptr() *GptImage25SunburstRequestAspectRatio {
	return &e
}
