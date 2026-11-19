package types

import (
	"fmt"
)

type ProvisionResellerWorkspaceSubscriptionBillingPeriod string

const (
	ProvisionResellerWorkspaceSubscriptionBillingPeriodMonthlyPeriod ProvisionResellerWorkspaceSubscriptionBillingPeriod = "monthly_period"
	ProvisionResellerWorkspaceSubscriptionBillingPeriodAnnualPeriod  ProvisionResellerWorkspaceSubscriptionBillingPeriod = "annual_period"
)

func NewProvisionResellerWorkspaceSubscriptionBillingPeriodFromString(s string) (ProvisionResellerWorkspaceSubscriptionBillingPeriod, error) {
	switch s {
	case "monthly_period":
		return ProvisionResellerWorkspaceSubscriptionBillingPeriodMonthlyPeriod, nil
	case "annual_period":
		return ProvisionResellerWorkspaceSubscriptionBillingPeriodAnnualPeriod, nil
	}
	var t ProvisionResellerWorkspaceSubscriptionBillingPeriod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ProvisionResellerWorkspaceSubscriptionBillingPeriod) Ptr() *ProvisionResellerWorkspaceSubscriptionBillingPeriod {
	return &e
}
