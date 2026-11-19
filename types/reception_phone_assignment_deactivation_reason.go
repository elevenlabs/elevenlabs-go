package types

import (
	"fmt"
)

// Why a phone assignment is currently inactive.
type ReceptionPhoneAssignmentDeactivationReason string

const (
	ReceptionPhoneAssignmentDeactivationReasonTierDowngrade ReceptionPhoneAssignmentDeactivationReason = "tier_downgrade"
	ReceptionPhoneAssignmentDeactivationReasonNoEntitlement ReceptionPhoneAssignmentDeactivationReason = "no_entitlement"
	ReceptionPhoneAssignmentDeactivationReasonUser          ReceptionPhoneAssignmentDeactivationReason = "user"
)

func NewReceptionPhoneAssignmentDeactivationReasonFromString(s string) (ReceptionPhoneAssignmentDeactivationReason, error) {
	switch s {
	case "tier_downgrade":
		return ReceptionPhoneAssignmentDeactivationReasonTierDowngrade, nil
	case "no_entitlement":
		return ReceptionPhoneAssignmentDeactivationReasonNoEntitlement, nil
	case "user":
		return ReceptionPhoneAssignmentDeactivationReasonUser, nil
	}
	var t ReceptionPhoneAssignmentDeactivationReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionPhoneAssignmentDeactivationReason) Ptr() *ReceptionPhoneAssignmentDeactivationReason {
	return &e
}
