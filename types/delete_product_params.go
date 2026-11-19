package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DeleteProductParams struct {
	extraProperties map[string]interface{}
}

func (d *DeleteProductParams) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DeleteProductParams) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DeleteProductParams to string"
	}
	return string(jsonData)
}

func (d *DeleteProductParams) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteProductParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteProductParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "smb_tool_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DeleteProductParams) MarshalJSON() ([]byte, error) {
	type embed DeleteProductParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
