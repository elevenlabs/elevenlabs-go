package types

import (
	"fmt"
)

type SipTrunkTransportEnum string

const (
	SipTrunkTransportEnumAuto SipTrunkTransportEnum = "auto"
	SipTrunkTransportEnumUdp  SipTrunkTransportEnum = "udp"
	SipTrunkTransportEnumTcp  SipTrunkTransportEnum = "tcp"
	SipTrunkTransportEnumTls  SipTrunkTransportEnum = "tls"
)

func NewSipTrunkTransportEnumFromString(s string) (SipTrunkTransportEnum, error) {
	switch s {
	case "auto":
		return SipTrunkTransportEnumAuto, nil
	case "udp":
		return SipTrunkTransportEnumUdp, nil
	case "tcp":
		return SipTrunkTransportEnumTcp, nil
	case "tls":
		return SipTrunkTransportEnumTls, nil
	}
	var t SipTrunkTransportEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SipTrunkTransportEnum) Ptr() *SipTrunkTransportEnum {
	return &e
}
