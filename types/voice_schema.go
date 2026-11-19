package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceSchema struct {
	Title           *string `json:"title,omitempty" xml:"title,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"500"`
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (v *VoiceSchema) GetTitle() (value *string) {
	if v == nil {
		return
	}
	return v.Title
}

func (v *VoiceSchema) GetDescription() (value *string) {
	if v == nil {
		return
	}
	return v.Description
}

func (v *VoiceSchema) GetType() (value *string) {
	if v == nil {
		return
	}
	return v.Type
}

func (v *VoiceSchema) SetTitle(title *string) {
	v.Title = title
}

func (v *VoiceSchema) SetDescription(description *string) {
	v.Description = description
}

func (v *VoiceSchema) SetType(type_ *string) {
	v.Type = type_
}

func (v *VoiceSchema) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VoiceSchema) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceSchema to string"
	}
	return string(jsonData)
}

func (v *VoiceSchema) UnmarshalJSON(data []byte) error {
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

func (v *VoiceSchema) MarshalJSON() ([]byte, error) {
	type embed VoiceSchema
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
