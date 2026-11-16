package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A system tool is a tool that is used to call a system method in the server
type SystemToolConfigInput struct {
	// The type of tool
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	Name string  `json:"name" xml:"name" required:"true" pattern:"^[a-zA-Z0-9_-]{1,64}$"`
	// Description of when the tool should be used and what it does. Leave empty to use the default description that's optimized for the specific tool type.
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
	ToolErrorHandlingMode *ToolErrorHandlingMode       `json:"tool_error_handling_mode,omitempty" xml:"tool_error_handling_mode,omitempty"`
	Params                *SystemToolConfigInputParams `json:"params" xml:"params" required:"true"`
	extraProperties       map[string]interface{}
}

func (s *SystemToolConfigInput) GetType() (value *string) {
	if s == nil {
		return
	}
	return s.Type
}

func (s *SystemToolConfigInput) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

func (s *SystemToolConfigInput) GetDescription() (value *string) {
	if s == nil {
		return
	}
	return s.Description
}

func (s *SystemToolConfigInput) GetResponseTimeoutSecs() (value *int) {
	if s == nil {
		return
	}
	return s.ResponseTimeoutSecs
}

func (s *SystemToolConfigInput) GetDisableInterruptions() (value *bool) {
	if s == nil {
		return
	}
	return s.DisableInterruptions
}

func (s *SystemToolConfigInput) GetInterruptionMode() (value *ToolInterruptionMode) {
	if s == nil {
		return
	}
	return s.InterruptionMode
}

func (s *SystemToolConfigInput) GetForcePreToolSpeech() (value *bool) {
	if s == nil {
		return
	}
	return s.ForcePreToolSpeech
}

func (s *SystemToolConfigInput) GetPreToolSpeech() (value *PreToolSpeechMode) {
	if s == nil {
		return
	}
	return s.PreToolSpeech
}

func (s *SystemToolConfigInput) GetAssignments() (value []*DynamicVariableAssignment) {
	if s == nil {
		return
	}
	return s.Assignments
}

func (s *SystemToolConfigInput) GetToolCallSound() (value *ToolCallSoundType) {
	if s == nil {
		return
	}
	return s.ToolCallSound
}

func (s *SystemToolConfigInput) GetToolCallSoundBehavior() (value *ToolCallSoundBehavior) {
	if s == nil {
		return
	}
	return s.ToolCallSoundBehavior
}

func (s *SystemToolConfigInput) GetToolErrorHandlingMode() (value *ToolErrorHandlingMode) {
	if s == nil {
		return
	}
	return s.ToolErrorHandlingMode
}

func (s *SystemToolConfigInput) GetParams() (value *SystemToolConfigInputParams) {
	if s == nil {
		return
	}
	return s.Params
}

func (s *SystemToolConfigInput) SetType(type_ *string) {
	s.Type = type_
}

func (s *SystemToolConfigInput) SetName(name string) {
	s.Name = name
}

func (s *SystemToolConfigInput) SetDescription(description *string) {
	s.Description = description
}

func (s *SystemToolConfigInput) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	s.ResponseTimeoutSecs = responseTimeoutSecs
}

func (s *SystemToolConfigInput) SetDisableInterruptions(disableInterruptions *bool) {
	s.DisableInterruptions = disableInterruptions
}

func (s *SystemToolConfigInput) SetInterruptionMode(interruptionMode *ToolInterruptionMode) {
	s.InterruptionMode = interruptionMode
}

func (s *SystemToolConfigInput) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	s.ForcePreToolSpeech = forcePreToolSpeech
}

func (s *SystemToolConfigInput) SetPreToolSpeech(preToolSpeech *PreToolSpeechMode) {
	s.PreToolSpeech = preToolSpeech
}

func (s *SystemToolConfigInput) SetAssignments(assignments []*DynamicVariableAssignment) {
	s.Assignments = assignments
}

func (s *SystemToolConfigInput) SetToolCallSound(toolCallSound *ToolCallSoundType) {
	s.ToolCallSound = toolCallSound
}

func (s *SystemToolConfigInput) SetToolCallSoundBehavior(toolCallSoundBehavior *ToolCallSoundBehavior) {
	s.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (s *SystemToolConfigInput) SetToolErrorHandlingMode(toolErrorHandlingMode *ToolErrorHandlingMode) {
	s.ToolErrorHandlingMode = toolErrorHandlingMode
}

func (s *SystemToolConfigInput) SetParams(params *SystemToolConfigInputParams) {
	s.Params = params
}

func (s *SystemToolConfigInput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SystemToolConfigInput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SystemToolConfigInput to string"
	}
	return string(jsonData)
}

func (s *SystemToolConfigInput) UnmarshalJSON(data []byte) error {
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

func (s *SystemToolConfigInput) MarshalJSON() ([]byte, error) {
	type embed SystemToolConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
