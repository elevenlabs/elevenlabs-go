package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstBooleanNodeInput struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Value of this literal.
	Value           bool `json:"value" xml:"value" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstBooleanNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstBooleanNodeInput) GetValue() (value bool) {
	if a == nil {
		return
	}
	return a.Value
}

func (a *AstBooleanNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstBooleanNodeInput) SetValue(value bool) {
	a.Value = value
}

func (a *AstBooleanNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstBooleanNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstBooleanNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstBooleanNodeInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstBooleanNodeInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstBooleanNodeInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstBooleanNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstBooleanNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
