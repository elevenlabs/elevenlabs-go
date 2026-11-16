package types

import (
	"fmt"
)

// Currency code for all prices (ISO 4217)
type ReceptionAgentInfoResponseCurrency string

const (
	ReceptionAgentInfoResponseCurrencyUsd ReceptionAgentInfoResponseCurrency = "USD"
	ReceptionAgentInfoResponseCurrencyEur ReceptionAgentInfoResponseCurrency = "EUR"
	ReceptionAgentInfoResponseCurrencyGbp ReceptionAgentInfoResponseCurrency = "GBP"
	ReceptionAgentInfoResponseCurrencyPln ReceptionAgentInfoResponseCurrency = "PLN"
	ReceptionAgentInfoResponseCurrencyCad ReceptionAgentInfoResponseCurrency = "CAD"
	ReceptionAgentInfoResponseCurrencyAud ReceptionAgentInfoResponseCurrency = "AUD"
	ReceptionAgentInfoResponseCurrencyChf ReceptionAgentInfoResponseCurrency = "CHF"
	ReceptionAgentInfoResponseCurrencySek ReceptionAgentInfoResponseCurrency = "SEK"
	ReceptionAgentInfoResponseCurrencyNok ReceptionAgentInfoResponseCurrency = "NOK"
	ReceptionAgentInfoResponseCurrencyDkk ReceptionAgentInfoResponseCurrency = "DKK"
	ReceptionAgentInfoResponseCurrencyIsk ReceptionAgentInfoResponseCurrency = "ISK"
	ReceptionAgentInfoResponseCurrencyCzk ReceptionAgentInfoResponseCurrency = "CZK"
	ReceptionAgentInfoResponseCurrencyHuf ReceptionAgentInfoResponseCurrency = "HUF"
	ReceptionAgentInfoResponseCurrencyRon ReceptionAgentInfoResponseCurrency = "RON"
	ReceptionAgentInfoResponseCurrencyBgn ReceptionAgentInfoResponseCurrency = "BGN"
	ReceptionAgentInfoResponseCurrencyJpy ReceptionAgentInfoResponseCurrency = "JPY"
	ReceptionAgentInfoResponseCurrencyCny ReceptionAgentInfoResponseCurrency = "CNY"
	ReceptionAgentInfoResponseCurrencyTwd ReceptionAgentInfoResponseCurrency = "TWD"
	ReceptionAgentInfoResponseCurrencyInr ReceptionAgentInfoResponseCurrency = "INR"
	ReceptionAgentInfoResponseCurrencyBrl ReceptionAgentInfoResponseCurrency = "BRL"
	ReceptionAgentInfoResponseCurrencyMxn ReceptionAgentInfoResponseCurrency = "MXN"
	ReceptionAgentInfoResponseCurrencyKrw ReceptionAgentInfoResponseCurrency = "KRW"
	ReceptionAgentInfoResponseCurrencySgd ReceptionAgentInfoResponseCurrency = "SGD"
	ReceptionAgentInfoResponseCurrencyHkd ReceptionAgentInfoResponseCurrency = "HKD"
	ReceptionAgentInfoResponseCurrencyNzd ReceptionAgentInfoResponseCurrency = "NZD"
	ReceptionAgentInfoResponseCurrencyZar ReceptionAgentInfoResponseCurrency = "ZAR"
	ReceptionAgentInfoResponseCurrencyAed ReceptionAgentInfoResponseCurrency = "AED"
	ReceptionAgentInfoResponseCurrencySar ReceptionAgentInfoResponseCurrency = "SAR"
	ReceptionAgentInfoResponseCurrencyQar ReceptionAgentInfoResponseCurrency = "QAR"
	ReceptionAgentInfoResponseCurrencyKwd ReceptionAgentInfoResponseCurrency = "KWD"
	ReceptionAgentInfoResponseCurrencyIls ReceptionAgentInfoResponseCurrency = "ILS"
	ReceptionAgentInfoResponseCurrencyTry ReceptionAgentInfoResponseCurrency = "TRY"
	ReceptionAgentInfoResponseCurrencyThb ReceptionAgentInfoResponseCurrency = "THB"
	ReceptionAgentInfoResponseCurrencyIdr ReceptionAgentInfoResponseCurrency = "IDR"
	ReceptionAgentInfoResponseCurrencyMyr ReceptionAgentInfoResponseCurrency = "MYR"
	ReceptionAgentInfoResponseCurrencyPhp ReceptionAgentInfoResponseCurrency = "PHP"
	ReceptionAgentInfoResponseCurrencyVnd ReceptionAgentInfoResponseCurrency = "VND"
	ReceptionAgentInfoResponseCurrencyBdt ReceptionAgentInfoResponseCurrency = "BDT"
	ReceptionAgentInfoResponseCurrencyRub ReceptionAgentInfoResponseCurrency = "RUB"
	ReceptionAgentInfoResponseCurrencyUah ReceptionAgentInfoResponseCurrency = "UAH"
	ReceptionAgentInfoResponseCurrencyClp ReceptionAgentInfoResponseCurrency = "CLP"
	ReceptionAgentInfoResponseCurrencyCop ReceptionAgentInfoResponseCurrency = "COP"
	ReceptionAgentInfoResponseCurrencyArs ReceptionAgentInfoResponseCurrency = "ARS"
	ReceptionAgentInfoResponseCurrencyPen ReceptionAgentInfoResponseCurrency = "PEN"
	ReceptionAgentInfoResponseCurrencyEgp ReceptionAgentInfoResponseCurrency = "EGP"
	ReceptionAgentInfoResponseCurrencyNgn ReceptionAgentInfoResponseCurrency = "NGN"
	ReceptionAgentInfoResponseCurrencyKes ReceptionAgentInfoResponseCurrency = "KES"
	ReceptionAgentInfoResponseCurrencyPkr ReceptionAgentInfoResponseCurrency = "PKR"
)

