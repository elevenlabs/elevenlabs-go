package types

import (
	"fmt"
)

// The aspect ratio of the output image. With `auto`, the model picks an aspect ratio based on the inputs.
type Gemini25FlashImageRequestAspectRatio string

const (
	Gemini25FlashImageRequestAspectRatioAuto   Gemini25FlashImageRequestAspectRatio = "auto"
	Gemini25FlashImageRequestAspectRatioOne1   Gemini25FlashImageRequestAspectRatio = "1:1"
	Gemini25FlashImageRequestAspectRatioTwo3   Gemini25FlashImageRequestAspectRatio = "2:3"
	Gemini25FlashImageRequestAspectRatioThree2 Gemini25FlashImageRequestAspectRatio = "3:2"
	Gemini25FlashImageRequestAspectRatioThree4 Gemini25FlashImageRequestAspectRatio = "3:4"
	Gemini25FlashImageRequestAspectRatioFour3  Gemini25FlashImageRequestAspectRatio = "4:3"
	Gemini25FlashImageRequestAspectRatioFour5  Gemini25FlashImageRequestAspectRatio = "4:5"
	Gemini25FlashImageRequestAspectRatioFive4  Gemini25FlashImageRequestAspectRatio = "5:4"
	Gemini25FlashImageRequestAspectRatioNine16 Gemini25FlashImageRequestAspectRatio = "9:16"
	Gemini25FlashImageRequestAspectRatioOne69  Gemini25FlashImageRequestAspectRatio = "16:9"
	Gemini25FlashImageRequestAspectRatioTwo19  Gemini25FlashImageRequestAspectRatio = "21:9"
)

func NewGemini25FlashImageRequestAspectRatioFromString(s string) (Gemini25FlashImageRequestAspectRatio, error) {
	switch s {
	case "auto":
		return Gemini25FlashImageRequestAspectRatioAuto, nil
	case "1:1":
		return Gemini25FlashImageRequestAspectRatioOne1, nil
	case "2:3":
		return Gemini25FlashImageRequestAspectRatioTwo3, nil
	case "3:2":
		return Gemini25FlashImageRequestAspectRatioThree2, nil
	case "3:4":
		return Gemini25FlashImageRequestAspectRatioThree4, nil
	case "4:3":
		return Gemini25FlashImageRequestAspectRatioFour3, nil
	case "4:5":
		return Gemini25FlashImageRequestAspectRatioFour5, nil
	case "5:4":
		return Gemini25FlashImageRequestAspectRatioFive4, nil
	case "9:16":
		return Gemini25FlashImageRequestAspectRatioNine16, nil
	case "16:9":
		return Gemini25FlashImageRequestAspectRatioOne69, nil
	case "21:9":
		return Gemini25FlashImageRequestAspectRatioTwo19, nil
	}
	var t Gemini25FlashImageRequestAspectRatio
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Gemini25FlashImageRequestAspectRatio) Ptr() *Gemini25FlashImageRequestAspectRatio {
	return &e
}
