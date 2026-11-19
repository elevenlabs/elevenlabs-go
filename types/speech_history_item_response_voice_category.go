package types

import (
	"fmt"
)

type SpeechHistoryItemResponseVoiceCategory string

const (
	SpeechHistoryItemResponseVoiceCategoryPremade      SpeechHistoryItemResponseVoiceCategory = "premade"
	SpeechHistoryItemResponseVoiceCategoryCloned       SpeechHistoryItemResponseVoiceCategory = "cloned"
	SpeechHistoryItemResponseVoiceCategoryGenerated    SpeechHistoryItemResponseVoiceCategory = "generated"
	SpeechHistoryItemResponseVoiceCategoryProfessional SpeechHistoryItemResponseVoiceCategory = "professional"
)

func NewSpeechHistoryItemResponseVoiceCategoryFromString(s string) (SpeechHistoryItemResponseVoiceCategory, error) {
	switch s {
	case "premade":
		return SpeechHistoryItemResponseVoiceCategoryPremade, nil
	case "cloned":
		return SpeechHistoryItemResponseVoiceCategoryCloned, nil
	case "generated":
		return SpeechHistoryItemResponseVoiceCategoryGenerated, nil
	case "professional":
		return SpeechHistoryItemResponseVoiceCategoryProfessional, nil
	}
	var t SpeechHistoryItemResponseVoiceCategory
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SpeechHistoryItemResponseVoiceCategory) Ptr() *SpeechHistoryItemResponseVoiceCategory {
	return &e
}
