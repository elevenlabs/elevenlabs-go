package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Update an existing staff member's information.
type UpdateStaffParams struct {
	// Whether to offer a working_hours timezone, set from the workspace's enable_per_resource_timezones. While that is off every staff member uses the workspace timezone, so the key is not offered.
	PerResourceTimezones *bool `json:"per_resource_timezones,omitempty" xml:"per_resource_timezones,omitempty"`
	extraProperties      map[string]interface{}
}

func (u *UpdateStaffParams) GetPerResourceTimezones() (value *bool) {
	if u == nil {
		return
	}
	return u.PerResourceTimezones
}

func (u *UpdateStaffParams) SetPerResourceTimezones(perResourceTimezones *bool) {
	u.PerResourceTimezones = perResourceTimezones
}

func (u *UpdateStaffParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateStaffParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateStaffParams to string"
	}
	return string(jsonData)
}

func (u *UpdateStaffParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateStaffParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateStaffParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateStaffParams) MarshalJSON() ([]byte, error) {
	type embed UpdateStaffParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
