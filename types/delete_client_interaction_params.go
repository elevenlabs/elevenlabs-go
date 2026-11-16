package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DeleteClientInteractionParams struct {
	extraProperties map[string]interface{}
}

func (d *DeleteClientInteractionParams) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteClientInteractionParams) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteClientInteractionParams to string"
	}
	return string(jsonData)
}

func (d *DeleteClientInteractionParams) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteClientInteractionParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteClientInteractionParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "smb_tool_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteClientInteractionParams) MarshalJSON() ([]byte, error) {
	type embed DeleteClientInteractionParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
