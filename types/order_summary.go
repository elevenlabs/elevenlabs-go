package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"time"
)

type OrderSummary struct {
	OrderId OrderId `json:"order_id" xml:"order_id" required:"true" pattern:"^prodorder_[a-z0-9]{26}$"`
	// The display name of the order.
	Name  string     `json:"name" xml:"name" required:"true"`
	State OrderState `json:"state" xml:"state" required:"true"`
	// The total price for all items in USD. Excluded from response until quotes are available.
	TotalAmountUsd *float64 `json:"total_amount_usd,omitempty" xml:"total_amount_usd,omitempty"`
	// Whether this is a sandbox order that auto-progresses without producer intervention.
	Sandbox *bool `json:"sandbox,omitempty" xml:"sandbox,omitempty"`
	// The timestamp when the order was submitted, if applicable.
	SubmittedAt *time.Time `json:"submitted_at,omitempty" xml:"submitted_at,omitempty"`
	// The timestamp when the order was last modified, if applicable.
	UpdatedAt *time.Time `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// The reason the order was cancelled, if applicable.
	CancelReason    *string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OrderSummary) GetOrderId() (value OrderId) {
	if o == nil {
		return
	}
	return o.OrderId
}

func (o *OrderSummary) GetName() (value string) {
	if o == nil {
		return
	}
	return o.Name
}

func (o *OrderSummary) GetState() (value OrderState) {
	if o == nil {
		return
	}
	return o.State
}

func (o *OrderSummary) GetTotalAmountUsd() (value *float64) {
	if o == nil {
		return
	}
	return o.TotalAmountUsd
}

func (o *OrderSummary) GetSandbox() (value *bool) {
	if o == nil {
		return
	}
	return o.Sandbox
}

func (o *OrderSummary) GetSubmittedAt() (value *time.Time) {
	if o == nil {
		return
	}
	return o.SubmittedAt
}

func (o *OrderSummary) GetUpdatedAt() (value *time.Time) {
	if o == nil {
		return
	}
	return o.UpdatedAt
}

func (o *OrderSummary) GetCancelReason() (value *string) {
	if o == nil {
		return
	}
	return o.CancelReason
}

func (o *OrderSummary) SetOrderId(orderId OrderId) {
	o.OrderId = orderId
}

func (o *OrderSummary) SetName(name string) {
	o.Name = name
}

func (o *OrderSummary) SetState(state OrderState) {
	o.State = state
}

func (o *OrderSummary) SetTotalAmountUsd(totalAmountUsd *float64) {
	o.TotalAmountUsd = totalAmountUsd
}

func (o *OrderSummary) SetSandbox(sandbox *bool) {
	o.Sandbox = sandbox
}

func (o *OrderSummary) SetSubmittedAt(submittedAt *time.Time) {
	o.SubmittedAt = submittedAt
}

func (o *OrderSummary) SetUpdatedAt(updatedAt *time.Time) {
	o.UpdatedAt = updatedAt
}

func (o *OrderSummary) SetCancelReason(cancelReason *string) {
	o.CancelReason = cancelReason
}

func (o *OrderSummary) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OrderSummary) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OrderSummary to string"
	}
	return string(jsonData)
}

func (o *OrderSummary) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, o); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OrderSummary) MarshalJSON() ([]byte, error) {
	type embed OrderSummary
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
