package types

import (
	"fmt"
)

// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
type Gemini3ProImageRequestAspectRatio string

const (
	Gemini3ProImageRequestAspectRatioAuto   Gemini3ProImageRequestAspectRatio = "auto"
	Gemini3ProImageRequestAspectRatioOne1   Gemini3ProImageRequestAspectRatio = "1:1"
	Gemini3ProImageRequestAspectRatioTwo3   Gemini3ProImageRequestAspectRatio = "2:3"
	Gemini3ProImageRequestAspectRatioThree2 Gemini3ProImageRequestAspectRatio = "3:2"
	Gemini3ProImageRequestAspectRatioThree4 Gemini3ProImageRequestAspectRatio = "3:4"
	Gemini3ProImageRequestAspectRatioFour3  Gemini3ProImageRequestAspectRatio = "4:3"
	Gemini3ProImageRequestAspectRatioFour5  Gemini3ProImageRequestAspectRatio = "4:5"
	Gemini3ProImageRequestAspectRatioFive4  Gemini3ProImageRequestAspectRatio = "5:4"
	Gemini3ProImageRequestAspectRatioNine16 Gemini3ProImageRequestAspectRatio = "9:16"
	Gemini3ProImageRequestAspectRatioOne69  Gemini3ProImageRequestAspectRatio = "16:9"
	Gemini3ProImageRequestAspectRatioTwo19  Gemini3ProImageRequestAspectRatio = "21:9"
)

func NewGemini3ProImageRequestAspectRatioFromString(s string) (Gemini3ProImageRequestAspectRatio, error) {
	switch s {
	case "auto":
		return Gemini3ProImageRequestAspectRatioAuto, nil
	case "1:1":
		return Gemini3ProImageRequestAspectRatioOne1, nil
	case "2:3":
		return Gemini3ProImageRequestAspectRatioTwo3, nil
	case "3:2":
		return Gemini3ProImageRequestAspectRatioThree2, nil
	case "3:4":
		return Gemini3ProImageRequestAspectRatioThree4, nil
	case "4:3":
		return Gemini3ProImageRequestAspectRatioFour3, nil
	case "4:5":
		return Gemini3ProImageRequestAspectRatioFour5, nil
	case "5:4":
		return Gemini3ProImageRequestAspectRatioFive4, nil
	case "9:16":
		return Gemini3ProImageRequestAspectRatioNine16, nil
	case "16:9":
		return Gemini3ProImageRequestAspectRatioOne69, nil
	case "21:9":
		return Gemini3ProImageRequestAspectRatioTwo19, nil
	}
	var t Gemini3ProImageRequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Gemini3ProImageRequestAspectRatio) Ptr() *Gemini3ProImageRequestAspectRatio {
	return &e
}
