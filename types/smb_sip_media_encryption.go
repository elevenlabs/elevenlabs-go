package types

import (
	"fmt"
)

// Media (SRTP) encryption policy for a SIP trunk.
//
// Values match the Agents platform “SIPMediaEncryptionEnum“ so the SMB
// adapter maps by value.
type SmbSipMediaEncryption string

const (
	SmbSipMediaEncryptionDisabled SmbSipMediaEncryption = "disabled"
	SmbSipMediaEncryptionAllowed  SmbSipMediaEncryption = "allowed"
	SmbSipMediaEncryptionRequired SmbSipMediaEncryption = "required"
)

func NewSmbSipMediaEncryptionFromString(s string) (SmbSipMediaEncryption, error) {
	switch s {
	case "disabled":
		return SmbSipMediaEncryptionDisabled, nil
	case "allowed":
		return SmbSipMediaEncryptionAllowed, nil
	case "required":
		return SmbSipMediaEncryptionRequired, nil
	}
	var t SmbSipMediaEncryption
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SmbSipMediaEncryption) Ptr() *SmbSipMediaEncryption {
	return &e
}
