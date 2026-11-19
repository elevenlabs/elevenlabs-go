package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type UpdateAssetParams struct {
	// Whether to offer an available_hours timezone, set from the workspace's enable_per_resource_timezones.
	PerResourceTimezones *bool `json:"per_resource_timezones,omitempty" xml:"per_resource_timezones,omitempty"`
	extraProperties      map[string]interface{}
}

func (u *UpdateAssetParams) GetPerResourceTimezones() (value *bool) {
	if u == nil {
		return
	}
	return u.PerResourceTimezones
}

func (u *UpdateAssetParams) SetPerResourceTimezones(perResourceTimezones *bool) {
	u.PerResourceTimezones = perResourceTimezones
}

func (u *UpdateAssetParams) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UpdateAssetParams) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateAssetParams to string"
	}
	return string(jsonData)
}

func (u *UpdateAssetParams) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateAssetParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateAssetParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *u, "smb_tool_type")
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UpdateAssetParams) MarshalJSON() ([]byte, error) {
	type embed UpdateAssetParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
