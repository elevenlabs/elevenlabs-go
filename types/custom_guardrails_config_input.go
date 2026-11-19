package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Config container for custom guardrails list
type CustomGuardrailsConfigInput struct {
	Configs         []*CustomGuardrailConfig `json:"configs,omitempty" xml:"configs,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CustomGuardrailsConfigInput) GetConfigs() (value []*CustomGuardrailConfig) {
	if c == nil {
		return
	}
	return c.Configs
}

func (c *CustomGuardrailsConfigInput) SetConfigs(configs []*CustomGuardrailConfig) {
	c.Configs = configs
}

func (c *CustomGuardrailsConfigInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CustomGuardrailsConfigInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CustomGuardrailsConfigInput to string"
	}
	return string(jsonData)
}

func (c *CustomGuardrailsConfigInput) UnmarshalJSON(data []byte) error {
	type unmarshaler CustomGuardrailsConfigInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CustomGuardrailsConfigInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CustomGuardrailsConfigInput) MarshalJSON() ([]byte, error) {
	type embed CustomGuardrailsConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
