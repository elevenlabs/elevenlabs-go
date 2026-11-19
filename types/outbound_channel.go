package types

import (
	"fmt"
)

type OutboundChannel string

const (
	OutboundChannelPhone    OutboundChannel = "phone"
	OutboundChannelWhatsapp OutboundChannel = "whatsapp"
)

func NewOutboundChannelFromString(s string) (OutboundChannel, error) {
	switch s {
	case "phone":
		return OutboundChannelPhone, nil
	case "whatsapp":
		return OutboundChannelWhatsapp, nil
	}
	var t OutboundChannel
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e OutboundChannel) Ptr() *OutboundChannel {
	return &e
}
