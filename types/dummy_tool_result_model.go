package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DummyToolResultModel struct {
	extraProperties map[string]interface{}
}

func (d *DummyToolResultModel) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DummyToolResultModel) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DummyToolResultModel to string"
	}
	return string(jsonData)
}

func (d *DummyToolResultModel) UnmarshalJSON(data []byte) error {
	type unmarshaler DummyToolResultModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DummyToolResultModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "result_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DummyToolResultModel) MarshalJSON() ([]byte, error) {
	type embed DummyToolResultModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
