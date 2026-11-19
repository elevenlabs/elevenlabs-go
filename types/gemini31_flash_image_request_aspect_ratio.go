package types

import (
	"fmt"
)

// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
type Gemini31FlashImageRequestAspectRatio string

const (
	Gemini31FlashImageRequestAspectRatioAuto   Gemini31FlashImageRequestAspectRatio = "auto"
	Gemini31FlashImageRequestAspectRatioOne1   Gemini31FlashImageRequestAspectRatio = "1:1"
	Gemini31FlashImageRequestAspectRatioTwo3   Gemini31FlashImageRequestAspectRatio = "2:3"
	Gemini31FlashImageRequestAspectRatioThree2 Gemini31FlashImageRequestAspectRatio = "3:2"
	Gemini31FlashImageRequestAspectRatioThree4 Gemini31FlashImageRequestAspectRatio = "3:4"
	Gemini31FlashImageRequestAspectRatioFour3  Gemini31FlashImageRequestAspectRatio = "4:3"
	Gemini31FlashImageRequestAspectRatioFour5  Gemini31FlashImageRequestAspectRatio = "4:5"
	Gemini31FlashImageRequestAspectRatioFive4  Gemini31FlashImageRequestAspectRatio = "5:4"
	Gemini31FlashImageRequestAspectRatioNine16 Gemini31FlashImageRequestAspectRatio = "9:16"
	Gemini31FlashImageRequestAspectRatioOne69  Gemini31FlashImageRequestAspectRatio = "16:9"
	Gemini31FlashImageRequestAspectRatioTwo19  Gemini31FlashImageRequestAspectRatio = "21:9"
	Gemini31FlashImageRequestAspectRatioOne4   Gemini31FlashImageRequestAspectRatio = "1:4"
	Gemini31FlashImageRequestAspectRatioFour1  Gemini31FlashImageRequestAspectRatio = "4:1"
	Gemini31FlashImageRequestAspectRatioOne8   Gemini31FlashImageRequestAspectRatio = "1:8"
	Gemini31FlashImageRequestAspectRatioEight1 Gemini31FlashImageRequestAspectRatio = "8:1"
)

func NewGemini31FlashImageRequestAspectRatioFromString(s string) (Gemini31FlashImageRequestAspectRatio, error) {
	switch s {
	case "auto":
		return Gemini31FlashImageRequestAspectRatioAuto, nil
	case "1:1":
		return Gemini31FlashImageRequestAspectRatioOne1, nil
	case "2:3":
		return Gemini31FlashImageRequestAspectRatioTwo3, nil
	case "3:2":
		return Gemini31FlashImageRequestAspectRatioThree2, nil
	case "3:4":
		return Gemini31FlashImageRequestAspectRatioThree4, nil
	case "4:3":
		return Gemini31FlashImageRequestAspectRatioFour3, nil
	case "4:5":
		return Gemini31FlashImageRequestAspectRatioFour5, nil
	case "5:4":
		return Gemini31FlashImageRequestAspectRatioFive4, nil
	case "9:16":
		return Gemini31FlashImageRequestAspectRatioNine16, nil
	case "16:9":
		return Gemini31FlashImageRequestAspectRatioOne69, nil
	case "21:9":
		return Gemini31FlashImageRequestAspectRatioTwo19, nil
	case "1:4":
		return Gemini31FlashImageRequestAspectRatioOne4, nil
	case "4:1":
		return Gemini31FlashImageRequestAspectRatioFour1, nil
	case "1:8":
		return Gemini31FlashImageRequestAspectRatioOne8, nil
	case "8:1":
		return Gemini31FlashImageRequestAspectRatioEight1, nil
	}
	var t Gemini31FlashImageRequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Gemini31FlashImageRequestAspectRatio) Ptr() *Gemini31FlashImageRequestAspectRatio {
	return &e
}
