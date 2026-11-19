package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CreateAssetParams struct {
	// Whether to offer an available_hours timezone, set from the workspace's enable_per_resource_timezones.
	PerResourceTimezones *bool `json:"per_resource_timezones,omitempty" xml:"per_resource_timezones,omitempty"`
	extraProperties      map[string]interface{}
}

func (c *CreateAssetParams) GetPerResourceTimezones() (value *bool) {
	if c == nil {
		return
	}
	return c.PerResourceTimezones
}

func (c *CreateAssetParams) SetPerResourceTimezones(perResourceTimezones *bool) {
	c.PerResourceTimezones = perResourceTimezones
}

func (c *CreateAssetParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateAssetParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateAssetParams to string"
	}
	return string(jsonData)
}

func (c *CreateAssetParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateAssetParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateAssetParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateAssetParams) MarshalJSON() ([]byte, error) {
	type embed CreateAssetParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
