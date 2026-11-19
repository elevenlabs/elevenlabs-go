package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// An MCP tool configuration that can be used to call MCP servers
type McpToolConfigInput struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	Name string  `json:"name" xml:"name" required:"true" pattern:"^[a-zA-Z0-9_-]{1,64}$"`
	// Description of when the tool should be used and what it does.
	Description string `json:"description" xml:"description" required:"true"`
	// The maximum time in seconds to wait for the MCP tool call to complete. Must be between 5 and 300 seconds (inclusive).
	ResponseTimeoutSecs *int `json:"response_timeout_secs,omitempty" xml:"response_timeout_secs,omitempty" min:"5" max:"300"`
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
	IntegrationType       IntegrationType        `json:"integration_type" xml:"integration_type" required:"true"`
	// Schema for any parameters the LLM needs to provide to the MCP tool.
	Parameters *ObjectJsonSchemaPropertyInput `json:"parameters,omitempty" xml:"parameters,omitempty"`
	// Defines the MCP server-level approval policy for tool execution.
	ApprovalPolicy *McpApprovalPolicy `json:"approval_policy,omitempty" xml:"approval_policy,omitempty"`
	// The name of the MCP tool to call
	McpToolName string `json:"mcp_tool_name" xml:"mcp_tool_name" required:"true"`
	// The description of the MCP tool to call
	McpToolDescription string `json:"mcp_tool_description" xml:"mcp_tool_description" required:"true"`
	// The id of the MCP server to call
	McpServerId string `json:"mcp_server_id" xml:"mcp_server_id" required:"true"`
	// The name of the MCP server to call
	McpServerName string `json:"mcp_server_name" xml:"mcp_server_name" required:"true"`
	// Original inputSchema dict for consistent hashing
	McpInputSchema map[string]any     `json:"mcp_input_schema,omitempty"`
	ExecutionMode  *ToolExecutionMode `json:"execution_mode,omitempty" xml:"execution_mode,omitempty"`
	// Input parameter overrides for this tool
	InputOverrides  map[string]*McpToolConfigInputInputOverridesValue `json:"input_overrides,omitempty"`
	extraProperties map[string]interface{}
}

func (m *McpToolConfigInput) GetType() (value *string) {
	if m == nil {
		return
	}
	return m.Type
}

func (m *McpToolConfigInput) GetName() (value string) {
	if m == nil {
		return
	}
	return m.Name
}

func (m *McpToolConfigInput) GetDescription() (value string) {
	if m == nil {
		return
	}
	return m.Description
}

func (m *McpToolConfigInput) GetResponseTimeoutSecs() (value *int) {
	if m == nil {
		return
	}
	return m.ResponseTimeoutSecs
}

func (m *McpToolConfigInput) GetDisableInterruptions() (value *bool) {
	if m == nil {
		return
	}
	return m.DisableInterruptions
}

func (m *McpToolConfigInput) GetInterruptionMode() (value *ToolInterruptionMode) {
	if m == nil {
		return
	}
	return m.InterruptionMode
}

func (m *McpToolConfigInput) GetForcePreToolSpeech() (value *bool) {
	if m == nil {
		return
	}
	return m.ForcePreToolSpeech
}

func (m *McpToolConfigInput) GetPreToolSpeech() (value *PreToolSpeechMode) {
	if m == nil {
		return
	}
	return m.PreToolSpeech
}

func (m *McpToolConfigInput) GetAssignments() (value []*DynamicVariableAssignment) {
	if m == nil {
		return
	}
	return m.Assignments
}

func (m *McpToolConfigInput) GetToolCallSound() (value *ToolCallSoundType) {
	if m == nil {
		return
	}
	return m.ToolCallSound
}

func (m *McpToolConfigInput) GetToolCallSoundBehavior() (value *ToolCallSoundBehavior) {
	if m == nil {
		return
	}
	return m.ToolCallSoundBehavior
}

func (m *McpToolConfigInput) GetToolErrorHandlingMode() (value *ToolErrorHandlingMode) {
	if m == nil {
		return
	}
	return m.ToolErrorHandlingMode
}

func (m *McpToolConfigInput) GetIntegrationType() (value IntegrationType) {
	if m == nil {
		return
	}
	return m.IntegrationType
}

func (m *McpToolConfigInput) GetParameters() (value *ObjectJsonSchemaPropertyInput) {
	if m == nil {
		return
	}
	return m.Parameters
}

func (m *McpToolConfigInput) GetApprovalPolicy() (value *McpApprovalPolicy) {
	if m == nil {
		return
	}
	return m.ApprovalPolicy
}

