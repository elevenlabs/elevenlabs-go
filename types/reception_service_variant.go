package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A variant of a service identified by its duration, with optional price override.
type ReceptionServiceVariant struct {
	// Stable ID of this option; keep it when editing so bookings and promotions still match.
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// How long this option takes, in minutes.
	DurationMinutes int `json:"duration_minutes" xml:"duration_minutes" required:"true" max:"43200"`
	// Price of this option in the workspace's currency; null means no listed price.
	Price *float64 `json:"price,omitempty" xml:"price,omitempty" min:"0"`
	// Whether customers can book this option.
	IsActive        *bool `json:"is_active,omitempty" xml:"is_active,omitempty"`
	extraProperties map[string]interface{}
}

func (r *ReceptionServiceVariant) GetId() (value *string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionServiceVariant) GetDurationMinutes() (value int) {
	if r == nil {
		return
	}
	return r.DurationMinutes
}

func (r *ReceptionServiceVariant) GetPrice() (value *float64) {
	if r == nil {
		return
	}
	return r.Price
}

func (r *ReceptionServiceVariant) GetIsActive() (value *bool) {
	if r == nil {
		return
	}
	return r.IsActive
}

func (r *ReceptionServiceVariant) SetId(id *string) {
	r.Id = id
}

func (r *ReceptionServiceVariant) SetDurationMinutes(durationMinutes int) {
	r.DurationMinutes = durationMinutes
}

func (r *ReceptionServiceVariant) SetPrice(price *float64) {
	r.Price = price
}

func (r *ReceptionServiceVariant) SetIsActive(isActive *bool) {
	r.IsActive = isActive
}

func (r *ReceptionServiceVariant) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionServiceVariant) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionServiceVariant to string"
	}
	return string(jsonData)
}

func (r *ReceptionServiceVariant) UnmarshalJSON(data []byte) error {
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

func (r *ReceptionServiceVariant) MarshalJSON() ([]byte, error) {
	type embed ReceptionServiceVariant
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
