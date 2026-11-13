package types

import (
	"fmt"
)

// A string parameter to control text normalization.
type TextToDialogueMultiContextRealtimeApplyTextNormalization string

const (
	TextToDialogueMultiContextRealtimeApplyTextNormalizationAuto         TextToDialogueMultiContextRealtimeApplyTextNormalization = "auto"
	TextToDialogueMultiContextRealtimeApplyTextNormalizationOn           TextToDialogueMultiContextRealtimeApplyTextNormalization = "on"
	TextToDialogueMultiContextRealtimeApplyTextNormalizationOff          TextToDialogueMultiContextRealtimeApplyTextNormalization = "off"
	TextToDialogueMultiContextRealtimeApplyTextNormalizationApplyEnglish TextToDialogueMultiContextRealtimeApplyTextNormalization = "apply_english"
)

func NewTextToDialogueMultiContextRealtimeApplyTextNormalizationFromString(s string) (TextToDialogueMultiContextRealtimeApplyTextNormalization, error) {
	switch s {
	case "auto":
		return TextToDialogueMultiContextRealtimeApplyTextNormalizationAuto, nil
	case "on":
		return TextToDialogueMultiContextRealtimeApplyTextNormalizationOn, nil
	case "off":
		return TextToDialogueMultiContextRealtimeApplyTextNormalizationOff, nil
	case "apply_english":
		return TextToDialogueMultiContextRealtimeApplyTextNormalizationApplyEnglish, nil
	}
	var t TextToDialogueMultiContextRealtimeApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TextToDialogueMultiContextRealtimeApplyTextNormalization) Ptr() *TextToDialogueMultiContextRealtimeApplyTextNormalization {
	return &e
}
