package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Create a new staff member in the system.
type CreateStaffParams struct {
	// Whether to offer a working_hours timezone, set from the workspace's enable_per_resource_timezones. While that is off every staff member uses the workspace timezone, so the key is not offered.
	PerResourceTimezones *bool `json:"per_resource_timezones,omitempty" xml:"per_resource_timezones,omitempty"`
	extraProperties      map[string]interface{}
}

func (c *CreateStaffParams) GetPerResourceTimezones() (value *bool) {
	if c == nil {
		return
	}
	return c.PerResourceTimezones
}

func (c *CreateStaffParams) SetPerResourceTimezones(perResourceTimezones *bool) {
	c.PerResourceTimezones = perResourceTimezones
}

func (c *CreateStaffParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateStaffParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateStaffParams to string"
	}
	return string(jsonData)
}

func (c *CreateStaffParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateStaffParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateStaffParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateStaffParams) MarshalJSON() ([]byte, error) {
	type embed CreateStaffParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
