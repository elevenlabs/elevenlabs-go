package types

import (
	"fmt"
)

type TtsModelFamily string

const (
	TtsModelFamilyTurbo            TtsModelFamily = "turbo"
	TtsModelFamilyFlash            TtsModelFamily = "flash"
	TtsModelFamilyMultilingual     TtsModelFamily = "multilingual"
	TtsModelFamilyV3Conversational TtsModelFamily = "v3_conversational"
	TtsModelFamilyV4               TtsModelFamily = "v4"
	TtsModelFamilyV4Turbo          TtsModelFamily = "v4_turbo"
)

func NewTtsModelFamilyFromString(s string) (TtsModelFamily, error) {
	switch s {
	case "turbo":
		return TtsModelFamilyTurbo, nil
	case "flash":
		return TtsModelFamilyFlash, nil
	case "multilingual":
		return TtsModelFamilyMultilingual, nil
	case "v3_conversational":
		return TtsModelFamilyV3Conversational, nil
	case "v4":
		return TtsModelFamilyV4, nil
	case "v4_turbo":
		return TtsModelFamilyV4Turbo, nil
	}
	var t TtsModelFamily
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TtsModelFamily) Ptr() *TtsModelFamily {
	return &e
}
