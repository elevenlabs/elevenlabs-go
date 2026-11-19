package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LlmArrayJsonSchemaPropertyInput struct {
	Type            *string                               `json:"type,omitempty" xml:"type,omitempty"`
	Description     *string                               `json:"description,omitempty" xml:"description,omitempty"`
	Items           *LlmArrayJsonSchemaPropertyInputItems `json:"items" xml:"items" required:"true"`
	extraProperties map[string]interface{}
}

func (l *LlmArrayJsonSchemaPropertyInput) GetType() (value *string) {
	if l == nil {
		return
	}
	return l.Type
}

func (l *LlmArrayJsonSchemaPropertyInput) GetDescription() (value *string) {
	if l == nil {
		return
	}
	return l.Description
}

func (l *LlmArrayJsonSchemaPropertyInput) GetItems() (value *LlmArrayJsonSchemaPropertyInputItems) {
	if l == nil {
		return
	}
	return l.Items
}

func (l *LlmArrayJsonSchemaPropertyInput) SetType(type_ *string) {
	l.Type = type_
}

func (l *LlmArrayJsonSchemaPropertyInput) SetDescription(description *string) {
	l.Description = description
}

func (l *LlmArrayJsonSchemaPropertyInput) SetItems(items *LlmArrayJsonSchemaPropertyInputItems) {
	l.Items = items
}

func (l *LlmArrayJsonSchemaPropertyInput) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmArrayJsonSchemaPropertyInput) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmArrayJsonSchemaPropertyInput to string"
	}
	return string(jsonData)
}

func (l *LlmArrayJsonSchemaPropertyInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LlmArrayJsonSchemaPropertyInput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LlmArrayJsonSchemaPropertyInput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmArrayJsonSchemaPropertyInput) MarshalJSON() ([]byte, error) {
	type embed LlmArrayJsonSchemaPropertyInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
