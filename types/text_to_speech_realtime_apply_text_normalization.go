package types

import (
	"fmt"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
type TextToSpeechRealtimeApplyTextNormalization string

const (
	TextToSpeechRealtimeApplyTextNormalizationAuto TextToSpeechRealtimeApplyTextNormalization = "auto"
	TextToSpeechRealtimeApplyTextNormalizationOn   TextToSpeechRealtimeApplyTextNormalization = "on"
	TextToSpeechRealtimeApplyTextNormalizationOff  TextToSpeechRealtimeApplyTextNormalization = "off"
)

func NewTextToSpeechRealtimeApplyTextNormalizationFromString(s string) (TextToSpeechRealtimeApplyTextNormalization, error) {
	switch s {
	case "auto":
		return TextToSpeechRealtimeApplyTextNormalizationAuto, nil
	case "on":
		return TextToSpeechRealtimeApplyTextNormalizationOn, nil
	case "off":
		return TextToSpeechRealtimeApplyTextNormalizationOff, nil
	}
	var t TextToSpeechRealtimeApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TextToSpeechRealtimeApplyTextNormalization) Ptr() *TextToSpeechRealtimeApplyTextNormalization {
	return &e
}
