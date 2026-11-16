package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ImageAvatar struct {
	// The type of the avatar
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The URL of the avatar
	Url             *string `json:"url,omitempty" xml:"url,omitempty"`
	extraProperties map[string]interface{}
}

func (i *ImageAvatar) GetType() (value *string) {
	if i == nil {
		return
	}
	return i.Type
}

func (i *ImageAvatar) GetUrl() (value *string) {
	if i == nil {
		return
	}
	return i.Url
}

func (i *ImageAvatar) SetType(type_ *string) {
	i.Type = type_
}

func (i *ImageAvatar) SetUrl(url *string) {
	i.Url = url
}

func (i *ImageAvatar) GetExtraProperties() map[string]interface{} {
	if i == nil {
		return nil
	}
	return i.extraProperties
}

func (i ImageAvatar) String() string {
	jsonData, err := json.MarshalIndent(i, "", "  ")
	if err != nil {
		return "error converting struct: ImageAvatar to string"
	}
	return string(jsonData)
}

func (i *ImageAvatar) UnmarshalJSON(data []byte) error {
	type unmarshaler ImageAvatar
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = ImageAvatar(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *i)
	if err != nil {
		return err
	}
	i.extraProperties = extraProperties
	return nil
}

func (i *ImageAvatar) MarshalJSON() ([]byte, error) {
	type embed ImageAvatar
	var marshaler = struct {
		embed
	}{
		embed: embed(*i),
	}
	return json.Marshal(marshaler)
}
