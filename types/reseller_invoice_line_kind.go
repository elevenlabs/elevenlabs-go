package types

import (
	"fmt"
)

type ResellerInvoiceLineKind string

const (
	ResellerInvoiceLineKindPlan       ResellerInvoiceLineKind = "plan"
	ResellerInvoiceLineKindTelephony  ResellerInvoiceLineKind = "telephony"
	ResellerInvoiceLineKindOverage    ResellerInvoiceLineKind = "overage"
	ResellerInvoiceLineKindProration  ResellerInvoiceLineKind = "proration"
	ResellerInvoiceLineKindAdjustment ResellerInvoiceLineKind = "adjustment"
)

func NewResellerInvoiceLineKindFromString(s string) (ResellerInvoiceLineKind, error) {
	switch s {
	case "plan":
		return ResellerInvoiceLineKindPlan, nil
	case "telephony":
		return ResellerInvoiceLineKindTelephony, nil
	case "overage":
		return ResellerInvoiceLineKindOverage, nil
	case "proration":
		return ResellerInvoiceLineKindProration, nil
	case "adjustment":
		return ResellerInvoiceLineKindAdjustment, nil
	}
	var t ResellerInvoiceLineKind
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ResellerInvoiceLineKind) Ptr() *ResellerInvoiceLineKind {
	return &e
}
