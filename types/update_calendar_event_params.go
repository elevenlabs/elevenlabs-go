package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateCalendarEventParams struct {
	extraProperties map[string]interface{}
}

func (u *UpdateCalendarEventParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateCalendarEventParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateCalendarEventParams to string"
	}
	return string(jsonData)
}

func (u *UpdateCalendarEventParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateCalendarEventParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateCalendarEventParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateCalendarEventParams) MarshalJSON() ([]byte, error) {
	type embed UpdateCalendarEventParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
