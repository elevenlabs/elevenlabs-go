package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Configuration for conversational events
type ConversationInitiationClientDataConversationConfigOverrideConversation struct {
	// If enabled audio will not be processed and only text will be used, use to avoid audio pricing.
	TextOnly *bool `json:"text_only,omitempty" xml:"text_only,omitempty"`
	// The maximum duration of a conversation in seconds
	MaxDurationSeconds *int `json:"max_duration_seconds,omitempty" xml:"max_duration_seconds,omitempty"`
	extraProperties    map[string]interface{}
}

func (c *ConversationInitiationClientDataConversationConfigOverrideConversation) GetTextOnly() (value *bool) {
	if c == nil {
		return
	}
	return c.TextOnly
}

func (c *ConversationInitiationClientDataConversationConfigOverrideConversation) GetMaxDurationSeconds() (value *int) {
	if c == nil {
		return
	}
	return c.MaxDurationSeconds
}

func (c *ConversationInitiationClientDataConversationConfigOverrideConversation) SetTextOnly(textOnly *bool) {
	c.TextOnly = textOnly
}

func (c *ConversationInitiationClientDataConversationConfigOverrideConversation) SetMaxDurationSeconds(maxDurationSeconds *int) {
	c.MaxDurationSeconds = maxDurationSeconds
}

func (c *ConversationInitiationClientDataConversationConfigOverrideConversation) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataConversationConfigOverrideConversation) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataConversationConfigOverrideConversation to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataConversationConfigOverrideConversation) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationInitiationClientDataConversationConfigOverrideConversation
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataConversationConfigOverrideConversation(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataConversationConfigOverrideConversation) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataConversationConfigOverrideConversation
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
