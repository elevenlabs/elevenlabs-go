package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentConfigOverrideOutput struct {
	// If non-empty, the first message the agent will say. If empty, the agent waits for the user to start the discussion.
	FirstMessage *string `json:"first_message,omitempty" xml:"first_message,omitempty"`
	// Language of the agent - used for ASR and TTS
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// If non-empty, the message the agent will send when max conversation duration is reached.
	MaxConversationDurationMessage *string `json:"max_conversation_duration_message,omitempty" xml:"max_conversation_duration_message,omitempty"`
	// The prompt for the agent
	Prompt          *PromptAgentApiModelOverrideOutput `json:"prompt,omitempty" xml:"prompt,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentConfigOverrideOutput) GetFirstMessage() (value *string) {
	if a == nil {
		return
	}
	return a.FirstMessage
}

func (a *AgentConfigOverrideOutput) GetLanguage() (value *string) {
	if a == nil {
		return
	}
	return a.Language
}

func (a *AgentConfigOverrideOutput) GetMaxConversationDurationMessage() (value *string) {
	if a == nil {
		return
	}
	return a.MaxConversationDurationMessage
}

func (a *AgentConfigOverrideOutput) GetPrompt() (value *PromptAgentApiModelOverrideOutput) {
	if a == nil {
		return
	}
	return a.Prompt
}

func (a *AgentConfigOverrideOutput) SetFirstMessage(firstMessage *string) {
	a.FirstMessage = firstMessage
}

func (a *AgentConfigOverrideOutput) SetLanguage(language *string) {
	a.Language = language
}

func (a *AgentConfigOverrideOutput) SetMaxConversationDurationMessage(maxConversationDurationMessage *string) {
	a.MaxConversationDurationMessage = maxConversationDurationMessage
}

func (a *AgentConfigOverrideOutput) SetPrompt(prompt *PromptAgentApiModelOverrideOutput) {
	a.Prompt = prompt
}

func (a *AgentConfigOverrideOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentConfigOverrideOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentConfigOverrideOutput to string"
	}
	return string(jsonData)
}

func (a *AgentConfigOverrideOutput) UnmarshalJSON(data []byte) error {
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

func (a *AgentConfigOverrideOutput) MarshalJSON() ([]byte, error) {
	type embed AgentConfigOverrideOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
