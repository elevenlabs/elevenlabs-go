package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Wrapper for anyOf/allOf composition constraints scoped to required fields.
type RequiredConstraints struct {
	AnyOf           []*RequiredConstraint `json:"any_of,omitempty" xml:"any_of,omitempty"`
	AllOf           []*RequiredConstraint `json:"all_of,omitempty" xml:"all_of,omitempty"`
	extraProperties map[string]interface{}
}

func (r *RequiredConstraints) GetAnyOf() (value []*RequiredConstraint) {
	if r == nil {
		return
	}
	return r.AnyOf
}

func (r *RequiredConstraints) GetAllOf() (value []*RequiredConstraint) {
	if r == nil {
		return
	}
	return r.AllOf
}

func (r *RequiredConstraints) SetAnyOf(anyOf []*RequiredConstraint) {
	r.AnyOf = anyOf
}

func (r *RequiredConstraints) SetAllOf(allOf []*RequiredConstraint) {
	r.AllOf = allOf
}

func (r *RequiredConstraints) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RequiredConstraints) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RequiredConstraints to string"
	}
	return string(jsonData)
}

func (r *RequiredConstraints) UnmarshalJSON(data []byte) error {
	type unmarshaler RequiredConstraints
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RequiredConstraints(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RequiredConstraints) MarshalJSON() ([]byte, error) {
	type embed RequiredConstraints
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
