package types

import (
	"fmt"
)

// Current state of the project
type AudioNativeProjectSettingsResponseModelStatus string

const (
	AudioNativeProjectSettingsResponseModelStatusProcessing AudioNativeProjectSettingsResponseModelStatus = "processing"
	AudioNativeProjectSettingsResponseModelStatusReady      AudioNativeProjectSettingsResponseModelStatus = "ready"
)

func NewAudioNativeProjectSettingsResponseModelStatusFromString(s string) (AudioNativeProjectSettingsResponseModelStatus, error) {
	switch s {
	case "processing":
		return AudioNativeProjectSettingsResponseModelStatusProcessing, nil
	case "ready":
		return AudioNativeProjectSettingsResponseModelStatusReady, nil
	}
	var t AudioNativeProjectSettingsResponseModelStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e AudioNativeProjectSettingsResponseModelStatus) Ptr() *AudioNativeProjectSettingsResponseModelStatus {
	return &e
}
