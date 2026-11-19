package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Container for custom guardrails, matching ModerationGuardrail pattern
type CustomGuardrailInput struct {
	// Config container for custom guardrails list
	Config          *CustomGuardrailsConfigInput `json:"config,omitempty" xml:"config,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CustomGuardrailInput) GetConfig() (value *CustomGuardrailsConfigInput) {
	if c == nil {
		return
	}
	return c.Config
}

func (c *CustomGuardrailInput) SetConfig(config *CustomGuardrailsConfigInput) {
	c.Config = config
}

func (c *CustomGuardrailInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CustomGuardrailInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CustomGuardrailInput to string"
	}
	return string(jsonData)
}

func (c *CustomGuardrailInput) UnmarshalJSON(data []byte) error {
	type unmarshaler CustomGuardrailInput
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CustomGuardrailInput(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CustomGuardrailInput) MarshalJSON() ([]byte, error) {
	type embed CustomGuardrailInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
