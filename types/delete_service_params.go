package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Delete an existing service from the system.
type DeleteServiceParams struct {
	extraProperties map[string]interface{}
}

func (d *DeleteServiceParams) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteServiceParams) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteServiceParams to string"
	}
	return string(jsonData)
}

func (d *DeleteServiceParams) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteServiceParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteServiceParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "smb_tool_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteServiceParams) MarshalJSON() ([]byte, error) {
	type embed DeleteServiceParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
