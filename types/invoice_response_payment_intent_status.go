package types

import (
	"fmt"
)

type InvoiceResponsePaymentIntentStatus string

const (
	InvoiceResponsePaymentIntentStatusCanceled              InvoiceResponsePaymentIntentStatus = "canceled"
	InvoiceResponsePaymentIntentStatusProcessing            InvoiceResponsePaymentIntentStatus = "processing"
	InvoiceResponsePaymentIntentStatusRequiresAction        InvoiceResponsePaymentIntentStatus = "requires_action"
	InvoiceResponsePaymentIntentStatusRequiresCapture       InvoiceResponsePaymentIntentStatus = "requires_capture"
	InvoiceResponsePaymentIntentStatusRequiresConfirmation  InvoiceResponsePaymentIntentStatus = "requires_confirmation"
	InvoiceResponsePaymentIntentStatusRequiresPaymentMethod InvoiceResponsePaymentIntentStatus = "requires_payment_method"
	InvoiceResponsePaymentIntentStatusSucceeded             InvoiceResponsePaymentIntentStatus = "succeeded"
)

func NewInvoiceResponsePaymentIntentStatusFromString(s string) (InvoiceResponsePaymentIntentStatus, error) {
	switch s {
	case "canceled":
		return InvoiceResponsePaymentIntentStatusCanceled, nil
	case "processing":
		return InvoiceResponsePaymentIntentStatusProcessing, nil
	case "requires_action":
		return InvoiceResponsePaymentIntentStatusRequiresAction, nil
	case "requires_capture":
		return InvoiceResponsePaymentIntentStatusRequiresCapture, nil
	case "requires_confirmation":
		return InvoiceResponsePaymentIntentStatusRequiresConfirmation, nil
	case "requires_payment_method":
		return InvoiceResponsePaymentIntentStatusRequiresPaymentMethod, nil
	case "succeeded":
		return InvoiceResponsePaymentIntentStatusSucceeded, nil
	}
	var t InvoiceResponsePaymentIntentStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e InvoiceResponsePaymentIntentStatus) Ptr() *InvoiceResponsePaymentIntentStatus {
	return &e
}
