package types

import (
	"fmt"
)

type ReceptionNotificationChannel string

const (
	ReceptionNotificationChannelEmail ReceptionNotificationChannel = "email"
	ReceptionNotificationChannelSms   ReceptionNotificationChannel = "sms"
)

func NewReceptionNotificationChannelFromString(s string) (ReceptionNotificationChannel, error) {
	switch s {
	case "email":
		return ReceptionNotificationChannelEmail, nil
	case "sms":
		return ReceptionNotificationChannelSms, nil
	}
	var t ReceptionNotificationChannel
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionNotificationChannel) Ptr() *ReceptionNotificationChannel {
	return &e
}
