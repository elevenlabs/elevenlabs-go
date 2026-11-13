package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SoftTimeoutConfigWorkflowOverride struct {
	// Time in seconds before showing the predefined message while waiting for LLM response. Set to -1 to disable.
	TimeoutSeconds *float64 `json:"timeout_seconds,omitempty" xml:"timeout_seconds,omitempty"`
	// Message to show when the first soft timeout is reached while waiting for LLM response. Supports dynamic variables (e.g., {{system__time}}, {{custom_variable}}).
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Extra static filler messages for subsequent soft timeouts in the same LLM generation. The first timeout uses `message`. If fewer messages are configured than `max_soft_timeouts_per_generation`, the last configured message is repeated; otherwise a built-in filler is used.
	AdditionalSoftTimeoutMessages []string `json:"additional_soft_timeout_messages,omitempty" xml:"additional_soft_timeout_messages,omitempty"`
	// If enabled, the soft timeout message will be generated dynamically instead of using the static message.
	UseLlmGeneratedMessage *bool `json:"use_llm_generated_message,omitempty" xml:"use_llm_generated_message,omitempty"`
	// If enabled, shuffle the order of static soft timeout messages once at the start of each turn. Only applies when use_llm_generated_message is false.
	RandomizeFillers *bool `json:"randomize_fillers,omitempty" xml:"randomize_fillers,omitempty"`
	// Maximum filler messages while waiting for a single LLM response. Fires every timeout_seconds until the LLM streams content or this limit is reached.
	MaxSoftTimeoutsPerGeneration *int `json:"max_soft_timeouts_per_generation,omitempty" xml:"max_soft_timeouts_per_generation,omitempty"`
	// Custom prompt for generating the soft timeout filler message when use_llm_generated_message is enabled. Recent conversation context is provided as a separate user message. If not set, the default prompt will be used. Supports dynamic variables (e.g., {{system__time}}, {{custom_variable}}).
	LlmGeneratedMessagePromptOverride *string `json:"llm_generated_message_prompt_override,omitempty" xml:"llm_generated_message_prompt_override,omitempty"`
	// When true, soft timeout fillers are suppressed until the conversation has at least one real user message. Prevents fillers during the agent's opening turn (e.g. workflow generate-immediately / tool calls before the user speaks).
	DisableUntilFirstUserMessage *bool `json:"disable_until_first_user_message,omitempty" xml:"disable_until_first_user_message,omitempty"`
	extraProperties              map[string]interface{}
}

func (s *SoftTimeoutConfigWorkflowOverride) GetTimeoutSeconds() (value *float64) {
	if s == nil {
		return
	}
	return s.TimeoutSeconds
}

func (s *SoftTimeoutConfigWorkflowOverride) GetMessage() (value *string) {
	if s == nil {
		return
	}
	return s.Message
}

func (s *SoftTimeoutConfigWorkflowOverride) GetAdditionalSoftTimeoutMessages() (value []string) {
	if s == nil {
		return
	}
	return s.AdditionalSoftTimeoutMessages
}

func (s *SoftTimeoutConfigWorkflowOverride) GetUseLlmGeneratedMessage() (value *bool) {
	if s == nil {
		return
	}
	return s.UseLlmGeneratedMessage
}

func (s *SoftTimeoutConfigWorkflowOverride) GetRandomizeFillers() (value *bool) {
	if s == nil {
		return
	}
	return s.RandomizeFillers
}

func (s *SoftTimeoutConfigWorkflowOverride) GetMaxSoftTimeoutsPerGeneration() (value *int) {
	if s == nil {
		return
	}
	return s.MaxSoftTimeoutsPerGeneration
}

func (s *SoftTimeoutConfigWorkflowOverride) GetLlmGeneratedMessagePromptOverride() (value *string) {
	if s == nil {
		return
	}
	return s.LlmGeneratedMessagePromptOverride
}

func (s *SoftTimeoutConfigWorkflowOverride) GetDisableUntilFirstUserMessage() (value *bool) {
	if s == nil {
		return
	}
	return s.DisableUntilFirstUserMessage
}

func (s *SoftTimeoutConfigWorkflowOverride) SetTimeoutSeconds(timeoutSeconds *float64) {
	s.TimeoutSeconds = timeoutSeconds
}

func (s *SoftTimeoutConfigWorkflowOverride) SetMessage(message *string) {
	s.Message = message
}

func (s *SoftTimeoutConfigWorkflowOverride) SetAdditionalSoftTimeoutMessages(additionalSoftTimeoutMessages []string) {
	s.AdditionalSoftTimeoutMessages = additionalSoftTimeoutMessages
}

func (s *SoftTimeoutConfigWorkflowOverride) SetUseLlmGeneratedMessage(useLlmGeneratedMessage *bool) {
	s.UseLlmGeneratedMessage = useLlmGeneratedMessage
}

func (s *SoftTimeoutConfigWorkflowOverride) SetRandomizeFillers(randomizeFillers *bool) {
	s.RandomizeFillers = randomizeFillers
}

func (s *SoftTimeoutConfigWorkflowOverride) SetMaxSoftTimeoutsPerGeneration(maxSoftTimeoutsPerGeneration *int) {
	s.MaxSoftTimeoutsPerGeneration = maxSoftTimeoutsPerGeneration
}

func (s *SoftTimeoutConfigWorkflowOverride) SetLlmGeneratedMessagePromptOverride(llmGeneratedMessagePromptOverride *string) {
	s.LlmGeneratedMessagePromptOverride = llmGeneratedMessagePromptOverride
}

func (s *SoftTimeoutConfigWorkflowOverride) SetDisableUntilFirstUserMessage(disableUntilFirstUserMessage *bool) {
	s.DisableUntilFirstUserMessage = disableUntilFirstUserMessage
}

func (s *SoftTimeoutConfigWorkflowOverride) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SoftTimeoutConfigWorkflowOverride) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SoftTimeoutConfigWorkflowOverride to string"
	}
	return string(jsonData)
}

func (s *SoftTimeoutConfigWorkflowOverride) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SoftTimeoutConfigWorkflowOverride) MarshalJSON() ([]byte, error) {
	type embed SoftTimeoutConfigWorkflowOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
