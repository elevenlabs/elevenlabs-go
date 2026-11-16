package types

import (
	"fmt"
)

// The resolution of the output image.
type Gemini3ProImageRequestResolution string

const (
	Gemini3ProImageRequestResolutionOneK  Gemini3ProImageRequestResolution = "1K"
	Gemini3ProImageRequestResolutionTwoK  Gemini3ProImageRequestResolution = "2K"
	Gemini3ProImageRequestResolutionFourK Gemini3ProImageRequestResolution = "4K"
)

func NewGemini3ProImageRequestResolutionFromString(s string) (Gemini3ProImageRequestResolution, error) {
	switch s {
	case "1K":
		return Gemini3ProImageRequestResolutionOneK, nil
	case "2K":
		return Gemini3ProImageRequestResolutionTwoK, nil
	case "4K":
		return Gemini3ProImageRequestResolutionFourK, nil
	}
	var t Gemini3ProImageRequestResolution
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Gemini3ProImageRequestResolution) Ptr() *Gemini3ProImageRequestResolution {
	return &e
}
