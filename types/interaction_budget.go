package types

import (
	"fmt"
)

type InteractionBudget string

const (
	InteractionBudgetRealtime    InteractionBudget = "realtime"
	InteractionBudgetFiveMinutes InteractionBudget = "5_minutes"
	InteractionBudgetOne0Minutes InteractionBudget = "10_minutes"
	InteractionBudgetOneHour     InteractionBudget = "1_hour"
)

func NewInteractionBudgetFromString(s string) (InteractionBudget, error) {
	switch s {
	case "realtime":
		return InteractionBudgetRealtime, nil
	case "5_minutes":
		return InteractionBudgetFiveMinutes, nil
	case "10_minutes":
		return InteractionBudgetOne0Minutes, nil
	case "1_hour":
		return InteractionBudgetOneHour, nil
	}
	var t InteractionBudget
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e InteractionBudget) Ptr() *InteractionBudget {
	return &e
}
