package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CreateHolidayParams struct {
	extraProperties map[string]interface{}
}

func (c *CreateHolidayParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateHolidayParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateHolidayParams to string"
	}
	return string(jsonData)
}

func (c *CreateHolidayParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateHolidayParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateHolidayParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateHolidayParams) MarshalJSON() ([]byte, error) {
	type embed CreateHolidayParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
