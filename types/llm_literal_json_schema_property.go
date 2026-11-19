package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LlmLiteralJsonSchemaProperty struct {
	Type        *LlmLiteralJsonSchemaPropertyType `json:"type" xml:"type" required:"true"`
	Description string                            `json:"description" xml:"description" required:"true"`
	// List of allowed string values for string type parameters
	Enum            []string `json:"enum,omitempty" xml:"enum,omitempty"`
	extraProperties map[string]interface{}
}

func (l *LlmLiteralJsonSchemaProperty) GetType() (value *LlmLiteralJsonSchemaPropertyType) {
	if l == nil {
		return
	}
	return l.Type
}

func (l *LlmLiteralJsonSchemaProperty) GetDescription() (value string) {
	if l == nil {
		return
	}
	return l.Description
}

func (l *LlmLiteralJsonSchemaProperty) GetEnum() (value []string) {
	if l == nil {
		return
	}
	return l.Enum
}

func (l *LlmLiteralJsonSchemaProperty) SetType(type_ *LlmLiteralJsonSchemaPropertyType) {
	l.Type = type_
}

func (l *LlmLiteralJsonSchemaProperty) SetDescription(description string) {
	l.Description = description
}

func (l *LlmLiteralJsonSchemaProperty) SetEnum(enum []string) {
	l.Enum = enum
}

func (l *LlmLiteralJsonSchemaProperty) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmLiteralJsonSchemaProperty) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmLiteralJsonSchemaProperty to string"
	}
	return string(jsonData)
}

func (l *LlmLiteralJsonSchemaProperty) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmLiteralJsonSchemaProperty) MarshalJSON() ([]byte, error) {
	type embed LlmLiteralJsonSchemaProperty
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
