package types

import (
	"fmt"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
type BodyTextToDialogueStreamWithTimestampsApplyTextNormalization string

const (
	BodyTextToDialogueStreamWithTimestampsApplyTextNormalizationAuto BodyTextToDialogueStreamWithTimestampsApplyTextNormalization = "auto"
	BodyTextToDialogueStreamWithTimestampsApplyTextNormalizationOn   BodyTextToDialogueStreamWithTimestampsApplyTextNormalization = "on"
	BodyTextToDialogueStreamWithTimestampsApplyTextNormalizationOff  BodyTextToDialogueStreamWithTimestampsApplyTextNormalization = "off"
)

func NewBodyTextToDialogueStreamWithTimestampsApplyTextNormalizationFromString(s string) (BodyTextToDialogueStreamWithTimestampsApplyTextNormalization, error) {
	switch s {
	case "auto":
		return BodyTextToDialogueStreamWithTimestampsApplyTextNormalizationAuto, nil
	case "on":
		return BodyTextToDialogueStreamWithTimestampsApplyTextNormalizationOn, nil
	case "off":
		return BodyTextToDialogueStreamWithTimestampsApplyTextNormalizationOff, nil
	}
	var t BodyTextToDialogueStreamWithTimestampsApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyTextToDialogueStreamWithTimestampsApplyTextNormalization) Ptr() *BodyTextToDialogueStreamWithTimestampsApplyTextNormalization {
	return &e
}
