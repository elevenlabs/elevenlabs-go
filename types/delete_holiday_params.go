package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DeleteHolidayParams struct {
	extraProperties map[string]interface{}
}

func (d *DeleteHolidayParams) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteHolidayParams) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteHolidayParams to string"
	}
	return string(jsonData)
}

func (d *DeleteHolidayParams) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteHolidayParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteHolidayParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "smb_tool_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteHolidayParams) MarshalJSON() ([]byte, error) {
	type embed DeleteHolidayParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
