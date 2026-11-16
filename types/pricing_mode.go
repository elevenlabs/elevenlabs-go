package types

import (
	"fmt"
)

// Whether a catalog item's price and duration are known up front.
//
// - `fixed`: the catalog price and duration apply (the default).
// - `quote`: neither is known until the owner has seen the job. The
// receptionist collects intake answers and files a quote request instead
// of booking a slot or taking an order.
type PricingMode string

const (
	PricingModeFixed PricingMode = "fixed"
	PricingModeQuote PricingMode = "quote"
)

func NewPricingModeFromString(s string) (PricingMode, error) {
	switch s {
	case "fixed":
		return PricingModeFixed, nil
	case "quote":
		return PricingModeQuote, nil
	}
	var t PricingMode
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e PricingMode) Ptr() *PricingMode {
	return &e
}
