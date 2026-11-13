package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Create a new service (classic, rental, or group) in the system.
type CreateServiceParams struct {
	extraProperties map[string]interface{}
}

func (c *CreateServiceParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateServiceParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateServiceParams to string"
	}
	return string(jsonData)
}

func (c *CreateServiceParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateServiceParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateServiceParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateServiceParams) MarshalJSON() ([]byte, error) {
	type embed CreateServiceParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
