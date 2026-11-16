package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type LlmAnyOfJsonSchemaPropertyInput struct {
	AnyOf           []*LlmObjectJsonSchemaPropertyInput `json:"any_of" xml:"any_of" required:"true"`
	Description     *string                             `json:"description,omitempty" xml:"description,omitempty"`
	extraProperties map[string]interface{}
}

func (l *LlmAnyOfJsonSchemaPropertyInput) GetAnyOf() (value []*LlmObjectJsonSchemaPropertyInput) {
	if l == nil {
		return
	}
	return l.AnyOf
}

func (l *LlmAnyOfJsonSchemaPropertyInput) GetDescription() (value *string) {
	if l == nil {
		return
	}
	return l.Description
}

func (l *LlmAnyOfJsonSchemaPropertyInput) SetAnyOf(anyOf []*LlmObjectJsonSchemaPropertyInput) {
	l.AnyOf = anyOf
}

func (l *LlmAnyOfJsonSchemaPropertyInput) SetDescription(description *string) {
	l.Description = description
}

func (l *LlmAnyOfJsonSchemaPropertyInput) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmAnyOfJsonSchemaPropertyInput) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmAnyOfJsonSchemaPropertyInput to string"
	}
	return string(jsonData)
}

func (l *LlmAnyOfJsonSchemaPropertyInput) UnmarshalJSON(data []byte) error {
	type unmarshaler LlmAnyOfJsonSchemaPropertyInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LlmAnyOfJsonSchemaPropertyInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmAnyOfJsonSchemaPropertyInput) MarshalJSON() ([]byte, error) {
	type embed LlmAnyOfJsonSchemaPropertyInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
