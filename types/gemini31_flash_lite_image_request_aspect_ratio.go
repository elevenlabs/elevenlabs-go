package types

import (
	"fmt"
)

// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
type Gemini31FlashLiteImageRequestAspectRatio string

const (
	Gemini31FlashLiteImageRequestAspectRatioAuto   Gemini31FlashLiteImageRequestAspectRatio = "auto"
	Gemini31FlashLiteImageRequestAspectRatioOne1   Gemini31FlashLiteImageRequestAspectRatio = "1:1"
	Gemini31FlashLiteImageRequestAspectRatioTwo3   Gemini31FlashLiteImageRequestAspectRatio = "2:3"
	Gemini31FlashLiteImageRequestAspectRatioThree2 Gemini31FlashLiteImageRequestAspectRatio = "3:2"
	Gemini31FlashLiteImageRequestAspectRatioThree4 Gemini31FlashLiteImageRequestAspectRatio = "3:4"
	Gemini31FlashLiteImageRequestAspectRatioFour3  Gemini31FlashLiteImageRequestAspectRatio = "4:3"
	Gemini31FlashLiteImageRequestAspectRatioFour5  Gemini31FlashLiteImageRequestAspectRatio = "4:5"
	Gemini31FlashLiteImageRequestAspectRatioFive4  Gemini31FlashLiteImageRequestAspectRatio = "5:4"
	Gemini31FlashLiteImageRequestAspectRatioNine16 Gemini31FlashLiteImageRequestAspectRatio = "9:16"
	Gemini31FlashLiteImageRequestAspectRatioOne69  Gemini31FlashLiteImageRequestAspectRatio = "16:9"
	Gemini31FlashLiteImageRequestAspectRatioTwo19  Gemini31FlashLiteImageRequestAspectRatio = "21:9"
)

func NewGemini31FlashLiteImageRequestAspectRatioFromString(s string) (Gemini31FlashLiteImageRequestAspectRatio, error) {
	switch s {
	case "auto":
		return Gemini31FlashLiteImageRequestAspectRatioAuto, nil
	case "1:1":
		return Gemini31FlashLiteImageRequestAspectRatioOne1, nil
	case "2:3":
		return Gemini31FlashLiteImageRequestAspectRatioTwo3, nil
	case "3:2":
		return Gemini31FlashLiteImageRequestAspectRatioThree2, nil
	case "3:4":
		return Gemini31FlashLiteImageRequestAspectRatioThree4, nil
	case "4:3":
		return Gemini31FlashLiteImageRequestAspectRatioFour3, nil
	case "4:5":
		return Gemini31FlashLiteImageRequestAspectRatioFour5, nil
	case "5:4":
		return Gemini31FlashLiteImageRequestAspectRatioFive4, nil
	case "9:16":
		return Gemini31FlashLiteImageRequestAspectRatioNine16, nil
	case "16:9":
		return Gemini31FlashLiteImageRequestAspectRatioOne69, nil
	case "21:9":
		return Gemini31FlashLiteImageRequestAspectRatioTwo19, nil
	}
	var t Gemini31FlashLiteImageRequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Gemini31FlashLiteImageRequestAspectRatio) Ptr() *Gemini31FlashLiteImageRequestAspectRatio {
	return &e
}
