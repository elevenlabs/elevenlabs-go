package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type RestoreCalendarEventParams struct {
	extraProperties map[string]interface{}
}

func (r *RestoreCalendarEventParams) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RestoreCalendarEventParams) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RestoreCalendarEventParams to string"
	}
	return string(jsonData)
}

func (r *RestoreCalendarEventParams) UnmarshalJSON(data []byte) error {
	type unmarshaler RestoreCalendarEventParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RestoreCalendarEventParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *r, "smb_tool_type")
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RestoreCalendarEventParams) MarshalJSON() ([]byte, error) {
	type embed RestoreCalendarEventParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
