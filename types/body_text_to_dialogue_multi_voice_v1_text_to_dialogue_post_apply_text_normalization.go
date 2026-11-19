package types

import (
	"fmt"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
type BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalization string

const (
	BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalizationAuto BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalization = "auto"
	BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalizationOn   BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalization = "on"
	BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalizationOff  BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalization = "off"
)

func NewBodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalizationFromString(s string) (BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalization, error) {
	switch s {
	case "auto":
		return BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalizationAuto, nil
	case "on":
		return BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalizationOn, nil
	case "off":
		return BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalizationOff, nil
	}
	var t BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalization) Ptr() *BodyTextToDialogueMultiVoiceV1TextToDialoguePostApplyTextNormalization {
	return &e
}
