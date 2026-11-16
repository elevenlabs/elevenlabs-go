package types

import (
	"fmt"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
type BodyTextToDialogueFullWithTimestampsApplyTextNormalization string

const (
	BodyTextToDialogueFullWithTimestampsApplyTextNormalizationAuto BodyTextToDialogueFullWithTimestampsApplyTextNormalization = "auto"
	BodyTextToDialogueFullWithTimestampsApplyTextNormalizationOn   BodyTextToDialogueFullWithTimestampsApplyTextNormalization = "on"
	BodyTextToDialogueFullWithTimestampsApplyTextNormalizationOff  BodyTextToDialogueFullWithTimestampsApplyTextNormalization = "off"
)

func NewBodyTextToDialogueFullWithTimestampsApplyTextNormalizationFromString(s string) (BodyTextToDialogueFullWithTimestampsApplyTextNormalization, error) {
	switch s {
	case "auto":
		return BodyTextToDialogueFullWithTimestampsApplyTextNormalizationAuto, nil
	case "on":
		return BodyTextToDialogueFullWithTimestampsApplyTextNormalizationOn, nil
	case "off":
		return BodyTextToDialogueFullWithTimestampsApplyTextNormalizationOff, nil
	}
	var t BodyTextToDialogueFullWithTimestampsApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyTextToDialogueFullWithTimestampsApplyTextNormalization) Ptr() *BodyTextToDialogueFullWithTimestampsApplyTextNormalization {
	return &e
}
