package types

import (
	"fmt"
)

// Regulatory standing of a number in a country that requires paperwork.
//
// Null for numbers that need no paperwork, such as US and Canadian numbers.
type ReceptionPhoneRegulatoryStatus string

const (
	ReceptionPhoneRegulatoryStatusHeld     ReceptionPhoneRegulatoryStatus = "held"
	ReceptionPhoneRegulatoryStatusApproved ReceptionPhoneRegulatoryStatus = "approved"
)

func NewReceptionPhoneRegulatoryStatusFromString(s string) (ReceptionPhoneRegulatoryStatus, error) {
	switch s {
	case "held":
		return ReceptionPhoneRegulatoryStatusHeld, nil
	case "approved":
		return ReceptionPhoneRegulatoryStatusApproved, nil
	}
	var t ReceptionPhoneRegulatoryStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionPhoneRegulatoryStatus) Ptr() *ReceptionPhoneRegulatoryStatus {
	return &e
}
