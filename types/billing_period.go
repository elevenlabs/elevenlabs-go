package types

import (
	"fmt"
)

type BillingPeriod string

const (
	BillingPeriodMonthlyPeriod    BillingPeriod = "monthly_period"
	BillingPeriodThreeMonthPeriod BillingPeriod = "3_month_period"
	BillingPeriodSixMonthPeriod   BillingPeriod = "6_month_period"
	BillingPeriodAnnualPeriod     BillingPeriod = "annual_period"
)

func NewBillingPeriodFromString(s string) (BillingPeriod, error) {
	switch s {
	case "monthly_period":
		return BillingPeriodMonthlyPeriod, nil
	case "3_month_period":
		return BillingPeriodThreeMonthPeriod, nil
	case "6_month_period":
		return BillingPeriodSixMonthPeriod, nil
	case "annual_period":
		return BillingPeriodAnnualPeriod, nil
	}
	var t BillingPeriod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e BillingPeriod) Ptr() *BillingPeriod {
	return &e
}
