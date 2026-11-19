package types

import (
	"fmt"
)

// What a promotion discounts: the service price or add-on prices.
type PromotionTarget string

const (
	PromotionTargetService PromotionTarget = "service"
	PromotionTargetAddOn   PromotionTarget = "add_on"
)

func NewPromotionTargetFromString(s string) (PromotionTarget, error) {
	switch s {
	case "service":
		return PromotionTargetService, nil
	case "add_on":
		return PromotionTargetAddOn, nil
	}
	var t PromotionTarget
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e PromotionTarget) Ptr() *PromotionTarget {
	return &e
}
