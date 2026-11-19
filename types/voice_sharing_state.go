package types

import (
	"fmt"
)

// The status of the voice sharing.
type VoiceSharingState string

const (
	VoiceSharingStateEnabled        VoiceSharingState = "enabled"
	VoiceSharingStateDisabled       VoiceSharingState = "disabled"
	VoiceSharingStateCopied         VoiceSharingState = "copied"
	VoiceSharingStateCopiedDisabled VoiceSharingState = "copied_disabled"
)

func NewVoiceSharingStateFromString(s string) (VoiceSharingState, error) {
	switch s {
	case "enabled":
		return VoiceSharingStateEnabled, nil
	case "disabled":
		return VoiceSharingStateDisabled, nil
	case "copied":
		return VoiceSharingStateCopied, nil
	case "copied_disabled":
		return VoiceSharingStateCopiedDisabled, nil
	}
	var t VoiceSharingState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e VoiceSharingState) Ptr() *VoiceSharingState {
	return &e
}
