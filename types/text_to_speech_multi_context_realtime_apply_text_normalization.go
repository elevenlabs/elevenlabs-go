package types

import (
	"fmt"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
type TextToSpeechMultiContextRealtimeApplyTextNormalization string

const (
	TextToSpeechMultiContextRealtimeApplyTextNormalizationAuto TextToSpeechMultiContextRealtimeApplyTextNormalization = "auto"
	TextToSpeechMultiContextRealtimeApplyTextNormalizationOn   TextToSpeechMultiContextRealtimeApplyTextNormalization = "on"
	TextToSpeechMultiContextRealtimeApplyTextNormalizationOff  TextToSpeechMultiContextRealtimeApplyTextNormalization = "off"
)

func NewTextToSpeechMultiContextRealtimeApplyTextNormalizationFromString(s string) (TextToSpeechMultiContextRealtimeApplyTextNormalization, error) {
	switch s {
	case "auto":
		return TextToSpeechMultiContextRealtimeApplyTextNormalizationAuto, nil
	case "on":
		return TextToSpeechMultiContextRealtimeApplyTextNormalizationOn, nil
	case "off":
		return TextToSpeechMultiContextRealtimeApplyTextNormalizationOff, nil
	}
	var t TextToSpeechMultiContextRealtimeApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TextToSpeechMultiContextRealtimeApplyTextNormalization) Ptr() *TextToSpeechMultiContextRealtimeApplyTextNormalization {
	return &e
}
