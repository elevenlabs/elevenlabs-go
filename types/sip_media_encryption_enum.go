package types

import (
	"fmt"
)

type SipMediaEncryptionEnum string

const (
	SipMediaEncryptionEnumDisabled SipMediaEncryptionEnum = "disabled"
	SipMediaEncryptionEnumAllowed  SipMediaEncryptionEnum = "allowed"
	SipMediaEncryptionEnumRequired SipMediaEncryptionEnum = "required"
)

func NewSipMediaEncryptionEnumFromString(s string) (SipMediaEncryptionEnum, error) {
	switch s {
	case "disabled":
		return SipMediaEncryptionEnumDisabled, nil
	case "allowed":
		return SipMediaEncryptionEnumAllowed, nil
	case "required":
		return SipMediaEncryptionEnumRequired, nil
	}
	var t SipMediaEncryptionEnum
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SipMediaEncryptionEnum) Ptr() *SipMediaEncryptionEnum {
	return &e
}
