package types

import (
	"fmt"
)

// The category of the voice.
type VoiceSharingResponseCategory string

const (
	VoiceSharingResponseCategoryGenerated    VoiceSharingResponseCategory = "generated"
	VoiceSharingResponseCategoryCloned       VoiceSharingResponseCategory = "cloned"
	VoiceSharingResponseCategoryPremade      VoiceSharingResponseCategory = "premade"
	VoiceSharingResponseCategoryProfessional VoiceSharingResponseCategory = "professional"
	VoiceSharingResponseCategoryFamous       VoiceSharingResponseCategory = "famous"
	VoiceSharingResponseCategoryHighQuality  VoiceSharingResponseCategory = "high_quality"
)

func NewVoiceSharingResponseCategoryFromString(s string) (VoiceSharingResponseCategory, error) {
	switch s {
	case "generated":
		return VoiceSharingResponseCategoryGenerated, nil
	case "cloned":
		return VoiceSharingResponseCategoryCloned, nil
	case "premade":
		return VoiceSharingResponseCategoryPremade, nil
	case "professional":
		return VoiceSharingResponseCategoryProfessional, nil
	case "famous":
		return VoiceSharingResponseCategoryFamous, nil
	case "high_quality":
		return VoiceSharingResponseCategoryHighQuality, nil
	}
	var t VoiceSharingResponseCategory
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e VoiceSharingResponseCategory) Ptr() *VoiceSharingResponseCategory {
	return &e
}
