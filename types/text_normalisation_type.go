package types

import (
	"fmt"
)

// Method for converting numbers to words before sending to TTS
type TextNormalisationType string

const (
	TextNormalisationTypeSystemPrompt TextNormalisationType = "system_prompt"
	TextNormalisationTypeElevenlabs   TextNormalisationType = "elevenlabs"
)

func NewTextNormalisationTypeFromString(s string) (TextNormalisationType, error) {
	switch s {
	case "system_prompt":
		return TextNormalisationTypeSystemPrompt, nil
	case "elevenlabs":
		return TextNormalisationTypeElevenlabs, nil
	}
	var t TextNormalisationType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TextNormalisationType) Ptr() *TextNormalisationType {
	return &e
}
