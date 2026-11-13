package types

import (
	"fmt"
)

type WhatsAppAccountType string

const (
	WhatsAppAccountTypeCloudApi    WhatsAppAccountType = "cloud_api"
	WhatsAppAccountTypeCoexistence WhatsAppAccountType = "coexistence"
)

func NewWhatsAppAccountTypeFromString(s string) (WhatsAppAccountType, error) {
	switch s {
	case "cloud_api":
		return WhatsAppAccountTypeCloudApi, nil
	case "coexistence":
		return WhatsAppAccountTypeCoexistence, nil
	}
	var t WhatsAppAccountType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e WhatsAppAccountType) Ptr() *WhatsAppAccountType {
	return &e
}
