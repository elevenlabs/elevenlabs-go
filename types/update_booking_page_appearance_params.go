package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateBookingPageAppearanceParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateBookingPageAppearanceParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateBookingPageAppearanceParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateBookingPageAppearanceParams to string"
	}
	return string(jsonData)
}

func (u *UpdateBookingPageAppearanceParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateBookingPageAppearanceParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateBookingPageAppearanceParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateBookingPageAppearanceParams) MarshalJSON() ([]byte, error) {
	type embed UpdateBookingPageAppearanceParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
