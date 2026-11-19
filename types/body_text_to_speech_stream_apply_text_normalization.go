package types

import (
	"fmt"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
type BodyTextToSpeechStreamApplyTextNormalization string

const (
	BodyTextToSpeechStreamApplyTextNormalizationAuto BodyTextToSpeechStreamApplyTextNormalization = "auto"
	BodyTextToSpeechStreamApplyTextNormalizationOn   BodyTextToSpeechStreamApplyTextNormalization = "on"
	BodyTextToSpeechStreamApplyTextNormalizationOff  BodyTextToSpeechStreamApplyTextNormalization = "off"
)

func NewBodyTextToSpeechStreamApplyTextNormalizationFromString(s string) (BodyTextToSpeechStreamApplyTextNormalization, error) {
	switch s {
	case "auto":
		return BodyTextToSpeechStreamApplyTextNormalizationAuto, nil
	case "on":
		return BodyTextToSpeechStreamApplyTextNormalizationOn, nil
	case "off":
		return BodyTextToSpeechStreamApplyTextNormalizationOff, nil
	}
	var t BodyTextToSpeechStreamApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyTextToSpeechStreamApplyTextNormalization) Ptr() *BodyTextToSpeechStreamApplyTextNormalization {
	return &e
}
