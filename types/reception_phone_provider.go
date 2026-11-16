package types

import (
	"fmt"
)

// How calls to the number reach ElevenLabs: through Twilio or a SIP trunk.
type ReceptionPhoneProvider string

const (
	ReceptionPhoneProviderTwilio   ReceptionPhoneProvider = "twilio"
	ReceptionPhoneProviderSipTrunk ReceptionPhoneProvider = "sip_trunk"
)

func NewReceptionPhoneProviderFromString(s string) (ReceptionPhoneProvider, error) {
	switch s {
	case "twilio":
		return ReceptionPhoneProviderTwilio, nil
	case "sip_trunk":
		return ReceptionPhoneProviderSipTrunk, nil
	}
	var t ReceptionPhoneProvider
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionPhoneProvider) Ptr() *ReceptionPhoneProvider {
	return &e
}
