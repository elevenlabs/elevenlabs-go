package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationConfigOverride struct {
	// If enabled audio will not be processed and only text will be used, use to avoid audio pricing.
	TextOnly *bool `json:"text_only,omitempty" xml:"text_only,omitempty"`
	// The maximum duration of a conversation in seconds
	MaxDurationSeconds *int `json:"max_duration_seconds,omitempty" xml:"max_duration_seconds,omitempty"`
	extraProperties    map[string]interface{}
}

func (c *ConversationConfigOverride) GetTextOnly() (value *bool) {
	if c == nil {
		return
	}
	return c.TextOnly
}

func (c *ConversationConfigOverride) GetMaxDurationSeconds() (value *int) {
	if c == nil {
		return
	}
	return c.MaxDurationSeconds
}

func (c *ConversationConfigOverride) SetTextOnly(textOnly *bool) {
	c.TextOnly = textOnly
}

func (c *ConversationConfigOverride) SetMaxDurationSeconds(maxDurationSeconds *int) {
	c.MaxDurationSeconds = maxDurationSeconds
}

func (c *ConversationConfigOverride) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationConfigOverride) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationConfigOverride to string"
	}
	return string(jsonData)
}

func (c *ConversationConfigOverride) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationConfigOverride) MarshalJSON() ([]byte, error) {
	type embed ConversationConfigOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
