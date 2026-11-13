package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type LlmObjectJsonSchemaPropertyInput struct {
	Type            *string                                                     `json:"type,omitempty" xml:"type,omitempty"`
	Required        []string                                                    `json:"required,omitempty" xml:"required,omitempty"`
	Description     *string                                                     `json:"description,omitempty" xml:"description,omitempty"`
	Properties      map[string]*LlmObjectJsonSchemaPropertyInputPropertiesValue `json:"properties,omitempty"`
	extraProperties map[string]interface{}
}

func (l *LlmObjectJsonSchemaPropertyInput) GetType() (value *string) {
	if l == nil {
		return
	}
	return l.Type
}

func (l *LlmObjectJsonSchemaPropertyInput) GetRequired() (value []string) {
	if l == nil {
		return
	}
	return l.Required
}

func (l *LlmObjectJsonSchemaPropertyInput) GetDescription() (value *string) {
	if l == nil {
		return
	}
	return l.Description
}

func (l *LlmObjectJsonSchemaPropertyInput) GetProperties() (value map[string]*LlmObjectJsonSchemaPropertyInputPropertiesValue) {
	if l == nil {
		return
	}
	return l.Properties
}

func (l *LlmObjectJsonSchemaPropertyInput) SetType(type_ *string) {
	l.Type = type_
}

func (l *LlmObjectJsonSchemaPropertyInput) SetRequired(required []string) {
	l.Required = required
}

func (l *LlmObjectJsonSchemaPropertyInput) SetDescription(description *string) {
	l.Description = description
}

func (l *LlmObjectJsonSchemaPropertyInput) SetProperties(properties map[string]*LlmObjectJsonSchemaPropertyInputPropertiesValue) {
	l.Properties = properties
}

func (l *LlmObjectJsonSchemaPropertyInput) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmObjectJsonSchemaPropertyInput) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmObjectJsonSchemaPropertyInput to string"
	}
	return string(jsonData)
}

func (l *LlmObjectJsonSchemaPropertyInput) UnmarshalJSON(data []byte) error {
	type unmarshaler LlmObjectJsonSchemaPropertyInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = LlmObjectJsonSchemaPropertyInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmObjectJsonSchemaPropertyInput) MarshalJSON() ([]byte, error) {
	type embed LlmObjectJsonSchemaPropertyInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
