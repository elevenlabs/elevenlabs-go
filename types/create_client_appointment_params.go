package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CreateClientAppointmentParams struct {
	extraProperties map[string]interface{}
}

func (c *CreateClientAppointmentParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateClientAppointmentParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateClientAppointmentParams to string"
	}
	return string(jsonData)
}

func (c *CreateClientAppointmentParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateClientAppointmentParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateClientAppointmentParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateClientAppointmentParams) MarshalJSON() ([]byte, error) {
	type embed CreateClientAppointmentParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
