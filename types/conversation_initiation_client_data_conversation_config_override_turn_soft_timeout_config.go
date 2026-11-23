package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Configuration for soft timeout functionality. Provides immediate feedback during longer LLM responses.
type ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig struct {
	// Message to show when the first soft timeout is reached while waiting for LLM response. Supports dynamic variables (e.g., {{system__time}}, {{custom_variable}}).
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Extra static filler messages for subsequent soft timeouts in the same LLM generation. The first timeout uses `message`. If fewer messages are configured than `max_soft_timeouts_per_generation`, the last configured message is repeated; otherwise a built-in filler is used.
	AdditionalSoftTimeoutMessages []string `json:"additional_soft_timeout_messages,omitempty" xml:"additional_soft_timeout_messages,omitempty"`
	extraProperties               map[string]interface{}
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig) GetMessage() (value *string) {
	if c == nil {
		return
	}
	return c.Message
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig) GetAdditionalSoftTimeoutMessages() (value []string) {
	if c == nil {
		return
	}
	return c.AdditionalSoftTimeoutMessages
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig) SetMessage(message *string) {
	c.Message = message
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig) SetAdditionalSoftTimeoutMessages(additionalSoftTimeoutMessages []string) {
	c.AdditionalSoftTimeoutMessages = additionalSoftTimeoutMessages
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
