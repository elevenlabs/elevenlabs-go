package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Configuration for turn detection
type ConversationInitiationClientDataConversationConfigOverrideTurn struct {
	// Configuration for soft timeout functionality. Provides immediate feedback during longer LLM responses.
	SoftTimeoutConfig *ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig `json:"soft_timeout_config,omitempty" xml:"soft_timeout_config,omitempty"`
	// What the agent says when it detects a noisy environment.
	NoisyEnvironmentMessage *string `json:"noisy_environment_message,omitempty" xml:"noisy_environment_message,omitempty"`
	extraProperties         map[string]interface{}
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurn) GetSoftTimeoutConfig() (value *ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig) {
	if c == nil {
		return
	}
	return c.SoftTimeoutConfig
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurn) GetNoisyEnvironmentMessage() (value *string) {
	if c == nil {
		return
	}
	return c.NoisyEnvironmentMessage
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurn) SetSoftTimeoutConfig(softTimeoutConfig *ConversationInitiationClientDataConversationConfigOverrideTurnSoftTimeoutConfig) {
	c.SoftTimeoutConfig = softTimeoutConfig
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurn) SetNoisyEnvironmentMessage(noisyEnvironmentMessage *string) {
	c.NoisyEnvironmentMessage = noisyEnvironmentMessage
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurn) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataConversationConfigOverrideTurn) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataConversationConfigOverrideTurn to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurn) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationInitiationClientDataConversationConfigOverrideTurn
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataConversationConfigOverrideTurn(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTurn) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataConversationConfigOverrideTurn
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
