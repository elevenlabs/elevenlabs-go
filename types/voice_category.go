package types

import (
	"fmt"
)

type VoiceCategory string

const (
	VoiceCategoryPremade      VoiceCategory = "premade"
	VoiceCategoryCloned       VoiceCategory = "cloned"
	VoiceCategoryGenerated    VoiceCategory = "generated"
	VoiceCategoryProfessional VoiceCategory = "professional"
)

func NewVoiceCategoryFromString(s string) (VoiceCategory, error) {
	switch s {
	case "premade":
		return VoiceCategoryPremade, nil
	case "cloned":
		return VoiceCategoryCloned, nil
	case "generated":
		return VoiceCategoryGenerated, nil
	case "professional":
		return VoiceCategoryProfessional, nil
	}
	var t VoiceCategory
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e VoiceCategory) Ptr() *VoiceCategory {
	return &e
}
