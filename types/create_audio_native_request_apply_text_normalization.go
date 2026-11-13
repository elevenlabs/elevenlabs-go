package types

import (
	"fmt"
)

type CreateAudioNativeRequestApplyTextNormalization string

const (
	CreateAudioNativeRequestApplyTextNormalizationAuto         CreateAudioNativeRequestApplyTextNormalization = "auto"
	CreateAudioNativeRequestApplyTextNormalizationOn           CreateAudioNativeRequestApplyTextNormalization = "on"
	CreateAudioNativeRequestApplyTextNormalizationOff          CreateAudioNativeRequestApplyTextNormalization = "off"
	CreateAudioNativeRequestApplyTextNormalizationApplyEnglish CreateAudioNativeRequestApplyTextNormalization = "apply_english"
)

func NewCreateAudioNativeRequestApplyTextNormalizationFromString(s string) (CreateAudioNativeRequestApplyTextNormalization, error) {
	switch s {
	case "auto":
		return CreateAudioNativeRequestApplyTextNormalizationAuto, nil
	case "on":
		return CreateAudioNativeRequestApplyTextNormalizationOn, nil
	case "off":
		return CreateAudioNativeRequestApplyTextNormalizationOff, nil
	case "apply_english":
		return CreateAudioNativeRequestApplyTextNormalizationApplyEnglish, nil
	}
	var t CreateAudioNativeRequestApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CreateAudioNativeRequestApplyTextNormalization) Ptr() *CreateAudioNativeRequestApplyTextNormalization {
	return &e
}
