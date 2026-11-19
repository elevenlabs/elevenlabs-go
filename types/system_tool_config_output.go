package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A system tool is a tool that is used to call a system method in the server
type SystemToolConfigOutput struct {
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
	ToolErrorHandlingMode *ToolErrorHandlingMode        `json:"tool_error_handling_mode,omitempty" xml:"tool_error_handling_mode,omitempty"`
	Params                *SystemToolConfigOutputParams `json:"params" xml:"params" required:"true"`
	extraProperties       map[string]interface{}
}

func (s *SystemToolConfigOutput) GetType() (value *string) {
	if s == nil {
		return
	}
	return s.Type
}

func (s *SystemToolConfigOutput) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

func (s *SystemToolConfigOutput) GetDescription() (value *string) {
	if s == nil {
		return
	}
	return s.Description
}

func (s *SystemToolConfigOutput) GetResponseTimeoutSecs() (value *int) {
	if s == nil {
		return
	}
	return s.ResponseTimeoutSecs
}

func (s *SystemToolConfigOutput) GetDisableInterruptions() (value *bool) {
	if s == nil {
		return
	}
	return s.DisableInterruptions
}

func (s *SystemToolConfigOutput) GetInterruptionMode() (value *ToolInterruptionMode) {
	if s == nil {
		return
	}
	return s.InterruptionMode
}

func (s *SystemToolConfigOutput) GetForcePreToolSpeech() (value *bool) {
	if s == nil {
		return
	}
	return s.ForcePreToolSpeech
}

func (s *SystemToolConfigOutput) GetPreToolSpeech() (value *PreToolSpeechMode) {
	if s == nil {
		return
	}
	return s.PreToolSpeech
}

func (s *SystemToolConfigOutput) GetAssignments() (value []*DynamicVariableAssignment) {
	if s == nil {
		return
	}
	return s.Assignments
}

func (s *SystemToolConfigOutput) GetToolCallSound() (value *ToolCallSoundType) {
	if s == nil {
		return
	}
	return s.ToolCallSound
}

func (s *SystemToolConfigOutput) GetToolCallSoundBehavior() (value *ToolCallSoundBehavior) {
	if s == nil {
		return
	}
	return s.ToolCallSoundBehavior
}

func (s *SystemToolConfigOutput) GetToolErrorHandlingMode() (value *ToolErrorHandlingMode) {
	if s == nil {
		return
	}
	return s.ToolErrorHandlingMode
}

func (s *SystemToolConfigOutput) GetParams() (value *SystemToolConfigOutputParams) {
	if s == nil {
		return
	}
	return s.Params
}

func (s *SystemToolConfigOutput) SetType(type_ *string) {
	s.Type = type_
}

func (s *SystemToolConfigOutput) SetName(name string) {
	s.Name = name
}

func (s *SystemToolConfigOutput) SetDescription(description *string) {
	s.Description = description
}

func (s *SystemToolConfigOutput) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	s.ResponseTimeoutSecs = responseTimeoutSecs
}

func (s *SystemToolConfigOutput) SetDisableInterruptions(disableInterruptions *bool) {
	s.DisableInterruptions = disableInterruptions
}

func (s *SystemToolConfigOutput) SetInterruptionMode(interruptionMode *ToolInterruptionMode) {
	s.InterruptionMode = interruptionMode
}

func (s *SystemToolConfigOutput) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	s.ForcePreToolSpeech = forcePreToolSpeech
}

func (s *SystemToolConfigOutput) SetPreToolSpeech(preToolSpeech *PreToolSpeechMode) {
	s.PreToolSpeech = preToolSpeech
}

func (s *SystemToolConfigOutput) SetAssignments(assignments []*DynamicVariableAssignment) {
	s.Assignments = assignments
}

func (s *SystemToolConfigOutput) SetToolCallSound(toolCallSound *ToolCallSoundType) {
	s.ToolCallSound = toolCallSound
}

func (s *SystemToolConfigOutput) SetToolCallSoundBehavior(toolCallSoundBehavior *ToolCallSoundBehavior) {
	s.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (s *SystemToolConfigOutput) SetToolErrorHandlingMode(toolErrorHandlingMode *ToolErrorHandlingMode) {
	s.ToolErrorHandlingMode = toolErrorHandlingMode
}

func (s *SystemToolConfigOutput) SetParams(params *SystemToolConfigOutputParams) {
	s.Params = params
}

func (s *SystemToolConfigOutput) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SystemToolConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SystemToolConfigOutput to string"
	}
	return string(jsonData)
}

func (s *SystemToolConfigOutput) UnmarshalJSON(data []byte) error {
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

func (s *SystemToolConfigOutput) MarshalJSON() ([]byte, error) {
	type embed SystemToolConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
