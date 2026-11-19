package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A standalone add-on that can be tacked onto a service or rental booking.
//
// An add-on is a small, free-form extra (e.g. "Hot towel", "Valet parking")
// that extends the parent booking's duration and total price without
// creating a separate calendar event.
type ReceptionAddOn struct {
	// Stable identifier used by booking requests and calendar snapshots. Survives name/price edits, so rebooking and reporting stay consistent.
	Id          *string      `json:"id,omitempty" xml:"id,omitempty"`
	Name        StrippedStr  `json:"name" xml:"name" required:"true"`
	Description *StrippedStr `json:"description,omitempty" xml:"description,omitempty"`
	// Minutes added to the parent booking when this add-on is selected.
	AddDurationMinutes *int `json:"add_duration_minutes,omitempty" xml:"add_duration_minutes,omitempty" min:"0" max:"43200"`
	// Amount added to the parent booking's total when this add-on is selected.
	AddPrice        *float64 `json:"add_price,omitempty" xml:"add_price,omitempty" min:"0"`
	extraProperties map[string]interface{}
}

func (r *ReceptionAddOn) GetId() (value *string) {
	if r == nil {
		return
	}
	return r.Id
}

func (r *ReceptionAddOn) GetName() (value StrippedStr) {
	if r == nil {
		return
	}
	return r.Name
}

func (r *ReceptionAddOn) GetDescription() (value *StrippedStr) {
	if r == nil {
		return
	}
	return r.Description
}

func (r *ReceptionAddOn) GetAddDurationMinutes() (value *int) {
	if r == nil {
		return
	}
	return r.AddDurationMinutes
}

func (r *ReceptionAddOn) GetAddPrice() (value *float64) {
	if r == nil {
		return
	}
	return r.AddPrice
}

func (r *ReceptionAddOn) SetId(id *string) {
	r.Id = id
}

func (r *ReceptionAddOn) SetName(name StrippedStr) {
	r.Name = name
}

func (r *ReceptionAddOn) SetDescription(description *StrippedStr) {
	r.Description = description
}

func (r *ReceptionAddOn) SetAddDurationMinutes(addDurationMinutes *int) {
	r.AddDurationMinutes = addDurationMinutes
}

func (r *ReceptionAddOn) SetAddPrice(addPrice *float64) {
	r.AddPrice = addPrice
}

func (r *ReceptionAddOn) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r ReceptionAddOn) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: ReceptionAddOn to string"
	}
	return string(jsonData)
}

func (r *ReceptionAddOn) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias ReceptionAddOn
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = ReceptionAddOn(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *ReceptionAddOn) MarshalJSON() ([]byte, error) {
	type embed ReceptionAddOn
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
