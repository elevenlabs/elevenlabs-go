package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// A set of fields that must all be present to satisfy this constraint.
type RequiredConstraint struct {
	Required        []string `json:"required" xml:"required" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RequiredConstraint) GetRequired() (value []string) {
	if r == nil {
		return
	}
	return r.Required
}

func (r *RequiredConstraint) SetRequired(required []string) {
	r.Required = required
}

func (r *RequiredConstraint) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RequiredConstraint) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RequiredConstraint to string"
	}
	return string(jsonData)
}

func (r *RequiredConstraint) UnmarshalJSON(data []byte) error {
	type unmarshaler RequiredConstraint
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RequiredConstraint(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RequiredConstraint) MarshalJSON() ([]byte, error) {
	type embed RequiredConstraint
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
