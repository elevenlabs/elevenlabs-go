package types

import (
	"fmt"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
type BodyTextToSpeechFullWithTimestampsApplyTextNormalization string

const (
	BodyTextToSpeechFullWithTimestampsApplyTextNormalizationAuto BodyTextToSpeechFullWithTimestampsApplyTextNormalization = "auto"
	BodyTextToSpeechFullWithTimestampsApplyTextNormalizationOn   BodyTextToSpeechFullWithTimestampsApplyTextNormalization = "on"
	BodyTextToSpeechFullWithTimestampsApplyTextNormalizationOff  BodyTextToSpeechFullWithTimestampsApplyTextNormalization = "off"
)

func NewBodyTextToSpeechFullWithTimestampsApplyTextNormalizationFromString(s string) (BodyTextToSpeechFullWithTimestampsApplyTextNormalization, error) {
	switch s {
	case "auto":
		return BodyTextToSpeechFullWithTimestampsApplyTextNormalizationAuto, nil
	case "on":
		return BodyTextToSpeechFullWithTimestampsApplyTextNormalizationOn, nil
	case "off":
		return BodyTextToSpeechFullWithTimestampsApplyTextNormalizationOff, nil
	}
	var t BodyTextToSpeechFullWithTimestampsApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyTextToSpeechFullWithTimestampsApplyTextNormalization) Ptr() *BodyTextToSpeechFullWithTimestampsApplyTextNormalization {
	return &e
}
