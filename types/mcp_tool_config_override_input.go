package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type McpToolConfigOverrideInput struct {
	// The name of the MCP tool
	ToolName string `json:"tool_name" xml:"tool_name" required:"true"`
	// DEPRECATED: use `pre_tool_speech` instead. If set, overrides the server's force_pre_tool_speech setting for this tool.
	ForcePreToolSpeech *bool `json:"force_pre_tool_speech,omitempty" xml:"force_pre_tool_speech,omitempty"`
	// If set, overrides the server's pre_tool_speech setting for this tool.
	PreToolSpeech *PreToolSpeechMode `json:"pre_tool_speech,omitempty" xml:"pre_tool_speech,omitempty"`
	// DEPRECATED: use `interruption_mode` instead. If set, overrides the server's disable_interruptions setting for this tool.
	DisableInterruptions *bool `json:"disable_interruptions,omitempty" xml:"disable_interruptions,omitempty"`
	// If set, overrides the server's interruption_mode setting for this tool.
	InterruptionMode *ToolInterruptionMode `json:"interruption_mode,omitempty" xml:"interruption_mode,omitempty"`
	// Overrides the server's tool_call_sound setting for this tool. A sound name plays that sound; 'off' overrides to no sound (silence); null means do not override (inherit the server default).
	ToolCallSound *McpToolConfigOverrideInputToolCallSound `json:"tool_call_sound,omitempty" xml:"tool_call_sound,omitempty"`
	// If set, overrides the server's tool_call_sound_behavior setting for this tool
	ToolCallSoundBehavior *ToolCallSoundBehavior `json:"tool_call_sound_behavior,omitempty" xml:"tool_call_sound_behavior,omitempty"`
	// If set, overrides the server's execution_mode setting for this tool
	ExecutionMode *ToolExecutionMode `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	// If set, overrides the server's response timeout for this MCP tool (seconds).
	ResponseTimeoutSecs *int `json:"response_timeout_secs,omitempty" xml:"response_timeout_secs,omitempty" min:"5" max:"300"`
	// Dynamic variable assignments for this MCP tool
	Assignments []*DynamicVariableAssignment `json:"assignments,omitempty" xml:"assignments,omitempty"`
	// Mapping of json path to input override configuration
	InputOverrides map[string]*McpToolConfigOverrideInputInputOverridesValue `json:"input_overrides,omitempty"`
	// Mock responses with optional parameter conditions. Evaluated top-to-bottom; first match wins.
	ResponseMocks   []*ToolResponseMockConfigInput `json:"response_mocks,omitempty" xml:"response_mocks,omitempty"`
	extraProperties map[string]interface{}
}

func (m *McpToolConfigOverrideInput) GetToolName() (value string) {
	if m == nil {
		return
	}
	return m.ToolName
}

func (m *McpToolConfigOverrideInput) GetForcePreToolSpeech() (value *bool) {
	if m == nil {
		return
	}
	return m.ForcePreToolSpeech
}

func (m *McpToolConfigOverrideInput) GetPreToolSpeech() (value *PreToolSpeechMode) {
	if m == nil {
		return
	}
	return m.PreToolSpeech
}

func (m *McpToolConfigOverrideInput) GetDisableInterruptions() (value *bool) {
	if m == nil {
		return
	}
	return m.DisableInterruptions
}

func (m *McpToolConfigOverrideInput) GetInterruptionMode() (value *ToolInterruptionMode) {
	if m == nil {
		return
	}
	return m.InterruptionMode
}

func (m *McpToolConfigOverrideInput) GetToolCallSound() (value *McpToolConfigOverrideInputToolCallSound) {
	if m == nil {
		return
	}
	return m.ToolCallSound
}

func (m *McpToolConfigOverrideInput) GetToolCallSoundBehavior() (value *ToolCallSoundBehavior) {
	if m == nil {
		return
	}
	return m.ToolCallSoundBehavior
}

func (m *McpToolConfigOverrideInput) GetExecutionMode() (value *ToolExecutionMode) {
	if m == nil {
		return
	}
	return m.ExecutionMode
}

func (m *McpToolConfigOverrideInput) GetResponseTimeoutSecs() (value *int) {
	if m == nil {
		return
	}
	return m.ResponseTimeoutSecs
}

func (m *McpToolConfigOverrideInput) GetAssignments() (value []*DynamicVariableAssignment) {
	if m == nil {
		return
	}
	return m.Assignments
}

func (m *McpToolConfigOverrideInput) GetInputOverrides() (value map[string]*McpToolConfigOverrideInputInputOverridesValue) {
	if m == nil {
		return
	}
	return m.InputOverrides
}

func (m *McpToolConfigOverrideInput) GetResponseMocks() (value []*ToolResponseMockConfigInput) {
	if m == nil {
		return
	}
	return m.ResponseMocks
}

func (m *McpToolConfigOverrideInput) SetToolName(toolName string) {
	m.ToolName = toolName
}

func (m *McpToolConfigOverrideInput) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	m.ForcePreToolSpeech = forcePreToolSpeech
}

func (m *McpToolConfigOverrideInput) SetPreToolSpeech(preToolSpeech *PreToolSpeechMode) {
	m.PreToolSpeech = preToolSpeech
}

func (m *McpToolConfigOverrideInput) SetDisableInterruptions(disableInterruptions *bool) {
	m.DisableInterruptions = disableInterruptions
}

func (m *McpToolConfigOverrideInput) SetInterruptionMode(interruptionMode *ToolInterruptionMode) {
	m.InterruptionMode = interruptionMode
}

func (m *McpToolConfigOverrideInput) SetToolCallSound(toolCallSound *McpToolConfigOverrideInputToolCallSound) {
	m.ToolCallSound = toolCallSound
}

func (m *McpToolConfigOverrideInput) SetToolCallSoundBehavior(toolCallSoundBehavior *ToolCallSoundBehavior) {
	m.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (m *McpToolConfigOverrideInput) SetExecutionMode(executionMode *ToolExecutionMode) {
	m.ExecutionMode = executionMode
}

func (m *McpToolConfigOverrideInput) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	m.ResponseTimeoutSecs = responseTimeoutSecs
}

func (m *McpToolConfigOverrideInput) SetAssignments(assignments []*DynamicVariableAssignment) {
	m.Assignments = assignments
}

func (m *McpToolConfigOverrideInput) SetInputOverrides(inputOverrides map[string]*McpToolConfigOverrideInputInputOverridesValue) {
	m.InputOverrides = inputOverrides
}

func (m *McpToolConfigOverrideInput) SetResponseMocks(responseMocks []*ToolResponseMockConfigInput) {
	m.ResponseMocks = responseMocks
}

func (m *McpToolConfigOverrideInput) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpToolConfigOverrideInput) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolConfigOverrideInput to string"
	}
	return string(jsonData)
}

func (m *McpToolConfigOverrideInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpToolConfigOverrideInput) MarshalJSON() ([]byte, error) {
	type embed McpToolConfigOverrideInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
