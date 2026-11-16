package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type IntegerSchema struct {
	Title           *string `json:"title,omitempty" xml:"title,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (i *IntegerSchema) GetTitle() (value *string) {
	if i == nil {
		return
	}
	return i.Title
}

func (i *IntegerSchema) GetDescription() (value *string) {
	if i == nil {
		return
	}
	return i.Description
}

func (i *IntegerSchema) GetType() (value *string) {
	if i == nil {
		return
	}
	return i.Type
}

func (i *IntegerSchema) SetTitle(title *string) {
	i.Title = title
}

func (i *IntegerSchema) SetDescription(description *string) {
	i.Description = description
}

func (i *IntegerSchema) SetType(type_ *string) {
	i.Type = type_
}

func (i *IntegerSchema) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i IntegerSchema) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: IntegerSchema to string"
	}
	return string(jsonData)
}

func (i *IntegerSchema) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, i); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *IntegerSchema) MarshalJSON() ([]byte, error) {
	type embed IntegerSchema
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
