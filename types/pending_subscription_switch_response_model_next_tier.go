package types

import (
	"fmt"
)

// The tier to change to.
type PendingSubscriptionSwitchResponseModelNextTier string

const (
	PendingSubscriptionSwitchResponseModelNextTierFree               PendingSubscriptionSwitchResponseModelNextTier = "free"
	PendingSubscriptionSwitchResponseModelNextTierStarter            PendingSubscriptionSwitchResponseModelNextTier = "starter"
	PendingSubscriptionSwitchResponseModelNextTierGo                 PendingSubscriptionSwitchResponseModelNextTier = "go"
	PendingSubscriptionSwitchResponseModelNextTierCreator            PendingSubscriptionSwitchResponseModelNextTier = "creator"
	PendingSubscriptionSwitchResponseModelNextTierPro                PendingSubscriptionSwitchResponseModelNextTier = "pro"
	PendingSubscriptionSwitchResponseModelNextTierGrowingBusiness    PendingSubscriptionSwitchResponseModelNextTier = "growing_business"
	PendingSubscriptionSwitchResponseModelNextTierScale20240810      PendingSubscriptionSwitchResponseModelNextTier = "scale_2024_08_10"
	PendingSubscriptionSwitchResponseModelNextTierGrantTier120250723 PendingSubscriptionSwitchResponseModelNextTier = "grant_tier_1_2025_07_23"
	PendingSubscriptionSwitchResponseModelNextTierGrantTier220250723 PendingSubscriptionSwitchResponseModelNextTier = "grant_tier_2_2025_07_23"
	PendingSubscriptionSwitchResponseModelNextTierGrant              PendingSubscriptionSwitchResponseModelNextTier = "grant"
	PendingSubscriptionSwitchResponseModelNextTierTrial              PendingSubscriptionSwitchResponseModelNextTier = "trial"
	PendingSubscriptionSwitchResponseModelNextTierEnterprise         PendingSubscriptionSwitchResponseModelNextTier = "enterprise"
)

func NewPendingSubscriptionSwitchResponseModelNextTierFromString(s string) (PendingSubscriptionSwitchResponseModelNextTier, error) {
	switch s {
	case "free":
		return PendingSubscriptionSwitchResponseModelNextTierFree, nil
	case "starter":
		return PendingSubscriptionSwitchResponseModelNextTierStarter, nil
	case "go":
		return PendingSubscriptionSwitchResponseModelNextTierGo, nil
	case "creator":
		return PendingSubscriptionSwitchResponseModelNextTierCreator, nil
	case "pro":
		return PendingSubscriptionSwitchResponseModelNextTierPro, nil
	case "growing_business":
		return PendingSubscriptionSwitchResponseModelNextTierGrowingBusiness, nil
	case "scale_2024_08_10":
		return PendingSubscriptionSwitchResponseModelNextTierScale20240810, nil
	case "grant_tier_1_2025_07_23":
		return PendingSubscriptionSwitchResponseModelNextTierGrantTier120250723, nil
	case "grant_tier_2_2025_07_23":
		return PendingSubscriptionSwitchResponseModelNextTierGrantTier220250723, nil
	case "grant":
		return PendingSubscriptionSwitchResponseModelNextTierGrant, nil
	case "trial":
		return PendingSubscriptionSwitchResponseModelNextTierTrial, nil
	case "enterprise":
		return PendingSubscriptionSwitchResponseModelNextTierEnterprise, nil
	}
	var t PendingSubscriptionSwitchResponseModelNextTier
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e PendingSubscriptionSwitchResponseModelNextTier) Ptr() *PendingSubscriptionSwitchResponseModelNextTier {
	return &e
}
