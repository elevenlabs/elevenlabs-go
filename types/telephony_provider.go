package types

import (
	"fmt"
)

type TelephonyProvider string

const (
	TelephonyProviderTwilio   TelephonyProvider = "twilio"
	TelephonyProviderSipTrunk TelephonyProvider = "sip_trunk"
	TelephonyProviderExotel   TelephonyProvider = "exotel"
)

func NewTelephonyProviderFromString(s string) (TelephonyProvider, error) {
	switch s {
	case "twilio":
		return TelephonyProviderTwilio, nil
	case "sip_trunk":
		return TelephonyProviderSipTrunk, nil
	case "exotel":
		return TelephonyProviderExotel, nil
	}
	var t TelephonyProvider
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TelephonyProvider) Ptr() *TelephonyProvider {
	return &e
}
