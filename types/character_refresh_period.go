package types

import (
	"fmt"
)

type CharacterRefreshPeriod string

const (
	CharacterRefreshPeriodMonthlyPeriod    CharacterRefreshPeriod = "monthly_period"
	CharacterRefreshPeriodThreeMonthPeriod CharacterRefreshPeriod = "3_month_period"
	CharacterRefreshPeriodSixMonthPeriod   CharacterRefreshPeriod = "6_month_period"
	CharacterRefreshPeriodAnnualPeriod     CharacterRefreshPeriod = "annual_period"
)

func NewCharacterRefreshPeriodFromString(s string) (CharacterRefreshPeriod, error) {
	switch s {
	case "monthly_period":
		return CharacterRefreshPeriodMonthlyPeriod, nil
	case "3_month_period":
		return CharacterRefreshPeriodThreeMonthPeriod, nil
	case "6_month_period":
		return CharacterRefreshPeriodSixMonthPeriod, nil
	case "annual_period":
		return CharacterRefreshPeriodAnnualPeriod, nil
	}
	var t CharacterRefreshPeriod
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e CharacterRefreshPeriod) Ptr() *CharacterRefreshPeriod {
	return &e
}
