package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Delete an existing staff member from the system.
type DeleteStaffParams struct {
	extraProperties map[string]interface{}
}

func (d *DeleteStaffParams) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteStaffParams) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteStaffParams to string"
	}
	return string(jsonData)
}

func (d *DeleteStaffParams) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteStaffParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteStaffParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "smb_tool_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteStaffParams) MarshalJSON() ([]byte, error) {
	type embed DeleteStaffParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
