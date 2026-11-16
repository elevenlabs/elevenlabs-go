package types

import (
	"fmt"
)

type SmsConversationInfoDirection string

const (
	SmsConversationInfoDirectionInbound  SmsConversationInfoDirection = "inbound"
	SmsConversationInfoDirectionOutbound SmsConversationInfoDirection = "outbound"
)

func NewSmsConversationInfoDirectionFromString(s string) (SmsConversationInfoDirection, error) {
	switch s {
	case "inbound":
		return SmsConversationInfoDirectionInbound, nil
	case "outbound":
		return SmsConversationInfoDirectionOutbound, nil
	}
	var t SmsConversationInfoDirection
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SmsConversationInfoDirection) Ptr() *SmsConversationInfoDirection {
	return &e
}
