package types

import (
	"fmt"
)

// The category of the voice.
type LibraryVoiceResponseCategory string

const (
	LibraryVoiceResponseCategoryGenerated    LibraryVoiceResponseCategory = "generated"
	LibraryVoiceResponseCategoryCloned       LibraryVoiceResponseCategory = "cloned"
	LibraryVoiceResponseCategoryPremade      LibraryVoiceResponseCategory = "premade"
	LibraryVoiceResponseCategoryProfessional LibraryVoiceResponseCategory = "professional"
	LibraryVoiceResponseCategoryFamous       LibraryVoiceResponseCategory = "famous"
	LibraryVoiceResponseCategoryHighQuality  LibraryVoiceResponseCategory = "high_quality"
)

func NewLibraryVoiceResponseCategoryFromString(s string) (LibraryVoiceResponseCategory, error) {
	switch s {
	case "generated":
		return LibraryVoiceResponseCategoryGenerated, nil
	case "cloned":
		return LibraryVoiceResponseCategoryCloned, nil
	case "premade":
		return LibraryVoiceResponseCategoryPremade, nil
	case "professional":
		return LibraryVoiceResponseCategoryProfessional, nil
	case "famous":
		return LibraryVoiceResponseCategoryFamous, nil
	case "high_quality":
		return LibraryVoiceResponseCategoryHighQuality, nil
	}
	var t LibraryVoiceResponseCategory
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e LibraryVoiceResponseCategory) Ptr() *LibraryVoiceResponseCategory {
	return &e
}
