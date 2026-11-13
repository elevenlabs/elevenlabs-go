package types

import (
	"fmt"
)

type WhatsAppCallPermissionRequestState string

const (
	WhatsAppCallPermissionRequestStateBlocked  WhatsAppCallPermissionRequestState = "blocked"
	WhatsAppCallPermissionRequestStateSent     WhatsAppCallPermissionRequestState = "sent"
	WhatsAppCallPermissionRequestStateAccepted WhatsAppCallPermissionRequestState = "accepted"
	WhatsAppCallPermissionRequestStateDeclined WhatsAppCallPermissionRequestState = "declined"
)

func NewWhatsAppCallPermissionRequestStateFromString(s string) (WhatsAppCallPermissionRequestState, error) {
	switch s {
	case "blocked":
		return WhatsAppCallPermissionRequestStateBlocked, nil
	case "sent":
		return WhatsAppCallPermissionRequestStateSent, nil
	case "accepted":
		return WhatsAppCallPermissionRequestStateAccepted, nil
	case "declined":
		return WhatsAppCallPermissionRequestStateDeclined, nil
	}
	var t WhatsAppCallPermissionRequestState
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WhatsAppCallPermissionRequestState) Ptr() *WhatsAppCallPermissionRequestState {
	return &e
}
