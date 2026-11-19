package types

import (
	"fmt"
)

// Model to use for the voice generation. Possible values: eleven_multilingual_ttv_v2, eleven_ttv_v3.
type VoiceDesignRequestModelModelId string

const (
	VoiceDesignRequestModelModelIdElevenMultilingualTtvV2 VoiceDesignRequestModelModelId = "eleven_multilingual_ttv_v2"
	VoiceDesignRequestModelModelIdElevenTtvV3             VoiceDesignRequestModelModelId = "eleven_ttv_v3"
)

func NewVoiceDesignRequestModelModelIdFromString(s string) (VoiceDesignRequestModelModelId, error) {
	switch s {
	case "eleven_multilingual_ttv_v2":
		return VoiceDesignRequestModelModelIdElevenMultilingualTtvV2, nil
	case "eleven_ttv_v3":
		return VoiceDesignRequestModelModelIdElevenTtvV3, nil
	}
	var t VoiceDesignRequestModelModelId
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e VoiceDesignRequestModelModelId) Ptr() *VoiceDesignRequestModelModelId {
	return &e
}
