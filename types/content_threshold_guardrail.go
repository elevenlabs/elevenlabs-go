package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ContentThresholdGuardrail struct {
	IsEnabled       *bool                               `json:"is_enabled,omitempty" xml:"is_enabled,omitempty"`
	Threshold       *ContentThresholdGuardrailThreshold `json:"threshold,omitempty" xml:"threshold,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ContentThresholdGuardrail) GetIsEnabled() (value *bool) {
	if c == nil {
		return
	}
	return c.IsEnabled
}

func (c *ContentThresholdGuardrail) GetThreshold() (value *ContentThresholdGuardrailThreshold) {
	if c == nil {
		return
	}
	return c.Threshold
}

func (c *ContentThresholdGuardrail) SetIsEnabled(isEnabled *bool) {
	c.IsEnabled = isEnabled
}

func (c *ContentThresholdGuardrail) SetThreshold(threshold *ContentThresholdGuardrailThreshold) {
	c.Threshold = threshold
}

func (c *ContentThresholdGuardrail) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ContentThresholdGuardrail) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ContentThresholdGuardrail to string"
	}
	return string(jsonData)
}

func (c *ContentThresholdGuardrail) UnmarshalJSON(data []byte) error {
	type unmarshaler ContentThresholdGuardrail
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ContentThresholdGuardrail(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ContentThresholdGuardrail) MarshalJSON() ([]byte, error) {
	type embed ContentThresholdGuardrail
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
