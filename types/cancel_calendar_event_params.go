package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CancelCalendarEventParams struct {
	extraProperties map[string]interface{}
}

func (c *CancelCalendarEventParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CancelCalendarEventParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CancelCalendarEventParams to string"
	}
	return string(jsonData)
}

func (c *CancelCalendarEventParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CancelCalendarEventParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CancelCalendarEventParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CancelCalendarEventParams) MarshalJSON() ([]byte, error) {
	type embed CancelCalendarEventParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
