package types

import (
	"fmt"
)

// Which Twilio answering-machine-detection (AMD) mode to request for a call.
//
// `enable` returns a verdict as soon as Twilio can tell a human from a machine.
// `detect_message_end` additionally waits for the greeting to finish, which is what
// distinguishes the `machine_end_*` verdicts, at the cost of a later result.
type TwilioMachineDetectionMode string

const (
	TwilioMachineDetectionModeEnable           TwilioMachineDetectionMode = "enable"
	TwilioMachineDetectionModeDetectMessageEnd TwilioMachineDetectionMode = "detect_message_end"
)

func NewTwilioMachineDetectionModeFromString(s string) (TwilioMachineDetectionMode, error) {
	switch s {
	case "enable":
		return TwilioMachineDetectionModeEnable, nil
	case "detect_message_end":
		return TwilioMachineDetectionModeDetectMessageEnd, nil
	}
	var t TwilioMachineDetectionMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TwilioMachineDetectionMode) Ptr() *TwilioMachineDetectionMode {
	return &e
}
