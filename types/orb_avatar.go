package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type OrbAvatar struct {
	// The type of the avatar
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The first color of the avatar
	Color1 *string `json:"color_1,omitempty" xml:"color_1,omitempty"`
	// The second color of the avatar
	Color2          *string `json:"color_2,omitempty" xml:"color_2,omitempty"`
	extraProperties map[string]interface{}
}

func (o *OrbAvatar) GetType() (value *string) {
	if o == nil {
		return
	}
	return o.Type
}

func (o *OrbAvatar) GetColor1() (value *string) {
	if o == nil {
		return
	}
	return o.Color1
}

func (o *OrbAvatar) GetColor2() (value *string) {
	if o == nil {
		return
	}
	return o.Color2
}

func (o *OrbAvatar) SetType(type_ *string) {
	o.Type = type_
}

func (o *OrbAvatar) SetColor1(color1 *string) {
	o.Color1 = color1
}

func (o *OrbAvatar) SetColor2(color2 *string) {
	o.Color2 = color2
}

func (o *OrbAvatar) GetExtraProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.extraProperties
}

func (o OrbAvatar) String() string {
	jsonData, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return "error converting struct: OrbAvatar to string"
	}
	return string(jsonData)
}

func (o *OrbAvatar) UnmarshalJSON(data []byte) error {
	type unmarshaler OrbAvatar
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = OrbAvatar(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *o)
	if err != nil {
		return err
	}
	o.extraProperties = extraProperties
	return nil
}

func (o *OrbAvatar) MarshalJSON() ([]byte, error) {
	type embed OrbAvatar
	var marshaler = struct {
		embed
	}{
		embed: embed(*o),
	}
	return json.Marshal(marshaler)
}
