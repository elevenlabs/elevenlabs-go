package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type VadConfig struct {
	extraProperties map[string]interface{}
}

func (v *VadConfig) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VadConfig) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VadConfig to string"
	}
	return string(jsonData)
}

func (v *VadConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler VadConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = VadConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VadConfig) MarshalJSON() ([]byte, error) {
	type embed VadConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
