package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstNumberNodeInput struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Value of this literal.
	Value           *AstNumberNodeInputValue `json:"value" xml:"value" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstNumberNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstNumberNodeInput) GetValue() (value *AstNumberNodeInputValue) {
	if a == nil {
		return
	}
	return a.Value
}

func (a *AstNumberNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstNumberNodeInput) SetValue(value *AstNumberNodeInputValue) {
	a.Value = value
}

func (a *AstNumberNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstNumberNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstNumberNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstNumberNodeInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstNumberNodeInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstNumberNodeInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstNumberNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstNumberNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
