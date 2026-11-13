package types

import (
	"fmt"
)

type InvoiceResponsePaymentIntentStatussesItem string

const (
	InvoiceResponsePaymentIntentStatussesItemCanceled              InvoiceResponsePaymentIntentStatussesItem = "canceled"
	InvoiceResponsePaymentIntentStatussesItemProcessing            InvoiceResponsePaymentIntentStatussesItem = "processing"
	InvoiceResponsePaymentIntentStatussesItemRequiresAction        InvoiceResponsePaymentIntentStatussesItem = "requires_action"
	InvoiceResponsePaymentIntentStatussesItemRequiresCapture       InvoiceResponsePaymentIntentStatussesItem = "requires_capture"
	InvoiceResponsePaymentIntentStatussesItemRequiresConfirmation  InvoiceResponsePaymentIntentStatussesItem = "requires_confirmation"
	InvoiceResponsePaymentIntentStatussesItemRequiresPaymentMethod InvoiceResponsePaymentIntentStatussesItem = "requires_payment_method"
	InvoiceResponsePaymentIntentStatussesItemSucceeded             InvoiceResponsePaymentIntentStatussesItem = "succeeded"
)

func NewInvoiceResponsePaymentIntentStatussesItemFromString(s string) (InvoiceResponsePaymentIntentStatussesItem, error) {
	switch s {
	case "canceled":
		return InvoiceResponsePaymentIntentStatussesItemCanceled, nil
	case "processing":
		return InvoiceResponsePaymentIntentStatussesItemProcessing, nil
	case "requires_action":
		return InvoiceResponsePaymentIntentStatussesItemRequiresAction, nil
	case "requires_capture":
		return InvoiceResponsePaymentIntentStatussesItemRequiresCapture, nil
	case "requires_confirmation":
		return InvoiceResponsePaymentIntentStatussesItemRequiresConfirmation, nil
	case "requires_payment_method":
		return InvoiceResponsePaymentIntentStatussesItemRequiresPaymentMethod, nil
	case "succeeded":
		return InvoiceResponsePaymentIntentStatussesItemSucceeded, nil
	}
	var t InvoiceResponsePaymentIntentStatussesItem
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e InvoiceResponsePaymentIntentStatussesItem) Ptr() *InvoiceResponsePaymentIntentStatussesItem {
	return &e
}
