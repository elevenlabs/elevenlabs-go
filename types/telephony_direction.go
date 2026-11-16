package types

import (
	"fmt"
)

type TelephonyDirection string

const (
	TelephonyDirectionInbound  TelephonyDirection = "inbound"
	TelephonyDirectionOutbound TelephonyDirection = "outbound"
)

func NewTelephonyDirectionFromString(s string) (TelephonyDirection, error) {
	switch s {
	case "inbound":
		return TelephonyDirectionInbound, nil
	case "outbound":
		return TelephonyDirectionOutbound, nil
	}
	var t TelephonyDirection
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e TelephonyDirection) Ptr() *TelephonyDirection {
	return &e
}
