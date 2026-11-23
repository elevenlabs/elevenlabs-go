package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UserActivity struct {
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (u *UserActivity) GetType() (value *string) {
	if u == nil {
		return
	}
	return u.Type
}

func (u *UserActivity) SetType(type_ *string) {
	u.Type = type_
}

func (u *UserActivity) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UserActivity) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserActivity to string"
	}
	return string(jsonData)
}

func (u *UserActivity) UnmarshalJSON(data []byte) error {
	type unmarshaler UserActivity
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UserActivity(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UserActivity) MarshalJSON() ([]byte, error) {
	type embed UserActivity
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
