package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type LlmObjectJsonSchemaPropertyOutput struct {
	Type            *string                                                      `json:"type,omitempty" xml:"type,omitempty"`
	Required        []string                                                     `json:"required,omitempty" xml:"required,omitempty"`
	Description     *string                                                      `json:"description,omitempty" xml:"description,omitempty"`
	Properties      map[string]*LlmObjectJsonSchemaPropertyOutputPropertiesValue `json:"properties,omitempty"`
	extraProperties map[string]interface{}
}

func (l *LlmObjectJsonSchemaPropertyOutput) GetType() (value *string) {
	if l == nil {
		return
	}
	return l.Type
}

func (l *LlmObjectJsonSchemaPropertyOutput) GetRequired() (value []string) {
	if l == nil {
		return
	}
	return l.Required
}

func (l *LlmObjectJsonSchemaPropertyOutput) GetDescription() (value *string) {
	if l == nil {
		return
	}
	return l.Description
}

func (l *LlmObjectJsonSchemaPropertyOutput) GetProperties() (value map[string]*LlmObjectJsonSchemaPropertyOutputPropertiesValue) {
	if l == nil {
		return
	}
	return l.Properties
}

func (l *LlmObjectJsonSchemaPropertyOutput) SetType(type_ *string) {
	l.Type = type_
}

func (l *LlmObjectJsonSchemaPropertyOutput) SetRequired(required []string) {
	l.Required = required
}

func (l *LlmObjectJsonSchemaPropertyOutput) SetDescription(description *string) {
	l.Description = description
}

func (l *LlmObjectJsonSchemaPropertyOutput) SetProperties(properties map[string]*LlmObjectJsonSchemaPropertyOutputPropertiesValue) {
	l.Properties = properties
}

func (l *LlmObjectJsonSchemaPropertyOutput) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmObjectJsonSchemaPropertyOutput) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmObjectJsonSchemaPropertyOutput to string"
	}
	return string(jsonData)
}

func (l *LlmObjectJsonSchemaPropertyOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler LlmObjectJsonSchemaPropertyOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LlmObjectJsonSchemaPropertyOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmObjectJsonSchemaPropertyOutput) MarshalJSON() ([]byte, error) {
	type embed LlmObjectJsonSchemaPropertyOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
