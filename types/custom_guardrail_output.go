package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Container for custom guardrails, matching ModerationGuardrail pattern
type CustomGuardrailOutput struct {
	// Config container for custom guardrails list
	Config          *CustomGuardrailsConfigOutput `json:"config,omitempty" xml:"config,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CustomGuardrailOutput) GetConfig() (value *CustomGuardrailsConfigOutput) {
	if c == nil {
		return
	}
	return c.Config
}

func (c *CustomGuardrailOutput) SetConfig(config *CustomGuardrailsConfigOutput) {
	c.Config = config
}

func (c *CustomGuardrailOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CustomGuardrailOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CustomGuardrailOutput to string"
	}
	return string(jsonData)
}

func (c *CustomGuardrailOutput) UnmarshalJSON(data []byte) error {
	type unmarshaler CustomGuardrailOutput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CustomGuardrailOutput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CustomGuardrailOutput) MarshalJSON() ([]byte, error) {
	type embed CustomGuardrailOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
