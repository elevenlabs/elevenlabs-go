package types

import (
	"fmt"
)

type OutboundMode string

const (
	OutboundModeVoice   OutboundMode = "voice"
	OutboundModeMessage OutboundMode = "message"
)

func NewOutboundModeFromString(s string) (OutboundMode, error) {
	switch s {
	case "voice":
		return OutboundModeVoice, nil
	case "message":
		return OutboundModeMessage, nil
	}
	var t OutboundMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e OutboundMode) Ptr() *OutboundMode {
	return &e
}
