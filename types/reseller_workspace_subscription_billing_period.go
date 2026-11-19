package types

import (
	"fmt"
)

type ResellerWorkspaceSubscriptionBillingPeriod string

const (
	ResellerWorkspaceSubscriptionBillingPeriodMonthlyPeriod ResellerWorkspaceSubscriptionBillingPeriod = "monthly_period"
	ResellerWorkspaceSubscriptionBillingPeriodAnnualPeriod  ResellerWorkspaceSubscriptionBillingPeriod = "annual_period"
)

func NewResellerWorkspaceSubscriptionBillingPeriodFromString(s string) (ResellerWorkspaceSubscriptionBillingPeriod, error) {
	switch s {
	case "monthly_period":
		return ResellerWorkspaceSubscriptionBillingPeriodMonthlyPeriod, nil
	case "annual_period":
		return ResellerWorkspaceSubscriptionBillingPeriodAnnualPeriod, nil
	}
	var t ResellerWorkspaceSubscriptionBillingPeriod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ResellerWorkspaceSubscriptionBillingPeriod) Ptr() *ResellerWorkspaceSubscriptionBillingPeriod {
	return &e
}
