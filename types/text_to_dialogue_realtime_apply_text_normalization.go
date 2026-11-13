package types

import (
	"fmt"
)

// A string parameter to control text normalization.
type TextToDialogueRealtimeApplyTextNormalization string

const (
	TextToDialogueRealtimeApplyTextNormalizationAuto         TextToDialogueRealtimeApplyTextNormalization = "auto"
	TextToDialogueRealtimeApplyTextNormalizationOn           TextToDialogueRealtimeApplyTextNormalization = "on"
	TextToDialogueRealtimeApplyTextNormalizationOff          TextToDialogueRealtimeApplyTextNormalization = "off"
	TextToDialogueRealtimeApplyTextNormalizationApplyEnglish TextToDialogueRealtimeApplyTextNormalization = "apply_english"
)

func NewTextToDialogueRealtimeApplyTextNormalizationFromString(s string) (TextToDialogueRealtimeApplyTextNormalization, error) {
	switch s {
	case "auto":
		return TextToDialogueRealtimeApplyTextNormalizationAuto, nil
	case "on":
		return TextToDialogueRealtimeApplyTextNormalizationOn, nil
	case "off":
		return TextToDialogueRealtimeApplyTextNormalizationOff, nil
	case "apply_english":
		return TextToDialogueRealtimeApplyTextNormalizationApplyEnglish, nil
	}
	var t TextToDialogueRealtimeApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TextToDialogueRealtimeApplyTextNormalization) Ptr() *TextToDialogueRealtimeApplyTextNormalization {
	return &e
}
