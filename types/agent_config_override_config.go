package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentConfigOverrideConfig struct {
	// Whether to allow overriding the first_message field.
	FirstMessage *bool `json:"first_message,omitempty" xml:"first_message,omitempty"`
	// Whether to allow overriding the language field.
	Language *bool `json:"language,omitempty" xml:"language,omitempty"`
	// Whether to allow overriding the max_conversation_duration_message field.
	MaxConversationDurationMessage *bool                              `json:"max_conversation_duration_message,omitempty" xml:"max_conversation_duration_message,omitempty"`
	Prompt                         *PromptAgentApiModelOverrideConfig `json:"prompt,omitempty" xml:"prompt,omitempty"`
	extraProperties                map[string]interface{}
}

func (a *AgentConfigOverrideConfig) GetFirstMessage() (value *bool) {
	if a == nil {
		return
	}
	return a.FirstMessage
}

func (a *AgentConfigOverrideConfig) GetLanguage() (value *bool) {
	if a == nil {
		return
	}
	return a.Language
}

func (a *AgentConfigOverrideConfig) GetMaxConversationDurationMessage() (value *bool) {
	if a == nil {
		return
	}
	return a.MaxConversationDurationMessage
}

func (a *AgentConfigOverrideConfig) GetPrompt() (value *PromptAgentApiModelOverrideConfig) {
	if a == nil {
		return
	}
	return a.Prompt
}

func (a *AgentConfigOverrideConfig) SetFirstMessage(firstMessage *bool) {
	a.FirstMessage = firstMessage
}

func (a *AgentConfigOverrideConfig) SetLanguage(language *bool) {
	a.Language = language
}

func (a *AgentConfigOverrideConfig) SetMaxConversationDurationMessage(maxConversationDurationMessage *bool) {
	a.MaxConversationDurationMessage = maxConversationDurationMessage
}

func (a *AgentConfigOverrideConfig) SetPrompt(prompt *PromptAgentApiModelOverrideConfig) {
	a.Prompt = prompt
}

func (a *AgentConfigOverrideConfig) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentConfigOverrideConfig) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentConfigOverrideConfig to string"
	}
	return string(jsonData)
}

func (a *AgentConfigOverrideConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler AgentConfigOverrideConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AgentConfigOverrideConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentConfigOverrideConfig) MarshalJSON() ([]byte, error) {
	type embed AgentConfigOverrideConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
