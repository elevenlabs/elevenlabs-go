package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Delete a client, cascading deletion of all their appointments first.
type DeleteClientParams struct {
	extraProperties map[string]interface{}
}

func (d *DeleteClientParams) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteClientParams) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteClientParams to string"
	}
	return string(jsonData)
}

func (d *DeleteClientParams) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteClientParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteClientParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "smb_tool_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteClientParams) MarshalJSON() ([]byte, error) {
	type embed DeleteClientParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
