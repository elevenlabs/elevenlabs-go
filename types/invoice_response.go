package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type InvoiceResponse struct {
	// The amount due in cents.
	AmountDueCents int `json:"amount_due_cents" xml:"amount_due_cents" required:"true"`
	// The subtotal amount in cents before tax (exclusive of tax and discounts).
	SubtotalCents *int `json:"subtotal_cents,omitempty" xml:"subtotal_cents,omitempty"`
	// The tax amount in cents.
	TaxCents *int `json:"tax_cents,omitempty" xml:"tax_cents,omitempty"`
	// Deprecated. Use [discounts] instead. The discount applied to the invoice. E.g. [20.0f] for 20% off.
	DiscountPercentOff *float64 `json:"discount_percent_off,omitempty" xml:"discount_percent_off,omitempty"`
	// Deprecated. Use [discounts] instead. The discount applied to the invoice. E.g. [20.0f] for 20 cents off.
	DiscountAmountOff *float64 `json:"discount_amount_off,omitempty" xml:"discount_amount_off,omitempty"`
	// The discounts applied to the invoice.
	Discounts []*DiscountResponseModel `json:"discounts" xml:"discounts" required:"true"`
	// The Unix timestamp of the next payment attempt. -1 when there is no next payment attempt.
	NextPaymentAttemptUnix int `json:"next_payment_attempt_unix" xml:"next_payment_attempt_unix" required:"true"`
	// Deprecated. Use [payment_intent_statusses] instead. The status of this invoice's first payment intent. None when there is no payment intent.
	PaymentIntentStatus *InvoiceResponsePaymentIntentStatus `json:"payment_intent_status" xml:"payment_intent_status" required:"true"`
	// The statuses of this invoice's payment intents. Empty list when there are no payment intents.
	PaymentIntentStatusses []InvoiceResponsePaymentIntentStatussesItem `json:"payment_intent_statusses" xml:"payment_intent_statusses" required:"true"`
	extraProperties        map[string]interface{}
}

func (i *InvoiceResponse) GetAmountDueCents() (value int) {
	if i == nil {
		return
	}
	return i.AmountDueCents
}

func (i *InvoiceResponse) GetSubtotalCents() (value *int) {
	if i == nil {
		return
	}
	return i.SubtotalCents
}

func (i *InvoiceResponse) GetTaxCents() (value *int) {
	if i == nil {
		return
	}
	return i.TaxCents
}

func (i *InvoiceResponse) GetDiscountPercentOff() (value *float64) {
	if i == nil {
		return
	}
	return i.DiscountPercentOff
}

func (i *InvoiceResponse) GetDiscountAmountOff() (value *float64) {
	if i == nil {
		return
	}
	return i.DiscountAmountOff
}

func (i *InvoiceResponse) GetDiscounts() (value []*DiscountResponseModel) {
	if i == nil {
		return
	}
	return i.Discounts
}

func (i *InvoiceResponse) GetNextPaymentAttemptUnix() (value int) {
	if i == nil {
		return
	}
	return i.NextPaymentAttemptUnix
}

func (i *InvoiceResponse) GetPaymentIntentStatus() (value *InvoiceResponsePaymentIntentStatus) {
	if i == nil {
		return
	}
	return i.PaymentIntentStatus
}

func (i *InvoiceResponse) GetPaymentIntentStatusses() (value []InvoiceResponsePaymentIntentStatussesItem) {
	if i == nil {
		return
	}
	return i.PaymentIntentStatusses
}

func (i *InvoiceResponse) SetAmountDueCents(amountDueCents int) {
	i.AmountDueCents = amountDueCents
}

func (i *InvoiceResponse) SetSubtotalCents(subtotalCents *int) {
	i.SubtotalCents = subtotalCents
}

func (i *InvoiceResponse) SetTaxCents(taxCents *int) {
	i.TaxCents = taxCents
}

func (i *InvoiceResponse) SetDiscountPercentOff(discountPercentOff *float64) {
	i.DiscountPercentOff = discountPercentOff
}

func (i *InvoiceResponse) SetDiscountAmountOff(discountAmountOff *float64) {
	i.DiscountAmountOff = discountAmountOff
}

func (i *InvoiceResponse) SetDiscounts(discounts []*DiscountResponseModel) {
	i.Discounts = discounts
}

func (i *InvoiceResponse) SetNextPaymentAttemptUnix(nextPaymentAttemptUnix int) {
	i.NextPaymentAttemptUnix = nextPaymentAttemptUnix
}

func (i *InvoiceResponse) SetPaymentIntentStatus(paymentIntentStatus *InvoiceResponsePaymentIntentStatus) {
	i.PaymentIntentStatus = paymentIntentStatus
}

func (i *InvoiceResponse) SetPaymentIntentStatusses(paymentIntentStatusses []InvoiceResponsePaymentIntentStatussesItem) {
	i.PaymentIntentStatusses = paymentIntentStatusses
}

func (i *InvoiceResponse) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i InvoiceResponse) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: InvoiceResponse to string"
	}
	return string(jsonData)
}

func (i *InvoiceResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, i); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *InvoiceResponse) MarshalJSON() ([]byte, error) {
	type embed InvoiceResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
