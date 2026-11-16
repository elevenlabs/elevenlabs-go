package types

import (
	"fmt"
)

type LockReason string

const (
	LockReasonTrialEnded            LockReason = "trial_ended"
	LockReasonSubscriptionDowngrade LockReason = "subscription_downgrade"
	LockReasonExposedPublicly       LockReason = "exposed_publicly"
	LockReasonSelfDisabled          LockReason = "self_disabled"
)

func NewLockReasonFromString(s string) (LockReason, error) {
	switch s {
	case "trial_ended":
		return LockReasonTrialEnded, nil
	case "subscription_downgrade":
		return LockReasonSubscriptionDowngrade, nil
	case "exposed_publicly":
		return LockReasonExposedPublicly, nil
	case "self_disabled":
		return LockReasonSelfDisabled, nil
	}
	var t LockReason
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e LockReason) Ptr() *LockReason {
	return &e
}
