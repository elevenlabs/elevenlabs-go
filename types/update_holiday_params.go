package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateHolidayParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateHolidayParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateHolidayParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateHolidayParams to string"
	}
	return string(jsonData)
}

func (u *UpdateHolidayParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateHolidayParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateHolidayParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateHolidayParams) MarshalJSON() ([]byte, error) {
	type embed UpdateHolidayParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
