package types

import (
	"fmt"
)

type Currency string

const (
	CurrencyUsd Currency = "usd"
	CurrencyEur Currency = "eur"
	CurrencyInr Currency = "inr"
	CurrencyPln Currency = "pln"
	CurrencyGbp Currency = "gbp"
)

func NewCurrencyFromString(s string) (Currency, error) {
	switch s {
	case "usd":
		return CurrencyUsd, nil
	case "eur":
		return CurrencyEur, nil
	case "inr":
		return CurrencyInr, nil
	case "pln":
		return CurrencyPln, nil
	case "gbp":
		return CurrencyGbp, nil
	}
	var t Currency
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e Currency) Ptr() *Currency {
	return &e
}
