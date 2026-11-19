package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// One line of the workspace's invoice; an `adjustment` line covers any remainder, so lines sum to the charge.
type ResellerInvoiceLine struct {
	Kind        ResellerInvoiceLineKind `json:"kind" xml:"kind" required:"true"`
	Description string                  `json:"description" xml:"description" required:"true"`
	Quantity    *int                    `json:"quantity,omitempty" xml:"quantity,omitempty"`
	// Negative for a proration credit.
	AmountCents     int `json:"amount_cents" xml:"amount_cents" required:"true"`
	PeriodStartUnix int `json:"period_start_unix" xml:"period_start_unix" required:"true"`
	PeriodEndUnix   int `json:"period_end_unix" xml:"period_end_unix" required:"true"`
	extraProperties map[string]interface{}
}

func (r *ResellerInvoiceLine) GetKind() (value ResellerInvoiceLineKind) {
	if r == nil {
		return
	}
	return r.Kind
}

func (r *ResellerInvoiceLine) GetDescription() (value string) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ResellerInvoiceLine) GetQuantity() (value *int) {
	if r == nil {
		return
	}
	return r.Quantity
}

func (r *ResellerInvoiceLine) GetAmountCents() (value int) {
	if r == nil {
		return
	}
	return r.AmountCents
}

func (r *ResellerInvoiceLine) GetPeriodStartUnix() (value int) {
	if r == nil {
		return
	}
	return r.PeriodStartUnix
}

func (r *ResellerInvoiceLine) GetPeriodEndUnix() (value int) {
	if r == nil {
		return
	}
	return r.PeriodEndUnix
}

func (r *ResellerInvoiceLine) SetKind(kind ResellerInvoiceLineKind) {
	r.Kind = kind
}

func (r *ResellerInvoiceLine) SetDescription(description string) {
	r.Description = description
}

func (r *ResellerInvoiceLine) SetQuantity(quantity *int) {
	r.Quantity = quantity
}

func (r *ResellerInvoiceLine) SetAmountCents(amountCents int) {
	r.AmountCents = amountCents
}

func (r *ResellerInvoiceLine) SetPeriodStartUnix(periodStartUnix int) {
	r.PeriodStartUnix = periodStartUnix
}

func (r *ResellerInvoiceLine) SetPeriodEndUnix(periodEndUnix int) {
	r.PeriodEndUnix = periodEndUnix
}

func (r *ResellerInvoiceLine) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ResellerInvoiceLine) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ResellerInvoiceLine to string"
	}
	return string(jsonData)
}

func (r *ResellerInvoiceLine) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, r); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ResellerInvoiceLine) MarshalJSON() ([]byte, error) {
	type embed ResellerInvoiceLine
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
