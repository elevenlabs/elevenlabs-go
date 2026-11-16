package types

import (
	"fmt"
)

type SfxModelId string

const (
	SfxModelIdElevenTextToSoundV2 SfxModelId = "eleven_text_to_sound_v2"
	SfxModelIdElevenTextToSoundV3 SfxModelId = "eleven_text_to_sound_v3"
)

func NewSfxModelIdFromString(s string) (SfxModelId, error) {
	switch s {
	case "eleven_text_to_sound_v2":
		return SfxModelIdElevenTextToSoundV2, nil
	case "eleven_text_to_sound_v3":
		return SfxModelIdElevenTextToSoundV3, nil
	}
	var t SfxModelId
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SfxModelId) Ptr() *SfxModelId {
	return &e
}
