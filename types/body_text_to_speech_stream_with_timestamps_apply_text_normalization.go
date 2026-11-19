package types

import (
	"fmt"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
type BodyTextToSpeechStreamWithTimestampsApplyTextNormalization string

const (
	BodyTextToSpeechStreamWithTimestampsApplyTextNormalizationAuto BodyTextToSpeechStreamWithTimestampsApplyTextNormalization = "auto"
	BodyTextToSpeechStreamWithTimestampsApplyTextNormalizationOn   BodyTextToSpeechStreamWithTimestampsApplyTextNormalization = "on"
	BodyTextToSpeechStreamWithTimestampsApplyTextNormalizationOff  BodyTextToSpeechStreamWithTimestampsApplyTextNormalization = "off"
)

func NewBodyTextToSpeechStreamWithTimestampsApplyTextNormalizationFromString(s string) (BodyTextToSpeechStreamWithTimestampsApplyTextNormalization, error) {
	switch s {
	case "auto":
		return BodyTextToSpeechStreamWithTimestampsApplyTextNormalizationAuto, nil
	case "on":
		return BodyTextToSpeechStreamWithTimestampsApplyTextNormalizationOn, nil
	case "off":
		return BodyTextToSpeechStreamWithTimestampsApplyTextNormalizationOff, nil
	}
	var t BodyTextToSpeechStreamWithTimestampsApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyTextToSpeechStreamWithTimestampsApplyTextNormalization) Ptr() *BodyTextToSpeechStreamWithTimestampsApplyTextNormalization {
	return &e
}
