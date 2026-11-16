package types

import (
	"fmt"
)

type SubscriptionStatusType string

const (
	SubscriptionStatusTypeTrialing     SubscriptionStatusType = "trialing"
	SubscriptionStatusTypeActive       SubscriptionStatusType = "active"
	SubscriptionStatusTypeIncomplete   SubscriptionStatusType = "incomplete"
	SubscriptionStatusTypePastDue      SubscriptionStatusType = "past_due"
	SubscriptionStatusTypeFree         SubscriptionStatusType = "free"
	SubscriptionStatusTypeFreeDisabled SubscriptionStatusType = "free_disabled"
)

func NewSubscriptionStatusTypeFromString(s string) (SubscriptionStatusType, error) {
	switch s {
	case "trialing":
		return SubscriptionStatusTypeTrialing, nil
	case "active":
		return SubscriptionStatusTypeActive, nil
	case "incomplete":
		return SubscriptionStatusTypeIncomplete, nil
	case "past_due":
		return SubscriptionStatusTypePastDue, nil
	case "free":
		return SubscriptionStatusTypeFree, nil
	case "free_disabled":
		return SubscriptionStatusTypeFreeDisabled, nil
	}
	var t SubscriptionStatusType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SubscriptionStatusType) Ptr() *SubscriptionStatusType {
	return &e
}
