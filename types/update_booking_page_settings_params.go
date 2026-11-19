package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateBookingPageSettingsParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateBookingPageSettingsParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateBookingPageSettingsParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateBookingPageSettingsParams to string"
	}
	return string(jsonData)
}

func (u *UpdateBookingPageSettingsParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateBookingPageSettingsParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateBookingPageSettingsParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateBookingPageSettingsParams) MarshalJSON() ([]byte, error) {
	type embed UpdateBookingPageSettingsParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
