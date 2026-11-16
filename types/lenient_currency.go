package types

import (
	"fmt"
)

type LenientCurrency string

const (
	LenientCurrencyUsd LenientCurrency = "USD"
	LenientCurrencyEur LenientCurrency = "EUR"
	LenientCurrencyGbp LenientCurrency = "GBP"
	LenientCurrencyPln LenientCurrency = "PLN"
	LenientCurrencyCad LenientCurrency = "CAD"
	LenientCurrencyAud LenientCurrency = "AUD"
	LenientCurrencyChf LenientCurrency = "CHF"
	LenientCurrencySek LenientCurrency = "SEK"
	LenientCurrencyNok LenientCurrency = "NOK"
	LenientCurrencyDkk LenientCurrency = "DKK"
	LenientCurrencyIsk LenientCurrency = "ISK"
	LenientCurrencyCzk LenientCurrency = "CZK"
	LenientCurrencyHuf LenientCurrency = "HUF"
	LenientCurrencyRon LenientCurrency = "RON"
	LenientCurrencyBgn LenientCurrency = "BGN"
	LenientCurrencyJpy LenientCurrency = "JPY"
	LenientCurrencyCny LenientCurrency = "CNY"
	LenientCurrencyTwd LenientCurrency = "TWD"
	LenientCurrencyInr LenientCurrency = "INR"
	LenientCurrencyBrl LenientCurrency = "BRL"
	LenientCurrencyMxn LenientCurrency = "MXN"
	LenientCurrencyKrw LenientCurrency = "KRW"
	LenientCurrencySgd LenientCurrency = "SGD"
	LenientCurrencyHkd LenientCurrency = "HKD"
	LenientCurrencyNzd LenientCurrency = "NZD"
	LenientCurrencyZar LenientCurrency = "ZAR"
	LenientCurrencyAed LenientCurrency = "AED"
	LenientCurrencySar LenientCurrency = "SAR"
	LenientCurrencyQar LenientCurrency = "QAR"
	LenientCurrencyKwd LenientCurrency = "KWD"
	LenientCurrencyIls LenientCurrency = "ILS"
	LenientCurrencyTry LenientCurrency = "TRY"
	LenientCurrencyThb LenientCurrency = "THB"
	LenientCurrencyIdr LenientCurrency = "IDR"
	LenientCurrencyMyr LenientCurrency = "MYR"
	LenientCurrencyPhp LenientCurrency = "PHP"
	LenientCurrencyVnd LenientCurrency = "VND"
	LenientCurrencyBdt LenientCurrency = "BDT"
	LenientCurrencyRub LenientCurrency = "RUB"
	LenientCurrencyUah LenientCurrency = "UAH"
	LenientCurrencyClp LenientCurrency = "CLP"
	LenientCurrencyCop LenientCurrency = "COP"
	LenientCurrencyArs LenientCurrency = "ARS"
	LenientCurrencyPen LenientCurrency = "PEN"
	LenientCurrencyEgp LenientCurrency = "EGP"
	LenientCurrencyNgn LenientCurrency = "NGN"
	LenientCurrencyKes LenientCurrency = "KES"
	LenientCurrencyPkr LenientCurrency = "PKR"
)

func NewLenientCurrencyFromString(s string) (LenientCurrency, error) {
	switch s {
	case "USD":
		return LenientCurrencyUsd, nil
	case "EUR":
		return LenientCurrencyEur, nil
	case "GBP":
		return LenientCurrencyGbp, nil
	case "PLN":
		return LenientCurrencyPln, nil
	case "CAD":
		return LenientCurrencyCad, nil
	case "AUD":
		return LenientCurrencyAud, nil
	case "CHF":
		return LenientCurrencyChf, nil
	case "SEK":
		return LenientCurrencySek, nil
	case "NOK":
		return LenientCurrencyNok, nil
	case "DKK":
		return LenientCurrencyDkk, nil
	case "ISK":
		return LenientCurrencyIsk, nil
	case "CZK":
		return LenientCurrencyCzk, nil
	case "HUF":
		return LenientCurrencyHuf, nil
	case "RON":
		return LenientCurrencyRon, nil
	case "BGN":
		return LenientCurrencyBgn, nil
	case "JPY":
		return LenientCurrencyJpy, nil
	case "CNY":
		return LenientCurrencyCny, nil
	case "TWD":
		return LenientCurrencyTwd, nil
	case "INR":
		return LenientCurrencyInr, nil
	case "BRL":
		return LenientCurrencyBrl, nil
	case "MXN":
		return LenientCurrencyMxn, nil
	case "KRW":
		return LenientCurrencyKrw, nil
	case "SGD":
		return LenientCurrencySgd, nil
	case "HKD":
		return LenientCurrencyHkd, nil
	case "NZD":
		return LenientCurrencyNzd, nil
	case "ZAR":
		return LenientCurrencyZar, nil
	case "AED":
		return LenientCurrencyAed, nil
	case "SAR":
		return LenientCurrencySar, nil
	case "QAR":
		return LenientCurrencyQar, nil
	case "KWD":
		return LenientCurrencyKwd, nil
	case "ILS":
		return LenientCurrencyIls, nil
	case "TRY":
		return LenientCurrencyTry, nil
	case "THB":
		return LenientCurrencyThb, nil
	case "IDR":
		return LenientCurrencyIdr, nil
	case "MYR":
		return LenientCurrencyMyr, nil
	case "PHP":
		return LenientCurrencyPhp, nil
	case "VND":
		return LenientCurrencyVnd, nil
	case "BDT":
		return LenientCurrencyBdt, nil
	case "RUB":
		return LenientCurrencyRub, nil
	case "UAH":
		return LenientCurrencyUah, nil
	case "CLP":
		return LenientCurrencyClp, nil
	case "COP":
		return LenientCurrencyCop, nil
	case "ARS":
		return LenientCurrencyArs, nil
	case "PEN":
		return LenientCurrencyPen, nil
	case "EGP":
		return LenientCurrencyEgp, nil
	case "NGN":
		return LenientCurrencyNgn, nil
	case "KES":
		return LenientCurrencyKes, nil
	case "PKR":
		return LenientCurrencyPkr, nil
	}
	var t LenientCurrency
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e LenientCurrency) Ptr() *LenientCurrency {
	return &e
}
