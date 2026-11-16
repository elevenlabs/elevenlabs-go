package types

import (
	"fmt"
)

// Why a receptionist is currently inactive (live calls disabled).
type ReceptionCustomerFacingAgentDeactivationReason string

const (
	ReceptionCustomerFacingAgentDeactivationReasonTierDowngrade ReceptionCustomerFacingAgentDeactivationReason = "tier_downgrade"
	ReceptionCustomerFacingAgentDeactivationReasonUser          ReceptionCustomerFacingAgentDeactivationReason = "user"
)

func NewReceptionCustomerFacingAgentDeactivationReasonFromString(s string) (ReceptionCustomerFacingAgentDeactivationReason, error) {
	switch s {
	case "tier_downgrade":
		return ReceptionCustomerFacingAgentDeactivationReasonTierDowngrade, nil
	case "user":
		return ReceptionCustomerFacingAgentDeactivationReasonUser, nil
	}
	var t ReceptionCustomerFacingAgentDeactivationReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionCustomerFacingAgentDeactivationReason) Ptr() *ReceptionCustomerFacingAgentDeactivationReason {
	return &e
}
