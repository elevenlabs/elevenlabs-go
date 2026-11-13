package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BooleanSchema struct {
	Title           *string `json:"title,omitempty" xml:"title,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (b *BooleanSchema) GetTitle() (value *string) {
	if b == nil {
		return
	}
	return b.Title
}

func (b *BooleanSchema) GetDescription() (value *string) {
	if b == nil {
		return
	}
	return b.Description
}

func (b *BooleanSchema) GetType() (value *string) {
	if b == nil {
		return
	}
	return b.Type
}

func (b *BooleanSchema) SetTitle(title *string) {
	b.Title = title
}

func (b *BooleanSchema) SetDescription(description *string) {
	b.Description = description
}

func (b *BooleanSchema) SetType(type_ *string) {
	b.Type = type_
}

func (b *BooleanSchema) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BooleanSchema) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BooleanSchema to string"
	}
	return string(jsonData)
}

func (b *BooleanSchema) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BooleanSchema) MarshalJSON() ([]byte, error) {
	type embed BooleanSchema
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
