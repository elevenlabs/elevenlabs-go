package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// SMB tool configuration that wraps SMB tool parameters.
type SmbToolConfig struct {
	Name string `json:"name" xml:"name" required:"true" pattern:"^[a-zA-Z0-9_-]{1,64}$"`
	// Description of when the tool should be used and what it does.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The maximum time in seconds to wait for the tool call to complete.
	ResponseTimeoutSecs *int `json:"response_timeout_secs,omitempty" xml:"response_timeout_secs,omitempty"`
	// DEPRECATED: use `interruption_mode` instead. If true, the user will not be able to interrupt the agent while this tool is running.
	DisableInterruptions *bool                 `json:"disable_interruptions,omitempty" xml:"disable_interruptions,omitempty"`
	InterruptionMode     *ToolInterruptionMode `json:"interruption_mode,omitempty" xml:"interruption_mode,omitempty"`
	// DEPRECATED: use `pre_tool_speech` instead. If true, the agent will speak before the tool call.
	ForcePreToolSpeech *bool              `json:"force_pre_tool_speech,omitempty" xml:"force_pre_tool_speech,omitempty"`
	PreToolSpeech      *PreToolSpeechMode `json:"pre_tool_speech,omitempty" xml:"pre_tool_speech,omitempty"`
	// Configuration for extracting values from tool responses and assigning them to dynamic variables
	Assignments []*DynamicVariableAssignment `json:"assignments,omitempty" xml:"assignments,omitempty"`
	// Predefined tool call sound type to play during tool execution. If not specified, no tool call sound will be played.
	ToolCallSound *ToolCallSoundType `json:"tool_call_sound,omitempty" xml:"tool_call_sound,omitempty"`
	// Determines how the tool call sound should be played.
	ToolCallSoundBehavior *ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitempty" xml:"tool_call_sound_behavior,omitempty"`
	// Controls how tool errors are processed before being shared with the agent.
	ToolErrorHandlingMode *ToolErrorHandlingMode `json:"tool_error_handling_mode,omitempty" xml:"tool_error_handling_mode,omitempty"`
	// User-facing tooltip for the procedure reference picker. Not sent to the LLM.
	HumanDescription *string `json:"human_description,omitempty" xml:"human_description,omitempty"`
	// Whether this tool is enabled for the agent
	Enabled         *bool                `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Params          *SmbToolConfigParams `json:"params" xml:"params" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SmbToolConfig) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

func (s *SmbToolConfig) GetDescription() (value *string) {
	if s == nil {
		return
	}
	return s.Description
}

func (s *SmbToolConfig) GetResponseTimeoutSecs() (value *int) {
	if s == nil {
		return
	}
	return s.ResponseTimeoutSecs
}

func (s *SmbToolConfig) GetDisableInterruptions() (value *bool) {
	if s == nil {
		return
	}
	return s.DisableInterruptions
}

func (s *SmbToolConfig) GetInterruptionMode() (value *ToolInterruptionMode) {
	if s == nil {
		return
	}
	return s.InterruptionMode
}

func (s *SmbToolConfig) GetForcePreToolSpeech() (value *bool) {
	if s == nil {
		return
	}
	return s.ForcePreToolSpeech
}

func (s *SmbToolConfig) GetPreToolSpeech() (value *PreToolSpeechMode) {
	if s == nil {
		return
	}
	return s.PreToolSpeech
}

func (s *SmbToolConfig) GetAssignments() (value []*DynamicVariableAssignment) {
	if s == nil {
		return
	}
	return s.Assignments
}

func (s *SmbToolConfig) GetToolCallSound() (value *ToolCallSoundType) {
	if s == nil {
		return
	}
	return s.ToolCallSound
}

func (s *SmbToolConfig) GetToolCallSoundBehavior() (value *ToolCallSoundBehavior) {
	if s == nil {
		return
	}
	return s.ToolCallSoundBehavior
}

func (s *SmbToolConfig) GetToolErrorHandlingMode() (value *ToolErrorHandlingMode) {
	if s == nil {
		return
	}
	return s.ToolErrorHandlingMode
}

func (s *SmbToolConfig) GetHumanDescription() (value *string) {
	if s == nil {
		return
	}
	return s.HumanDescription
}

func (s *SmbToolConfig) GetEnabled() (value *bool) {
	if s == nil {
		return
	}
	return s.Enabled
}

func (s *SmbToolConfig) GetParams() (value *SmbToolConfigParams) {
	if s == nil {
		return
	}
	return s.Params
}

func (s *SmbToolConfig) SetName(name string) {
	s.Name = name
}

func (s *SmbToolConfig) SetDescription(description *string) {
	s.Description = description
}

func (s *SmbToolConfig) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	s.ResponseTimeoutSecs = responseTimeoutSecs
}

func (s *SmbToolConfig) SetDisableInterruptions(disableInterruptions *bool) {
	s.DisableInterruptions = disableInterruptions
}

func (s *SmbToolConfig) SetInterruptionMode(interruptionMode *ToolInterruptionMode) {
	s.InterruptionMode = interruptionMode
}

func (s *SmbToolConfig) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	s.ForcePreToolSpeech = forcePreToolSpeech
}

func (s *SmbToolConfig) SetPreToolSpeech(preToolSpeech *PreToolSpeechMode) {
	s.PreToolSpeech = preToolSpeech
}

func (s *SmbToolConfig) SetAssignments(assignments []*DynamicVariableAssignment) {
	s.Assignments = assignments
}

func (s *SmbToolConfig) SetToolCallSound(toolCallSound *ToolCallSoundType) {
	s.ToolCallSound = toolCallSound
}

func (s *SmbToolConfig) SetToolCallSoundBehavior(toolCallSoundBehavior *ToolCallSoundBehavior) {
	s.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (s *SmbToolConfig) SetToolErrorHandlingMode(toolErrorHandlingMode *ToolErrorHandlingMode) {
	s.ToolErrorHandlingMode = toolErrorHandlingMode
}

func (s *SmbToolConfig) SetHumanDescription(humanDescription *string) {
	s.HumanDescription = humanDescription
}

func (s *SmbToolConfig) SetEnabled(enabled *bool) {
	s.Enabled = enabled
}

func (s *SmbToolConfig) SetParams(params *SmbToolConfigParams) {
	s.Params = params
}

func (s *SmbToolConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SmbToolConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SmbToolConfig to string"
	}
	return string(jsonData)
}

func (s *SmbToolConfig) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SmbToolConfig) MarshalJSON() ([]byte, error) {
	type embed SmbToolConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
