package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DynamicVariablesConfig struct {
	// A dictionary of dynamic variable placeholders and their values
	DynamicVariablePlaceholders map[string]any `json:"dynamic_variable_placeholders,omitempty"`
	extraProperties             map[string]interface{}
}

func (d *DynamicVariablesConfig) GetDynamicVariablePlaceholders() (value map[string]any) {
	if d == nil {
		return
	}
	return d.DynamicVariablePlaceholders
}

func (d *DynamicVariablesConfig) SetDynamicVariablePlaceholders(dynamicVariablePlaceholders map[string]any) {
	d.DynamicVariablePlaceholders = dynamicVariablePlaceholders
}

func (d *DynamicVariablesConfig) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DynamicVariablesConfig) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DynamicVariablesConfig to string"
	}
	return string(jsonData)
}

func (d *DynamicVariablesConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler DynamicVariablesConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DynamicVariablesConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DynamicVariablesConfig) MarshalJSON() ([]byte, error) {
	type embed DynamicVariablesConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
