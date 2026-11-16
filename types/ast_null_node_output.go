package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstNullNodeOutput struct {
	type_           string
	extraProperties map[string]interface{}
}

func (a *AstNullNodeOutput) Type() string {
	return a.type_
}

func (a *AstNullNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstNullNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstNullNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstNullNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstNullNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstNullNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "null_literal" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "null_literal", *literals.Type)
	}
	a.type_ = "null_literal"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstNullNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstNullNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "null_literal",
	}
	return json.Marshal(marshaler)
}
