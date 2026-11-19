package types

import (
	"fmt"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
type BodyTextToSpeechFullApplyTextNormalization string

const (
	BodyTextToSpeechFullApplyTextNormalizationAuto BodyTextToSpeechFullApplyTextNormalization = "auto"
	BodyTextToSpeechFullApplyTextNormalizationOn   BodyTextToSpeechFullApplyTextNormalization = "on"
	BodyTextToSpeechFullApplyTextNormalizationOff  BodyTextToSpeechFullApplyTextNormalization = "off"
)

func NewBodyTextToSpeechFullApplyTextNormalizationFromString(s string) (BodyTextToSpeechFullApplyTextNormalization, error) {
	switch s {
	case "auto":
		return BodyTextToSpeechFullApplyTextNormalizationAuto, nil
	case "on":
		return BodyTextToSpeechFullApplyTextNormalizationOn, nil
	case "off":
		return BodyTextToSpeechFullApplyTextNormalizationOff, nil
	}
	var t BodyTextToSpeechFullApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyTextToSpeechFullApplyTextNormalization) Ptr() *BodyTextToSpeechFullApplyTextNormalization {
	return &e
}
