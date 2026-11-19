package types

import (
	"fmt"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
type TranslationRealtimeApplyTextNormalization string

const (
	TranslationRealtimeApplyTextNormalizationAuto TranslationRealtimeApplyTextNormalization = "auto"
	TranslationRealtimeApplyTextNormalizationOn   TranslationRealtimeApplyTextNormalization = "on"
	TranslationRealtimeApplyTextNormalizationOff  TranslationRealtimeApplyTextNormalization = "off"
)

func NewTranslationRealtimeApplyTextNormalizationFromString(s string) (TranslationRealtimeApplyTextNormalization, error) {
	switch s {
	case "auto":
		return TranslationRealtimeApplyTextNormalizationAuto, nil
	case "on":
		return TranslationRealtimeApplyTextNormalizationOn, nil
	case "off":
		return TranslationRealtimeApplyTextNormalizationOff, nil
	}
	var t TranslationRealtimeApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TranslationRealtimeApplyTextNormalization) Ptr() *TranslationRealtimeApplyTextNormalization {
	return &e
}
