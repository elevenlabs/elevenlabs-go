package types

import (
	"fmt"
)

type WhatsAppConversationInfoDirection string

const (
	WhatsAppConversationInfoDirectionInbound  WhatsAppConversationInfoDirection = "inbound"
	WhatsAppConversationInfoDirectionOutbound WhatsAppConversationInfoDirection = "outbound"
	WhatsAppConversationInfoDirectionUnknown  WhatsAppConversationInfoDirection = "unknown"
)

func NewWhatsAppConversationInfoDirectionFromString(s string) (WhatsAppConversationInfoDirection, error) {
	switch s {
	case "inbound":
		return WhatsAppConversationInfoDirectionInbound, nil
	case "outbound":
		return WhatsAppConversationInfoDirectionOutbound, nil
	case "unknown":
		return WhatsAppConversationInfoDirectionUnknown, nil
	}
	var t WhatsAppConversationInfoDirection
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WhatsAppConversationInfoDirection) Ptr() *WhatsAppConversationInfoDirection {
	return &e
}
