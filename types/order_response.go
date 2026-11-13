package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"time"
)

type OrderResponse struct {
	OrderId OrderId `json:"order_id" xml:"order_id" required:"true" pattern:"^prodorder_[a-z0-9]{26}$"`
	// The display name of the order.
	Name  string     `json:"name" xml:"name" required:"true"`
	State OrderState `json:"state" xml:"state" required:"true"`
	// The list of items in this order with their quotes.
	Items []*OrderItemInfo `json:"items" xml:"items" required:"true"`
	// The total price for all items in USD. Excluded from response until quotes are available.
	TotalAmountUsd *float64 `json:"total_amount_usd,omitempty" xml:"total_amount_usd,omitempty"`
	// Whether this is a sandbox order that auto-progresses without producer intervention.
	Sandbox *bool `json:"sandbox,omitempty" xml:"sandbox,omitempty"`
	// The timestamp when the order was created.
	CreatedAt time.Time `json:"created_at" xml:"created_at" required:"true"`
	// The timestamp when the order was submitted, if applicable.
	SubmittedAt *time.Time `json:"submitted_at,omitempty" xml:"submitted_at,omitempty"`
	// The timestamp when payment was received, if applicable.
	PaidAt *time.Time `json:"paid_at,omitempty" xml:"paid_at,omitempty"`
	// The timestamp when the order was accepted for production, if applicable.
	AcceptedAt *time.Time `json:"accepted_at,omitempty" xml:"accepted_at,omitempty"`
	// The timestamp when the order was completed, if applicable.
	CompletedAt *time.Time `json:"completed_at,omitempty" xml:"completed_at,omitempty"`
	// The reason the order was cancelled, if applicable.
	CancelReason    *string `json:"cancel_reason,omitempty" xml:"cancel_reason,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OrderResponse) GetOrderId() (value OrderId) {
	if o == nil {
		return
	}
	return o.OrderId
}

func (o *OrderResponse) GetName() (value string) {
	if o == nil {
		return
	}
	return o.Name
}

func (o *OrderResponse) GetState() (value OrderState) {
	if o == nil {
		return
	}
	return o.State
}

func (o *OrderResponse) GetItems() (value []*OrderItemInfo) {
	if o == nil {
		return
	}
	return o.Items
}

func (o *OrderResponse) GetTotalAmountUsd() (value *float64) {
	if o == nil {
		return
	}
	return o.TotalAmountUsd
}

func (o *OrderResponse) GetSandbox() (value *bool) {
	if o == nil {
		return
	}
	return o.Sandbox
}

func (o *OrderResponse) GetCreatedAt() (value time.Time) {
	if o == nil {
		return
	}
	return o.CreatedAt
}

func (o *OrderResponse) GetSubmittedAt() (value *time.Time) {
	if o == nil {
		return
	}
	return o.SubmittedAt
}

func (o *OrderResponse) GetPaidAt() (value *time.Time) {
	if o == nil {
		return
	}
	return o.PaidAt
}

func (o *OrderResponse) GetAcceptedAt() (value *time.Time) {
	if o == nil {
		return
	}
	return o.AcceptedAt
}

func (o *OrderResponse) GetCompletedAt() (value *time.Time) {
	if o == nil {
		return
	}
	return o.CompletedAt
}

func (o *OrderResponse) GetCancelReason() (value *string) {
	if o == nil {
		return
	}
	return o.CancelReason
}

func (o *OrderResponse) SetOrderId(orderId OrderId) {
	o.OrderId = orderId
}

func (o *OrderResponse) SetName(name string) {
	o.Name = name
}

func (o *OrderResponse) SetState(state OrderState) {
	o.State = state
}

func (o *OrderResponse) SetItems(items []*OrderItemInfo) {
	o.Items = items
}

func (o *OrderResponse) SetTotalAmountUsd(totalAmountUsd *float64) {
	o.TotalAmountUsd = totalAmountUsd
}

func (o *OrderResponse) SetSandbox(sandbox *bool) {
	o.Sandbox = sandbox
}

func (o *OrderResponse) SetCreatedAt(createdAt time.Time) {
	o.CreatedAt = createdAt
}

func (o *OrderResponse) SetSubmittedAt(submittedAt *time.Time) {
	o.SubmittedAt = submittedAt
}

func (o *OrderResponse) SetPaidAt(paidAt *time.Time) {
	o.PaidAt = paidAt
}

func (o *OrderResponse) SetAcceptedAt(acceptedAt *time.Time) {
	o.AcceptedAt = acceptedAt
}

func (o *OrderResponse) SetCompletedAt(completedAt *time.Time) {
	o.CompletedAt = completedAt
}

func (o *OrderResponse) SetCancelReason(cancelReason *string) {
	o.CancelReason = cancelReason
}

func (o *OrderResponse) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OrderResponse) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OrderResponse to string"
	}
	return string(jsonData)
}

func (o *OrderResponse) UnmarshalJSON(data []byte) error {
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

func (o *OrderResponse) MarshalJSON() ([]byte, error) {
	type embed OrderResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
