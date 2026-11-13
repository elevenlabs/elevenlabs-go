package types

import (
	"fmt"
)

// Carrier-level AMD verdict reported by Twilio via `AnsweredBy`.
type TwilioAnsweredBy string

const (
	TwilioAnsweredByHuman             TwilioAnsweredBy = "human"
	TwilioAnsweredByFax               TwilioAnsweredBy = "fax"
	TwilioAnsweredByUnknown           TwilioAnsweredBy = "unknown"
	TwilioAnsweredByMachineStart      TwilioAnsweredBy = "machine_start"
	TwilioAnsweredByMachineEndBeep    TwilioAnsweredBy = "machine_end_beep"
	TwilioAnsweredByMachineEndSilence TwilioAnsweredBy = "machine_end_silence"
	TwilioAnsweredByMachineEndOther   TwilioAnsweredBy = "machine_end_other"
)

func NewTwilioAnsweredByFromString(s string) (TwilioAnsweredBy, error) {
	switch s {
	case "human":
		return TwilioAnsweredByHuman, nil
	case "fax":
		return TwilioAnsweredByFax, nil
	case "unknown":
		return TwilioAnsweredByUnknown, nil
	case "machine_start":
		return TwilioAnsweredByMachineStart, nil
	case "machine_end_beep":
		return TwilioAnsweredByMachineEndBeep, nil
	case "machine_end_silence":
		return TwilioAnsweredByMachineEndSilence, nil
	case "machine_end_other":
		return TwilioAnsweredByMachineEndOther, nil
	}
	var t TwilioAnsweredBy
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TwilioAnsweredBy) Ptr() *TwilioAnsweredBy {
	return &e
}
