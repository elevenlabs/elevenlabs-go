package types

import (
	"fmt"
)

// Why a location is currently inactive.
type ReceptionLocationDeactivationReason string

const (
	ReceptionLocationDeactivationReasonTierDowngrade ReceptionLocationDeactivationReason = "tier_downgrade"
	ReceptionLocationDeactivationReasonUser          ReceptionLocationDeactivationReason = "user"
)

func NewReceptionLocationDeactivationReasonFromString(s string) (ReceptionLocationDeactivationReason, error) {
	switch s {
	case "tier_downgrade":
		return ReceptionLocationDeactivationReasonTierDowngrade, nil
	case "user":
		return ReceptionLocationDeactivationReasonUser, nil
	}
	var t ReceptionLocationDeactivationReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionLocationDeactivationReason) Ptr() *ReceptionLocationDeactivationReason {
	return &e
}