func NewReceptionAgentInfoResponseCurrencyFromString(s string) (ReceptionAgentInfoResponseCurrency, error) {
	switch s {
	case "USD":
		return ReceptionAgentInfoResponseCurrencyUsd, nil
	case "EUR":
		return ReceptionAgentInfoResponseCurrencyEur, nil
	case "GBP":
		return ReceptionAgentInfoResponseCurrencyGbp, nil
	case "PLN":
		return ReceptionAgentInfoResponseCurrencyPln, nil
	case "CAD":
		return ReceptionAgentInfoResponseCurrencyCad, nil
	case "AUD":
		return ReceptionAgentInfoResponseCurrencyAud, nil
	case "CHF":
		return ReceptionAgentInfoResponseCurrencyChf, nil
	case "SEK":
		return ReceptionAgentInfoResponseCurrencySek, nil
	case "NOK":
		return ReceptionAgentInfoResponseCurrencyNok, nil
	case "DKK":
		return ReceptionAgentInfoResponseCurrencyDkk, nil
	case "ISK":
		return ReceptionAgentInfoResponseCurrencyIsk, nil
	case "CZK":
		return ReceptionAgentInfoResponseCurrencyCzk, nil
	case "HUF":
		return ReceptionAgentInfoResponseCurrencyHuf, nil
	case "RON":
		return ReceptionAgentInfoResponseCurrencyRon, nil
	case "BGN":
		return ReceptionAgentInfoResponseCurrencyBgn, nil
	case "JPY":
		return ReceptionAgentInfoResponseCurrencyJpy, nil
	case "CNY":
		return ReceptionAgentInfoResponseCurrencyCny, nil
	case "TWD":
		return ReceptionAgentInfoResponseCurrencyTwd, nil
	case "INR":
		return ReceptionAgentInfoResponseCurrencyInr, nil
	case "BRL":
		return ReceptionAgentInfoResponseCurrencyBrl, nil
	case "MXN":
		return ReceptionAgentInfoResponseCurrencyMxn, nil
	case "KRW":
		return ReceptionAgentInfoResponseCurrencyKrw, nil
	case "SGD":
		return ReceptionAgentInfoResponseCurrencySgd, nil
	case "HKD":
		return ReceptionAgentInfoResponseCurrencyHkd, nil
	case "NZD":
		return ReceptionAgentInfoResponseCurrencyNzd, nil
	case "ZAR":
		return ReceptionAgentInfoResponseCurrencyZar, nil
	case "AED":
		return ReceptionAgentInfoResponseCurrencyAed, nil
	case "SAR":
		return ReceptionAgentInfoResponseCurrencySar, nil
	case "QAR":
		return ReceptionAgentInfoResponseCurrencyQar, nil
	case "KWD":
		return ReceptionAgentInfoResponseCurrencyKwd, nil
	case "ILS":
		return ReceptionAgentInfoResponseCurrencyIls, nil
	case "TRY":
		return ReceptionAgentInfoResponseCurrencyTry, nil
	case "THB":
		return ReceptionAgentInfoResponseCurrencyThb, nil
	case "IDR":
		return ReceptionAgentInfoResponseCurrencyIdr, nil
	case "MYR":
		return ReceptionAgentInfoResponseCurrencyMyr, nil
	case "PHP":
		return ReceptionAgentInfoResponseCurrencyPhp, nil
	case "VND":
		return ReceptionAgentInfoResponseCurrencyVnd, nil
	case "BDT":
		return ReceptionAgentInfoResponseCurrencyBdt, nil
	case "RUB":
		return ReceptionAgentInfoResponseCurrencyRub, nil
	case "UAH":
		return ReceptionAgentInfoResponseCurrencyUah, nil
	case "CLP":
		return ReceptionAgentInfoResponseCurrencyClp, nil
	case "COP":
		return ReceptionAgentInfoResponseCurrencyCop, nil
	case "ARS":
		return ReceptionAgentInfoResponseCurrencyArs, nil
	case "PEN":
		return ReceptionAgentInfoResponseCurrencyPen, nil
	case "EGP":
		return ReceptionAgentInfoResponseCurrencyEgp, nil
	case "NGN":
		return ReceptionAgentInfoResponseCurrencyNgn, nil
	case "KES":
		return ReceptionAgentInfoResponseCurrencyKes, nil
	case "PKR":
		return ReceptionAgentInfoResponseCurrencyPkr, nil
	}
	var t ReceptionAgentInfoResponseCurrency
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ReceptionAgentInfoResponseCurrency) Ptr() *ReceptionAgentInfoResponseCurrency {
	return &e
}
