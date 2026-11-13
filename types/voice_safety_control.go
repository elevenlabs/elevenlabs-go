package types

import (
	"fmt"
)

type VoiceSafetyControl string

const (
	VoiceSafetyControlNone              VoiceSafetyControl = "NONE"
	VoiceSafetyControlBan               VoiceSafetyControl = "BAN"
	VoiceSafetyControlCaptcha           VoiceSafetyControl = "CAPTCHA"
	VoiceSafetyControlEnterpriseBan     VoiceSafetyControl = "ENTERPRISE_BAN"
	VoiceSafetyControlEnterpriseCaptcha VoiceSafetyControl = "ENTERPRISE_CAPTCHA"
)

func NewVoiceSafetyControlFromString(s string) (VoiceSafetyControl, error) {
	switch s {
	case "NONE":
		return VoiceSafetyControlNone, nil
	case "BAN":
		return VoiceSafetyControlBan, nil
	case "CAPTCHA":
		return VoiceSafetyControlCaptcha, nil
	case "ENTERPRISE_BAN":
		return VoiceSafetyControlEnterpriseBan, nil
	case "ENTERPRISE_CAPTCHA":
		return VoiceSafetyControlEnterpriseCaptcha, nil
	}
	var t VoiceSafetyControl
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e VoiceSafetyControl) Ptr() *VoiceSafetyControl {
	return &e
}
