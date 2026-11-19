package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A single item extracted from a page — could be a service, person, product, or asset.
type RawItem struct {
	Name            string   `json:"name" xml:"name" required:"true"`
	Description     *string  `json:"description,omitempty" xml:"description,omitempty"`
	Price           *float64 `json:"price,omitempty" xml:"price,omitempty"`
	DurationMinutes *int     `json:"duration_minutes,omitempty" xml:"duration_minutes,omitempty"`
	extraProperties map[string]interface{}
}

func (r *RawItem) GetName() (value string) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *RawItem) GetDescription() (value *string) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *RawItem) GetPrice() (value *float64) {
	if r == nil {
		return
	}
	return r.Price
}

func (r *RawItem) GetDurationMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.DurationMinutes
}

func (r *RawItem) SetName(name string) {
	r.Name = name
}

func (r *RawItem) SetDescription(description *string) {
	r.Description = description
}

func (r *RawItem) SetPrice(price *float64) {
	r.Price = price
}

func (r *RawItem) SetDurationMinutes(durationMinutes *int) {
	r.DurationMinutes = durationMinutes
}

func (r *RawItem) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RawItem) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RawItem to string"
	}
	return string(jsonData)
}

func (r *RawItem) UnmarshalJSON(data []byte) error {
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

func (r *RawItem) MarshalJSON() ([]byte, error) {
	type embed RawItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
