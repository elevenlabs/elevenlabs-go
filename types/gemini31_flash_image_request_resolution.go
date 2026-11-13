package types

import (
	"fmt"
)

// The resolution of the output image.
type Gemini31FlashImageRequestResolution string

const (
	Gemini31FlashImageRequestResolutionFive12 Gemini31FlashImageRequestResolution = "512"
	Gemini31FlashImageRequestResolutionOneK   Gemini31FlashImageRequestResolution = "1K"
	Gemini31FlashImageRequestResolutionTwoK   Gemini31FlashImageRequestResolution = "2K"
	Gemini31FlashImageRequestResolutionFourK  Gemini31FlashImageRequestResolution = "4K"
)

func NewGemini31FlashImageRequestResolutionFromString(s string) (Gemini31FlashImageRequestResolution, error) {
	switch s {
	case "512":
		return Gemini31FlashImageRequestResolutionFive12, nil
	case "1K":
		return Gemini31FlashImageRequestResolutionOneK, nil
	case "2K":
		return Gemini31FlashImageRequestResolutionTwoK, nil
	case "4K":
		return Gemini31FlashImageRequestResolutionFourK, nil
	}
	var t Gemini31FlashImageRequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Gemini31FlashImageRequestResolution) Ptr() *Gemini31FlashImageRequestResolution {
	return &e
}
