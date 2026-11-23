package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Agent specific configuration
type ConversationInitiationClientDataConversationConfigOverrideAgent struct {
	// If non-empty, the first message the agent will say. If empty, the agent waits for the user to start the discussion.
	FirstMessage *string `json:"first_message,omitempty" xml:"first_message,omitempty"`
	// Language of the agent - used for ASR and TTS
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// If non-empty, the message the agent will send when max conversation duration is reached.
	MaxConversationDurationMessage *string `json:"max_conversation_duration_message,omitempty" xml:"max_conversation_duration_message,omitempty"`
	// The prompt for the agent
	Prompt          *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt `json:"prompt,omitempty" xml:"prompt,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgent) GetFirstMessage() (value *string) {
	if c == nil {
		return
	}
	return c.FirstMessage
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgent) GetLanguage() (value *string) {
	if c == nil {
		return
	}
	return c.Language
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgent) GetMaxConversationDurationMessage() (value *string) {
	if c == nil {
		return
	}
	return c.MaxConversationDurationMessage
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgent) GetPrompt() (value *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) {
	if c == nil {
		return
	}
	return c.Prompt
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgent) SetFirstMessage(firstMessage *string) {
	c.FirstMessage = firstMessage
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgent) SetLanguage(language *string) {
	c.Language = language
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgent) SetMaxConversationDurationMessage(maxConversationDurationMessage *string) {
	c.MaxConversationDurationMessage = maxConversationDurationMessage
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgent) SetPrompt(prompt *ConversationInitiationClientDataConversationConfigOverrideAgentPrompt) {
	c.Prompt = prompt
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgent) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataConversationConfigOverrideAgent) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataConversationConfigOverrideAgent to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgent) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationInitiationClientDataConversationConfigOverrideAgent
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataConversationConfigOverrideAgent(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataConversationConfigOverrideAgent) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataConversationConfigOverrideAgent
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
