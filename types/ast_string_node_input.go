package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstStringNodeInput struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Value of this literal.
	Value           string `json:"value" xml:"value" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstStringNodeInput) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstStringNodeInput) GetValue() (value string) {
	if a == nil {
		return
	}
	return a.Value
}

func (a *AstStringNodeInput) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstStringNodeInput) SetValue(value string) {
	a.Value = value
}

func (a *AstStringNodeInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstStringNodeInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstStringNodeInput to string"
	}
	return string(jsonData)
}

func (a *AstStringNodeInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstStringNodeInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstStringNodeInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstStringNodeInput) MarshalJSON() ([]byte, error) {
	type embed AstStringNodeInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
