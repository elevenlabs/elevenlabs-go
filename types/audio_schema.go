package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioSchema struct {
	Title           *string `json:"title,omitempty" xml:"title,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AudioSchema) GetTitle() (value *string) {
	if a == nil {
		return
	}
	return a.Title
}

func (a *AudioSchema) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AudioSchema) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AudioSchema) SetTitle(title *string) {
	a.Title = title
}

func (a *AudioSchema) SetDescription(description *string) {
	a.Description = description
}

func (a *AudioSchema) SetType(type_ *string) {
	a.Type = type_
}

func (a *AudioSchema) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioSchema) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioSchema to string"
	}
	return string(jsonData)
}

func (a *AudioSchema) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioSchema) MarshalJSON() ([]byte, error) {
	type embed AudioSchema
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
