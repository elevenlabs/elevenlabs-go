package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A client tool is one that sends an event to the user's client to trigger something client side
type ClientToolConfigOutput struct {
	Name string `json:"name" xml:"name" required:"true" pattern:"^[a-zA-Z0-9_-]{1,64}$"`
	// Description of when the tool should be used and what it does.
	Description string `json:"description" xml:"description" required:"true"`
	// The maximum time in seconds to wait for the tool call to complete. Must be between 1 and 120 seconds (inclusive), or -1 to wait for the client's response indefinitely (requires expects_response).
	ResponseTimeoutSecs *int `json:"response_timeout_secs,omitempty" xml:"response_timeout_secs,omitempty" min:"-1" max:"120"`
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
	// Schema for any parameters to pass to the client
	Parameters *ObjectJsonSchemaPropertyOutput `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// If true, calling this tool should block the conversation until the client responds with some response which is passed to the llm. If false then we will continue the conversation without waiting for the client to respond, this is useful to show content to a user but not block the conversation
	ExpectsResponse  *bool                   `json:"expects_response,omitempty" xml:"expects_response,omitempty"`
	DynamicVariables *DynamicVariablesConfig `json:"dynamic_variables,omitempty" xml:"dynamic_variables,omitempty"`
	ExecutionMode    *ToolExecutionMode      `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	extraProperties  map[string]interface{}
}

func (c *ClientToolConfigOutput) GetName() (value string) {
	if c == nil {
		return
	}
	return c.Name
}

func (c *ClientToolConfigOutput) GetDescription() (value string) {
	if c == nil {
		return
	}
	return c.Description
}

func (c *ClientToolConfigOutput) GetResponseTimeoutSecs() (value *int) {
	if c == nil {
		return
	}
	return c.ResponseTimeoutSecs
}

func (c *ClientToolConfigOutput) GetDisableInterruptions() (value *bool) {
	if c == nil {
		return
	}
	return c.DisableInterruptions
}

func (c *ClientToolConfigOutput) GetInterruptionMode() (value *ToolInterruptionMode) {
	if c == nil {
		return
	}
	return c.InterruptionMode
}

func (c *ClientToolConfigOutput) GetForcePreToolSpeech() (value *bool) {
	if c == nil {
		return
	}
	return c.ForcePreToolSpeech
}

func (c *ClientToolConfigOutput) GetPreToolSpeech() (value *PreToolSpeechMode) {
	if c == nil {
		return
	}
	return c.PreToolSpeech
}

func (c *ClientToolConfigOutput) GetAssignments() (value []*DynamicVariableAssignment) {
	if c == nil {
		return
	}
	return c.Assignments
}

func (c *ClientToolConfigOutput) GetToolCallSound() (value *ToolCallSoundType) {
	if c == nil {
		return
	}
	return c.ToolCallSound
}

func (c *ClientToolConfigOutput) GetToolCallSoundBehavior() (value *ToolCallSoundBehavior) {
	if c == nil {
		return
	}
	return c.ToolCallSoundBehavior
}

func (c *ClientToolConfigOutput) GetToolErrorHandlingMode() (value *ToolErrorHandlingMode) {
	if c == nil {
		return
	}
	return c.ToolErrorHandlingMode
}

func (c *ClientToolConfigOutput) GetParameters() (value *ObjectJsonSchemaPropertyOutput) {
	if c == nil {
		return
	}
	return c.Parameters
}

func (c *ClientToolConfigOutput) GetExpectsResponse() (value *bool) {
	if c == nil {
		return
	}
	return c.ExpectsResponse
}

func (c *ClientToolConfigOutput) GetDynamicVariables() (value *DynamicVariablesConfig) {
	if c == nil {
		return
	}
	return c.DynamicVariables
}

func (c *ClientToolConfigOutput) GetExecutionMode() (value *ToolExecutionMode) {
	if c == nil {
		return
	}
	return c.ExecutionMode
}

func (c *ClientToolConfigOutput) SetName(name string) {
	c.Name = name
}

func (c *ClientToolConfigOutput) SetDescription(description string) {
	c.Description = description
}

func (c *ClientToolConfigOutput) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	c.ResponseTimeoutSecs = responseTimeoutSecs
}

func (c *ClientToolConfigOutput) SetDisableInterruptions(disableInterruptions *bool) {
	c.DisableInterruptions = disableInterruptions
}

func (c *ClientToolConfigOutput) SetInterruptionMode(interruptionMode *ToolInterruptionMode) {
	c.InterruptionMode = interruptionMode
}

func (c *ClientToolConfigOutput) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	c.ForcePreToolSpeech = forcePreToolSpeech
}

func (c *ClientToolConfigOutput) SetPreToolSpeech(preToolSpeech *PreToolSpeechMode) {
	c.PreToolSpeech = preToolSpeech
}

func (c *ClientToolConfigOutput) SetAssignments(assignments []*DynamicVariableAssignment) {
	c.Assignments = assignments
}

func (c *ClientToolConfigOutput) SetToolCallSound(toolCallSound *ToolCallSoundType) {
	c.ToolCallSound = toolCallSound
}

func (c *ClientToolConfigOutput) SetToolCallSoundBehavior(toolCallSoundBehavior *ToolCallSoundBehavior) {
	c.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (c *ClientToolConfigOutput) SetToolErrorHandlingMode(toolErrorHandlingMode *ToolErrorHandlingMode) {
	c.ToolErrorHandlingMode = toolErrorHandlingMode
}

func (c *ClientToolConfigOutput) SetParameters(parameters *ObjectJsonSchemaPropertyOutput) {
	c.Parameters = parameters
}

func (c *ClientToolConfigOutput) SetExpectsResponse(expectsResponse *bool) {
	c.ExpectsResponse = expectsResponse
}

func (c *ClientToolConfigOutput) SetDynamicVariables(dynamicVariables *DynamicVariablesConfig) {
	c.DynamicVariables = dynamicVariables
}

func (c *ClientToolConfigOutput) SetExecutionMode(executionMode *ToolExecutionMode) {
	c.ExecutionMode = executionMode
}

func (c *ClientToolConfigOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ClientToolConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ClientToolConfigOutput to string"
	}
	return string(jsonData)
}

func (c *ClientToolConfigOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ClientToolConfigOutput) MarshalJSON() ([]byte, error) {
	type embed ClientToolConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
