package types

import (
	"fmt"
)

// TTS model used to synthesize the translated speech.
type TranslationRealtimeTtsModelId string

const (
	TranslationRealtimeTtsModelIdElevenFlashV25 TranslationRealtimeTtsModelId = "eleven_flash_v2_5"
	TranslationRealtimeTtsModelIdElevenV3       TranslationRealtimeTtsModelId = "eleven_v3"
	TranslationRealtimeTtsModelIdElevenV4Turbo  TranslationRealtimeTtsModelId = "eleven_v4_turbo"
)

func NewTranslationRealtimeTtsModelIdFromString(s string) (TranslationRealtimeTtsModelId, error) {
	switch s {
	case "eleven_flash_v2_5":
		return TranslationRealtimeTtsModelIdElevenFlashV25, nil
	case "eleven_v3":
		return TranslationRealtimeTtsModelIdElevenV3, nil
	case "eleven_v4_turbo":
		return TranslationRealtimeTtsModelIdElevenV4Turbo, nil
	}
	var t TranslationRealtimeTtsModelId
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TranslationRealtimeTtsModelId) Ptr() *TranslationRealtimeTtsModelId {
	return &e
}