func (m *McpToolConfigInput) GetMcpToolName() (value string) {
	if m == nil {
		return
	}
	return m.McpToolName
}

func (m *McpToolConfigInput) GetMcpToolDescription() (value string) {
	if m == nil {
		return
	}
	return m.McpToolDescription
}

func (m *McpToolConfigInput) GetMcpServerId() (value string) {
	if m == nil {
		return
	}
	return m.McpServerId
}

func (m *McpToolConfigInput) GetMcpServerName() (value string) {
	if m == nil {
		return
	}
	return m.McpServerName
}

func (m *McpToolConfigInput) GetMcpInputSchema() (value map[string]any) {
	if m == nil {
		return
	}
	return m.McpInputSchema
}

func (m *McpToolConfigInput) GetExecutionMode() (value *ToolExecutionMode) {
	if m == nil {
		return
	}
	return m.ExecutionMode
}

func (m *McpToolConfigInput) GetInputOverrides() (value map[string]*McpToolConfigInputInputOverridesValue) {
	if m == nil {
		return
	}
	return m.InputOverrides
}

func (m *McpToolConfigInput) SetType(type_ *string) {
	m.Type = type_
}

func (m *McpToolConfigInput) SetName(name string) {
	m.Name = name
}

func (m *McpToolConfigInput) SetDescription(description string) {
	m.Description = description
}

func (m *McpToolConfigInput) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	m.ResponseTimeoutSecs = responseTimeoutSecs
}

func (m *McpToolConfigInput) SetDisableInterruptions(disableInterruptions *bool) {
	m.DisableInterruptions = disableInterruptions
}

func (m *McpToolConfigInput) SetInterruptionMode(interruptionMode *ToolInterruptionMode) {
	m.InterruptionMode = interruptionMode
}

func (m *McpToolConfigInput) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	m.ForcePreToolSpeech = forcePreToolSpeech
}

func (m *McpToolConfigInput) SetPreToolSpeech(preToolSpeech *PreToolSpeechMode) {
	m.PreToolSpeech = preToolSpeech
}

func (m *McpToolConfigInput) SetAssignments(assignments []*DynamicVariableAssignment) {
	m.Assignments = assignments
}

func (m *McpToolConfigInput) SetToolCallSound(toolCallSound *ToolCallSoundType) {
	m.ToolCallSound = toolCallSound
}

func (m *McpToolConfigInput) SetToolCallSoundBehavior(toolCallSoundBehavior *ToolCallSoundBehavior) {
	m.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (m *McpToolConfigInput) SetToolErrorHandlingMode(toolErrorHandlingMode *ToolErrorHandlingMode) {
	m.ToolErrorHandlingMode = toolErrorHandlingMode
}

func (m *McpToolConfigInput) SetIntegrationType(integrationType IntegrationType) {
	m.IntegrationType = integrationType
}

func (m *McpToolConfigInput) SetParameters(parameters *ObjectJsonSchemaPropertyInput) {
	m.Parameters = parameters
}

func (m *McpToolConfigInput) SetApprovalPolicy(approvalPolicy *McpApprovalPolicy) {
	m.ApprovalPolicy = approvalPolicy
}

func (m *McpToolConfigInput) SetMcpToolName(mcpToolName string) {
	m.McpToolName = mcpToolName
}

func (m *McpToolConfigInput) SetMcpToolDescription(mcpToolDescription string) {
	m.McpToolDescription = mcpToolDescription
}

func (m *McpToolConfigInput) SetMcpServerId(mcpServerId string) {
	m.McpServerId = mcpServerId
}

func (m *McpToolConfigInput) SetMcpServerName(mcpServerName string) {
	m.McpServerName = mcpServerName
}

func (m *McpToolConfigInput) SetMcpInputSchema(mcpInputSchema map[string]any) {
	m.McpInputSchema = mcpInputSchema
}

func (m *McpToolConfigInput) SetExecutionMode(executionMode *ToolExecutionMode) {
	m.ExecutionMode = executionMode
}

func (m *McpToolConfigInput) SetInputOverrides(inputOverrides map[string]*McpToolConfigInputInputOverridesValue) {
	m.InputOverrides = inputOverrides
}

func (m *McpToolConfigInput) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpToolConfigInput) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolConfigInput to string"
	}
	return string(jsonData)
}

func (m *McpToolConfigInput) UnmarshalJSON(data []byte) error {
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

func (m *McpToolConfigInput) MarshalJSON() ([]byte, error) {
	type embed McpToolConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
