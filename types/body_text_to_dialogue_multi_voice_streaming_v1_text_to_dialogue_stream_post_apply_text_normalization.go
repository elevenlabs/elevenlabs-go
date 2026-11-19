package types

import (
	"fmt"
)

// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
type BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalization string

const (
	BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalizationAuto BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalization = "auto"
	BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalizationOn   BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalization = "on"
	BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalizationOff  BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalization = "off"
)

func NewBodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalizationFromString(s string) (BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalization, error) {
	switch s {
	case "auto":
		return BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalizationAuto, nil
	case "on":
		return BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalizationOn, nil
	case "off":
		return BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalizationOff, nil
	}
	var t BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalization
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalization) Ptr() *BodyTextToDialogueMultiVoiceStreamingV1TextToDialogueStreamPostApplyTextNormalization {
	return &e
}
