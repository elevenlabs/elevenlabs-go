package types

import (
	"fmt"
)

// `open` while the card is still being retried; `void` once retries ran out and the workspace's plan was ended.
type ResellerInvoiceStatus string

const (
	ResellerInvoiceStatusPaid ResellerInvoiceStatus = "paid"
	ResellerInvoiceStatusOpen ResellerInvoiceStatus = "open"
	ResellerInvoiceStatusVoid ResellerInvoiceStatus = "void"
)

func NewResellerInvoiceStatusFromString(s string) (ResellerInvoiceStatus, error) {
	switch s {
	case "paid":
		return ResellerInvoiceStatusPaid, nil
	case "open":
		return ResellerInvoiceStatusOpen, nil
	case "void":
		return ResellerInvoiceStatusVoid, nil
	}
	var t ResellerInvoiceStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ResellerInvoiceStatus) Ptr() *ResellerInvoiceStatus {
	return &e
}
