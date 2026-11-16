package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstStringNodeOutput struct {
	type_ string
	// Value of this literal.
	Value           string `json:"value" xml:"value" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstStringNodeOutput) GetValue() (value string) {
	if a == nil {
		return
	}
	return a.Value
}

func (a *AstStringNodeOutput) SetValue(value string) {
	a.Value = value
}

func (a *AstStringNodeOutput) Type() string {
	return a.type_
}

func (a *AstStringNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstStringNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstStringNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstStringNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstStringNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstStringNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "string_literal" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "string_literal", *literals.Type)
	}
	a.type_ = "string_literal"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstStringNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstStringNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "string_literal",
	}
	return json.Marshal(marshaler)
}
