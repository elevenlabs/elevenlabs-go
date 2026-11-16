package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstBooleanNodeOutput struct {
	type_ string
	// Value of this literal.
	Value           bool `json:"value" xml:"value" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstBooleanNodeOutput) GetValue() (value bool) {
	if a == nil {
		return
	}
	return a.Value
}

func (a *AstBooleanNodeOutput) SetValue(value bool) {
	a.Value = value
}

func (a *AstBooleanNodeOutput) Type() string {
	return a.type_
}

func (a *AstBooleanNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstBooleanNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstBooleanNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstBooleanNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstBooleanNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstBooleanNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "boolean_literal" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "boolean_literal", *literals.Type)
	}
	a.type_ = "boolean_literal"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstBooleanNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstBooleanNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "boolean_literal",
	}
	return json.Marshal(marshaler)
}
