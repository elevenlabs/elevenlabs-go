package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstllmNodeInputZero struct {
	Type            *string                       `json:"type,omitempty" xml:"type,omitempty"`
	ValueSchema     *LlmLiteralJsonSchemaProperty `json:"value_schema" xml:"value_schema" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstllmNodeInputZero) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AstllmNodeInputZero) GetValueSchema() (value *LlmLiteralJsonSchemaProperty) {
	if a == nil {
		return
	}
	return a.ValueSchema
}

func (a *AstllmNodeInputZero) SetType(type_ *string) {
	a.Type = type_
}

func (a *AstllmNodeInputZero) SetValueSchema(valueSchema *LlmLiteralJsonSchemaProperty) {
	a.ValueSchema = valueSchema
}

func (a *AstllmNodeInputZero) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstllmNodeInputZero) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstllmNodeInputZero to string"
	}
	return string(jsonData)
}

func (a *AstllmNodeInputZero) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstllmNodeInputZero
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstllmNodeInputZero(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstllmNodeInputZero) MarshalJSON() ([]byte, error) {
	type embed AstllmNodeInputZero
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
