package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DynamicVariablesConfigWorkflowOverride struct {
	// A dictionary of dynamic variable placeholders and their values
	DynamicVariablePlaceholders map[string]any `json:"dynamic_variable_placeholders,omitempty"`
	extraProperties             map[string]interface{}
}

func (d *DynamicVariablesConfigWorkflowOverride) GetDynamicVariablePlaceholders() (value map[string]any) {
	if d == nil {
		return
	}
	return d.DynamicVariablePlaceholders
}

func (d *DynamicVariablesConfigWorkflowOverride) SetDynamicVariablePlaceholders(dynamicVariablePlaceholders map[string]any) {
	d.DynamicVariablePlaceholders = dynamicVariablePlaceholders
}

func (d *DynamicVariablesConfigWorkflowOverride) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DynamicVariablesConfigWorkflowOverride) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DynamicVariablesConfigWorkflowOverride to string"
	}
	return string(jsonData)
}

func (d *DynamicVariablesConfigWorkflowOverride) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DynamicVariablesConfigWorkflowOverride) MarshalJSON() ([]byte, error) {
	type embed DynamicVariablesConfigWorkflowOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
