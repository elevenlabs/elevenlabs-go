package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// An MCP tool configuration that can be used to call MCP servers
type McpToolConfigOutput struct {
	Name string `json:"name" xml:"name" required:"true" pattern:"^[a-zA-Z0-9_-]{1,64}$"`
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
	Parameters *ObjectJsonSchemaPropertyOutput `json:"parameters,omitempty" xml:"parameters,omitempty"`
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
	InputOverrides  map[string]*McpToolConfigOutputInputOverridesValue `json:"input_overrides,omitempty"`
	extraProperties map[string]interface{}
}

func (m *McpToolConfigOutput) GetName() (value string) {
	if m == nil {
		return
	}
	return m.Name
}

func (m *McpToolConfigOutput) GetDescription() (value string) {
	if m == nil {
		return
	}
	return m.Description
}

func (m *McpToolConfigOutput) GetResponseTimeoutSecs() (value *int) {
	if m == nil {
		return
	}
	return m.ResponseTimeoutSecs
}

func (m *McpToolConfigOutput) GetDisableInterruptions() (value *bool) {
	if m == nil {
		return
	}
	return m.DisableInterruptions
}

func (m *McpToolConfigOutput) GetInterruptionMode() (value *ToolInterruptionMode) {
	if m == nil {
		return
	}
	return m.InterruptionMode
}

func (m *McpToolConfigOutput) GetForcePreToolSpeech() (value *bool) {
	if m == nil {
		return
	}
	return m.ForcePreToolSpeech
}

func (m *McpToolConfigOutput) GetPreToolSpeech() (value *PreToolSpeechMode) {
	if m == nil {
		return
	}
	return m.PreToolSpeech
}

func (m *McpToolConfigOutput) GetAssignments() (value []*DynamicVariableAssignment) {
	if m == nil {
		return
	}
	return m.Assignments
}

func (m *McpToolConfigOutput) GetToolCallSound() (value *ToolCallSoundType) {
	if m == nil {
		return
	}
	return m.ToolCallSound
}

func (m *McpToolConfigOutput) GetToolCallSoundBehavior() (value *ToolCallSoundBehavior) {
	if m == nil {
		return
	}
	return m.ToolCallSoundBehavior
}

func (m *McpToolConfigOutput) GetToolErrorHandlingMode() (value *ToolErrorHandlingMode) {
	if m == nil {
		return
	}
	return m.ToolErrorHandlingMode
}

func (m *McpToolConfigOutput) GetIntegrationType() (value IntegrationType) {
	if m == nil {
		return
	}
	return m.IntegrationType
}

func (m *McpToolConfigOutput) GetParameters() (value *ObjectJsonSchemaPropertyOutput) {
	if m == nil {
		return
	}
	return m.Parameters
}

func (m *McpToolConfigOutput) GetApprovalPolicy() (value *McpApprovalPolicy) {
	if m == nil {
		return
	}
	return m.ApprovalPolicy
}

func (m *McpToolConfigOutput) GetMcpToolName() (value string) {
	if m == nil {
		return
	}
	return m.McpToolName
}

func (m *McpToolConfigOutput) GetMcpToolDescription() (value string) {
	if m == nil {
		return
	}
	return m.McpToolDescription
}

func (m *McpToolConfigOutput) GetMcpServerId() (value string) {
	if m == nil {
		return
	}
	return m.McpServerId
}

func (m *McpToolConfigOutput) GetMcpServerName() (value string) {
	if m == nil {
		return
	}
	return m.McpServerName
}

func (m *McpToolConfigOutput) GetMcpInputSchema() (value map[string]any) {
	if m == nil {
		return
	}
	return m.McpInputSchema
}

func (m *McpToolConfigOutput) GetExecutionMode() (value *ToolExecutionMode) {
	if m == nil {
		return
	}
	return m.ExecutionMode
}

func (m *McpToolConfigOutput) GetInputOverrides() (value map[string]*McpToolConfigOutputInputOverridesValue) {
	if m == nil {
		return
	}
	return m.InputOverrides
}

func (m *McpToolConfigOutput) SetName(name string) {
	m.Name = name
}

func (m *McpToolConfigOutput) SetDescription(description string) {
	m.Description = description
}

func (m *McpToolConfigOutput) SetResponseTimeoutSecs(responseTimeoutSecs *int) {
	m.ResponseTimeoutSecs = responseTimeoutSecs
}

func (m *McpToolConfigOutput) SetDisableInterruptions(disableInterruptions *bool) {
	m.DisableInterruptions = disableInterruptions
}

func (m *McpToolConfigOutput) SetInterruptionMode(interruptionMode *ToolInterruptionMode) {
	m.InterruptionMode = interruptionMode
}

func (m *McpToolConfigOutput) SetForcePreToolSpeech(forcePreToolSpeech *bool) {
	m.ForcePreToolSpeech = forcePreToolSpeech
}

func (m *McpToolConfigOutput) SetPreToolSpeech(preToolSpeech *PreToolSpeechMode) {
	m.PreToolSpeech = preToolSpeech
}

func (m *McpToolConfigOutput) SetAssignments(assignments []*DynamicVariableAssignment) {
	m.Assignments = assignments
}

func (m *McpToolConfigOutput) SetToolCallSound(toolCallSound *ToolCallSoundType) {
	m.ToolCallSound = toolCallSound
}

func (m *McpToolConfigOutput) SetToolCallSoundBehavior(toolCallSoundBehavior *ToolCallSoundBehavior) {
	m.ToolCallSoundBehavior = toolCallSoundBehavior
}

func (m *McpToolConfigOutput) SetToolErrorHandlingMode(toolErrorHandlingMode *ToolErrorHandlingMode) {
	m.ToolErrorHandlingMode = toolErrorHandlingMode
}

func (m *McpToolConfigOutput) SetIntegrationType(integrationType IntegrationType) {
	m.IntegrationType = integrationType
}

func (m *McpToolConfigOutput) SetParameters(parameters *ObjectJsonSchemaPropertyOutput) {
	m.Parameters = parameters
}

func (m *McpToolConfigOutput) SetApprovalPolicy(approvalPolicy *McpApprovalPolicy) {
	m.ApprovalPolicy = approvalPolicy
}

func (m *McpToolConfigOutput) SetMcpToolName(mcpToolName string) {
	m.McpToolName = mcpToolName
}

func (m *McpToolConfigOutput) SetMcpToolDescription(mcpToolDescription string) {
	m.McpToolDescription = mcpToolDescription
}

func (m *McpToolConfigOutput) SetMcpServerId(mcpServerId string) {
	m.McpServerId = mcpServerId
}

func (m *McpToolConfigOutput) SetMcpServerName(mcpServerName string) {
	m.McpServerName = mcpServerName
}

func (m *McpToolConfigOutput) SetMcpInputSchema(mcpInputSchema map[string]any) {
	m.McpInputSchema = mcpInputSchema
}

func (m *McpToolConfigOutput) SetExecutionMode(executionMode *ToolExecutionMode) {
	m.ExecutionMode = executionMode
}

func (m *McpToolConfigOutput) SetInputOverrides(inputOverrides map[string]*McpToolConfigOutputInputOverridesValue) {
	m.InputOverrides = inputOverrides
}

func (m *McpToolConfigOutput) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpToolConfigOutput) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolConfigOutput to string"
	}
	return string(jsonData)
}

func (m *McpToolConfigOutput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m, "type")
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpToolConfigOutput) MarshalJSON() ([]byte, error) {
	type embed McpToolConfigOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
