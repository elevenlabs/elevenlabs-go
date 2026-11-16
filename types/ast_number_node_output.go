package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstNumberNodeOutput struct {
	type_ string
	// Value of this literal.
	Value           *AstNumberNodeOutputValue `json:"value" xml:"value" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstNumberNodeOutput) GetValue() (value *AstNumberNodeOutputValue) {
	if a == nil {
		return
	}
	return a.Value
}

func (a *AstNumberNodeOutput) SetValue(value *AstNumberNodeOutputValue) {
	a.Value = value
}

func (a *AstNumberNodeOutput) Type() string {
	return a.type_
}

func (a *AstNumberNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstNumberNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstNumberNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstNumberNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstNumberNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstNumberNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "number_literal" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "number_literal", *literals.Type)
	}
	a.type_ = "number_literal"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstNumberNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstNumberNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "number_literal",
	}
	return json.Marshal(marshaler)
}
