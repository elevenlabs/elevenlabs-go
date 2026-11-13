package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type LlmArrayJsonSchemaPropertyOutput struct {
	Type            *string                                `json:"type,omitempty" xml:"type,omitempty"`
	Description     *string                                `json:"description,omitempty" xml:"description,omitempty"`
	Items           *LlmArrayJsonSchemaPropertyOutputItems `json:"items" xml:"items" required:"true"`
	extraProperties map[string]interface{}
}

func (l *LlmArrayJsonSchemaPropertyOutput) GetType() (value *string) {
	if l == nil {
		return
	}
	return l.Type
}

func (l *LlmArrayJsonSchemaPropertyOutput) GetDescription() (value *string) {
	if l == nil {
		return
	}
	return l.Description
}

func (l *LlmArrayJsonSchemaPropertyOutput) GetItems() (value *LlmArrayJsonSchemaPropertyOutputItems) {
	if l == nil {
		return
	}
	return l.Items
}

func (l *LlmArrayJsonSchemaPropertyOutput) SetType(type_ *string) {
	l.Type = type_
}

func (l *LlmArrayJsonSchemaPropertyOutput) SetDescription(description *string) {
	l.Description = description
}

func (l *LlmArrayJsonSchemaPropertyOutput) SetItems(items *LlmArrayJsonSchemaPropertyOutputItems) {
	l.Items = items
}

func (l *LlmArrayJsonSchemaPropertyOutput) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l LlmArrayJsonSchemaPropertyOutput) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: LlmArrayJsonSchemaPropertyOutput to string"
	}
	return string(jsonData)
}

func (l *LlmArrayJsonSchemaPropertyOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, l); err != nil {
		return err
	}
	type alias LlmArrayJsonSchemaPropertyOutput
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*l = LlmArrayJsonSchemaPropertyOutput(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *LlmArrayJsonSchemaPropertyOutput) MarshalJSON() ([]byte, error) {
	type embed LlmArrayJsonSchemaPropertyOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
