package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UrlAvatar struct {
	// The type of the avatar
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The custom URL of the avatar
	CustomUrl       *string `json:"custom_url,omitempty" xml:"custom_url,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UrlAvatar) GetType() (value *string) {
	if u == nil {
		return
	}
	return u.Type
}

func (u *UrlAvatar) GetCustomUrl() (value *string) {
	if u == nil {
		return
	}
	return u.CustomUrl
}

func (u *UrlAvatar) SetType(type_ *string) {
	u.Type = type_
}

func (u *UrlAvatar) SetCustomUrl(customUrl *string) {
	u.CustomUrl = customUrl
}

func (u *UrlAvatar) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UrlAvatar) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UrlAvatar to string"
	}
	return string(jsonData)
}

func (u *UrlAvatar) UnmarshalJSON(data []byte) error {
	type unmarshaler UrlAvatar
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UrlAvatar(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UrlAvatar) MarshalJSON() ([]byte, error) {
	type embed UrlAvatar
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
