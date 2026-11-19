package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Config container for custom guardrails list
type CustomGuardrailsConfigOutput struct {
	Configs         []*CustomGuardrailConfig `json:"configs,omitempty" xml:"configs,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CustomGuardrailsConfigOutput) GetConfigs() (value []*CustomGuardrailConfig) {
	if c == nil {
		return
	}
	return c.Configs
}

func (c *CustomGuardrailsConfigOutput) SetConfigs(configs []*CustomGuardrailConfig) {
	c.Configs = configs
}

func (c *CustomGuardrailsConfigOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CustomGuardrailsConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CustomGuardrailsConfigOutput to string"
	}
	return string(jsonData)
}

func (c *CustomGuardrailsConfigOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler CustomGuardrailsConfigOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CustomGuardrailsConfigOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CustomGuardrailsConfigOutput) MarshalJSON() ([]byte, error) {
	type embed CustomGuardrailsConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
