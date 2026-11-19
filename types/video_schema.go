package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VideoSchema struct {
	Title           *string `json:"title,omitempty" xml:"title,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VideoSchema) GetTitle() (value *string) {
	if v == nil {
		return
	}
	return v.Title
}

func (v *VideoSchema) GetDescription() (value *string) {
	if v == nil {
		return
	}
	return v.Description
}

func (v *VideoSchema) GetType() (value *string) {
	if v == nil {
		return
	}
	return v.Type
}

func (v *VideoSchema) SetTitle(title *string) {
	v.Title = title
}

func (v *VideoSchema) SetDescription(description *string) {
	v.Description = description
}

func (v *VideoSchema) SetType(type_ *string) {
	v.Type = type_
}

func (v *VideoSchema) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VideoSchema) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VideoSchema to string"
	}
	return string(jsonData)
}

func (v *VideoSchema) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, v); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VideoSchema) MarshalJSON() ([]byte, error) {
	type embed VideoSchema
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
