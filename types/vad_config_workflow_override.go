package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type VadConfigWorkflowOverride struct {
	extraProperties map[string]interface{}
}

func (v *VadConfigWorkflowOverride) GetExtraProperties() map[string]interface{} {
	if v == nil {
		return nil
	}
	return v.extraProperties
}

func (v VadConfigWorkflowOverride) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VadConfigWorkflowOverride to string"
	}
	return string(jsonData)
}

func (v *VadConfigWorkflowOverride) UnmarshalJSON(data []byte) error {
	type unmarshaler VadConfigWorkflowOverride
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*v = VadConfigWorkflowOverride(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *v)
	if err != nil {
		return err
	}
	v.extraProperties = extraProperties
	return nil
}

func (v *VadConfigWorkflowOverride) MarshalJSON() ([]byte, error) {
	type embed VadConfigWorkflowOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
