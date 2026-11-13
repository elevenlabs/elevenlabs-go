package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentConfigApiModelWorkflowOverrideOutput struct {
	// If non-empty, the first message the agent will say. If empty, the agent waits for the user to start the discussion.
	FirstMessage *string `json:"first_message,omitempty" xml:"first_message,omitempty"`
	// Language of the agent - used for ASR and TTS
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	// When enabled, the agent mixes English words and phrases into responses in Hindi, Malayalam, Telugu, or Kannada.
	HinglishMode *bool `json:"hinglish_mode,omitempty" xml:"hinglish_mode,omitempty"`
	// Configuration for dynamic variables
	DynamicVariables *DynamicVariablesConfigWorkflowOverride `json:"dynamic_variables,omitempty" xml:"dynamic_variables,omitempty"`
	// If true, the user will not be able to interrupt the agent while the first message is being delivered.
	DisableFirstMessageInterruptions *bool `json:"disable_first_message_interruptions,omitempty" xml:"disable_first_message_interruptions,omitempty"`
	// If non-empty, the message the agent will send when max conversation duration is reached.
	MaxConversationDurationMessage *string `json:"max_conversation_duration_message,omitempty" xml:"max_conversation_duration_message,omitempty"`
	// Per-channel response behavior overrides for text conversations. Built-in channel defaults apply when unset.
	TextBehaviorOverrides map[string]*BehaviorOverride `json:"text_behavior_overrides,omitempty"`
	// The prompt for the agent
	Prompt          *PromptAgentApiModelWorkflowOverrideOutput `json:"prompt,omitempty" xml:"prompt,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) GetFirstMessage() (value *string) {
	if a == nil {
		return
	}
	return a.FirstMessage
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) GetLanguage() (value *string) {
	if a == nil {
		return
	}
	return a.Language
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) GetHinglishMode() (value *bool) {
	if a == nil {
		return
	}
	return a.HinglishMode
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) GetDynamicVariables() (value *DynamicVariablesConfigWorkflowOverride) {
	if a == nil {
		return
	}
	return a.DynamicVariables
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) GetDisableFirstMessageInterruptions() (value *bool) {
	if a == nil {
		return
	}
	return a.DisableFirstMessageInterruptions
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) GetMaxConversationDurationMessage() (value *string) {
	if a == nil {
		return
	}
	return a.MaxConversationDurationMessage
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) GetTextBehaviorOverrides() (value map[string]*BehaviorOverride) {
	if a == nil {
		return
	}
	return a.TextBehaviorOverrides
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) GetPrompt() (value *PromptAgentApiModelWorkflowOverrideOutput) {
	if a == nil {
		return
	}
	return a.Prompt
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) SetFirstMessage(firstMessage *string) {
	a.FirstMessage = firstMessage
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) SetLanguage(language *string) {
	a.Language = language
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) SetHinglishMode(hinglishMode *bool) {
	a.HinglishMode = hinglishMode
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) SetDynamicVariables(dynamicVariables *DynamicVariablesConfigWorkflowOverride) {
	a.DynamicVariables = dynamicVariables
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) SetDisableFirstMessageInterruptions(disableFirstMessageInterruptions *bool) {
	a.DisableFirstMessageInterruptions = disableFirstMessageInterruptions
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) SetMaxConversationDurationMessage(maxConversationDurationMessage *string) {
	a.MaxConversationDurationMessage = maxConversationDurationMessage
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) SetTextBehaviorOverrides(textBehaviorOverrides map[string]*BehaviorOverride) {
	a.TextBehaviorOverrides = textBehaviorOverrides
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) SetPrompt(prompt *PromptAgentApiModelWorkflowOverrideOutput) {
	a.Prompt = prompt
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentConfigApiModelWorkflowOverrideOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentConfigApiModelWorkflowOverrideOutput to string"
	}
	return string(jsonData)
}

func (a *AgentConfigApiModelWorkflowOverrideOutput) UnmarshalJSON(data []byte) error {
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

func (a *AgentConfigApiModelWorkflowOverrideOutput) MarshalJSON() ([]byte, error) {
	type embed AgentConfigApiModelWorkflowOverrideOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
