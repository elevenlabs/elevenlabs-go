package types

import (
	"fmt"
)

type TtsConversationalModel string

const (
	TtsConversationalModelElevenTurboV2          TtsConversationalModel = "eleven_turbo_v2"
	TtsConversationalModelElevenTurboV25         TtsConversationalModel = "eleven_turbo_v2_5"
	TtsConversationalModelElevenFlashV2          TtsConversationalModel = "eleven_flash_v2"
	TtsConversationalModelElevenFlashV25         TtsConversationalModel = "eleven_flash_v2_5"
	TtsConversationalModelElevenMultilingualV2   TtsConversationalModel = "eleven_multilingual_v2"
	TtsConversationalModelElevenV3Conversational TtsConversationalModel = "eleven_v3_conversational"
	TtsConversationalModelElevenV4               TtsConversationalModel = "eleven_v4"
	TtsConversationalModelElevenV4Turbo          TtsConversationalModel = "eleven_v4_turbo"
)

func NewTtsConversationalModelFromString(s string) (TtsConversationalModel, error) {
	switch s {
	case "eleven_turbo_v2":
		return TtsConversationalModelElevenTurboV2, nil
	case "eleven_turbo_v2_5":
		return TtsConversationalModelElevenTurboV25, nil
	case "eleven_flash_v2":
		return TtsConversationalModelElevenFlashV2, nil
	case "eleven_flash_v2_5":
		return TtsConversationalModelElevenFlashV25, nil
	case "eleven_multilingual_v2":
		return TtsConversationalModelElevenMultilingualV2, nil
	case "eleven_v3_conversational":
		return TtsConversationalModelElevenV3Conversational, nil
	case "eleven_v4":
		return TtsConversationalModelElevenV4, nil
	case "eleven_v4_turbo":
		return TtsConversationalModelElevenV4Turbo, nil
	}
	var t TtsConversationalModel
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TtsConversationalModel) Ptr() *TtsConversationalModel {
	return &e
}
