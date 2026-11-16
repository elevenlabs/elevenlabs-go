package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentConfigOverrideInput struct {
	// If non-empty, the first message the agent will say. If empty, the agent waits for the user to start the discussion.
	FirstMessage *string `json:"first_message,omitempty" xml:"first_message,omitempty"`
	// Language of the agent - used for ASR and TTS
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// If non-empty, the message the agent will send when max conversation duration is reached.
	MaxConversationDurationMessage *string `json:"max_conversation_duration_message,omitempty" xml:"max_conversation_duration_message,omitempty"`
	// The prompt for the agent
	Prompt          *PromptAgentApiModelOverrideInput `json:"prompt,omitempty" xml:"prompt,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentConfigOverrideInput) GetFirstMessage() (value *string) {
	if a == nil {
		return
	}
	return a.FirstMessage
}

func (a *AgentConfigOverrideInput) GetLanguage() (value *string) {
	if a == nil {
		return
	}
	return a.Language
}

func (a *AgentConfigOverrideInput) GetMaxConversationDurationMessage() (value *string) {
	if a == nil {
		return
	}
	return a.MaxConversationDurationMessage
}

func (a *AgentConfigOverrideInput) GetPrompt() (value *PromptAgentApiModelOverrideInput) {
	if a == nil {
		return
	}
	return a.Prompt
}

func (a *AgentConfigOverrideInput) SetFirstMessage(firstMessage *string) {
	a.FirstMessage = firstMessage
}

func (a *AgentConfigOverrideInput) SetLanguage(language *string) {
	a.Language = language
}

func (a *AgentConfigOverrideInput) SetMaxConversationDurationMessage(maxConversationDurationMessage *string) {
	a.MaxConversationDurationMessage = maxConversationDurationMessage
}

func (a *AgentConfigOverrideInput) SetPrompt(prompt *PromptAgentApiModelOverrideInput) {
	a.Prompt = prompt
}

func (a *AgentConfigOverrideInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentConfigOverrideInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentConfigOverrideInput to string"
	}
	return string(jsonData)
}

func (a *AgentConfigOverrideInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentConfigOverrideInput) MarshalJSON() ([]byte, error) {
	type embed AgentConfigOverrideInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
