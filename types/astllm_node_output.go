package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AstllmNodeOutput struct {
	type_       string
	ValueSchema *LlmLiteralJsonSchemaProperty `json:"value_schema" xml:"value_schema" required:"true"`
	// The prompt to evaluate to a boolean value. Deprecated. Use a boolean schema instead.
	Prompt          string `json:"prompt" xml:"prompt" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AstllmNodeOutput) GetValueSchema() (value *LlmLiteralJsonSchemaProperty) {
	if a == nil {
		return
	}
	return a.ValueSchema
}

func (a *AstllmNodeOutput) GetPrompt() (value string) {
	if a == nil {
		return
	}
	return a.Prompt
}

func (a *AstllmNodeOutput) SetValueSchema(valueSchema *LlmLiteralJsonSchemaProperty) {
	a.ValueSchema = valueSchema
}

func (a *AstllmNodeOutput) SetPrompt(prompt string) {
	a.Prompt = prompt
}

func (a *AstllmNodeOutput) Type() string {
	return a.type_
}

func (a *AstllmNodeOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AstllmNodeOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AstllmNodeOutput to string"
	}
	return string(jsonData)
}

func (a *AstllmNodeOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AstllmNodeOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AstllmNodeOutput(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "llm" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "llm", *literals.Type)
	}
	a.type_ = "llm"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AstllmNodeOutput) MarshalJSON() ([]byte, error) {
	type embed AstllmNodeOutput
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "llm",
	}
	return json.Marshal(marshaler)
}
