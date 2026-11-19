package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ImageSchema struct {
	Title           *string `json:"title,omitempty" xml:"title,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (i *ImageSchema) GetTitle() (value *string) {
	if i == nil {
		return
	}
	return i.Title
}

func (i *ImageSchema) GetDescription() (value *string) {
	if i == nil {
		return
	}
	return i.Description
}

func (i *ImageSchema) GetType() (value *string) {
	if i == nil {
		return
	}
	return i.Type
}

func (i *ImageSchema) SetTitle(title *string) {
	i.Title = title
}

func (i *ImageSchema) SetDescription(description *string) {
	i.Description = description
}

func (i *ImageSchema) SetType(type_ *string) {
	i.Type = type_
}

func (i *ImageSchema) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i ImageSchema) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: ImageSchema to string"
	}
	return string(jsonData)
}

func (i *ImageSchema) UnmarshalJSON(data []byte) error {
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

func (i *ImageSchema) MarshalJSON() ([]byte, error) {
	type embed ImageSchema
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
