package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"time"
)

type SubmitOrderResponse struct {
	OrderId OrderId    `json:"order_id" xml:"order_id" required:"true" pattern:"^prodorder_[a-z0-9]{26}$"`
	State   OrderState `json:"state" xml:"state" required:"true"`
	// The timestamp when the order was submitted.
	SubmittedAt     time.Time `json:"submitted_at" xml:"submitted_at" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubmitOrderResponse) GetOrderId() (value OrderId) {
	if s == nil {
		return
	}
	return s.OrderId
}

func (s *SubmitOrderResponse) GetState() (value OrderState) {
	if s == nil {
		return
	}
	return s.State
}

func (s *SubmitOrderResponse) GetSubmittedAt() (value time.Time) {
	if s == nil {
		return
	}
	return s.SubmittedAt
}

func (s *SubmitOrderResponse) SetOrderId(orderId OrderId) {
	s.OrderId = orderId
}

func (s *SubmitOrderResponse) SetState(state OrderState) {
	s.State = state
}

func (s *SubmitOrderResponse) SetSubmittedAt(submittedAt time.Time) {
	s.SubmittedAt = submittedAt
}

func (s *SubmitOrderResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubmitOrderResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubmitOrderResponse to string"
	}
	return string(jsonData)
}

func (s *SubmitOrderResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubmitOrderResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubmitOrderResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubmitOrderResponse) MarshalJSON() ([]byte, error) {
	type embed SubmitOrderResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
