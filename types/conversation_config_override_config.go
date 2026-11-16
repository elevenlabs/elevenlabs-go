package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationConfigOverrideConfig struct {
	// Whether to allow overriding the text_only field.
	TextOnly *bool `json:"text_only,omitempty" xml:"text_only,omitempty"`
	// Whether to allow overriding the max_duration_seconds field.
	MaxDurationSeconds *bool `json:"max_duration_seconds,omitempty" xml:"max_duration_seconds,omitempty"`
	extraProperties    map[string]interface{}
}

func (c *ConversationConfigOverrideConfig) GetTextOnly() (value *bool) {
	if c == nil {
		return
	}
	return c.TextOnly
}

func (c *ConversationConfigOverrideConfig) GetMaxDurationSeconds() (value *bool) {
	if c == nil {
		return
	}
	return c.MaxDurationSeconds
}

func (c *ConversationConfigOverrideConfig) SetTextOnly(textOnly *bool) {
	c.TextOnly = textOnly
}

func (c *ConversationConfigOverrideConfig) SetMaxDurationSeconds(maxDurationSeconds *bool) {
	c.MaxDurationSeconds = maxDurationSeconds
}

func (c *ConversationConfigOverrideConfig) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationConfigOverrideConfig) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationConfigOverrideConfig to string"
	}
	return string(jsonData)
}

func (c *ConversationConfigOverrideConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationConfigOverrideConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationConfigOverrideConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationConfigOverrideConfig) MarshalJSON() ([]byte, error) {
	type embed ConversationConfigOverrideConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
