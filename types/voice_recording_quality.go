package types

import (
	"fmt"
)

type VoiceRecordingQuality string

const (
	VoiceRecordingQualityStudio VoiceRecordingQuality = "studio"
	VoiceRecordingQualityGood   VoiceRecordingQuality = "good"
	VoiceRecordingQualityOk     VoiceRecordingQuality = "ok"
	VoiceRecordingQualityPoor   VoiceRecordingQuality = "poor"
	VoiceRecordingQualityBad    VoiceRecordingQuality = "bad"
)

func NewVoiceRecordingQualityFromString(s string) (VoiceRecordingQuality, error) {
	switch s {
	case "studio":
		return VoiceRecordingQualityStudio, nil
	case "good":
		return VoiceRecordingQualityGood, nil
	case "ok":
		return VoiceRecordingQualityOk, nil
	case "poor":
		return VoiceRecordingQualityPoor, nil
	case "bad":
		return VoiceRecordingQualityBad, nil
	}
	var t VoiceRecordingQuality
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e VoiceRecordingQuality) Ptr() *VoiceRecordingQuality {
	return &e
}
