package mcpservers

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

// Request model for adding approval for a single MCP tool.
type McpToolAddApprovalRequestModel struct {
	// ID of the MCP Server.
	McpServerId string `json:"-" url:"-"`
	// The name of the MCP tool
	ToolName string `json:"tool_name" xml:"tool_name" required:"true"`
	// The description of the MCP tool
	ToolDescription string `json:"tool_description" xml:"tool_description" required:"true"`
	// The input schema of the MCP tool (the schema defined on the MCP server before ElevenLabs does any extra processing)
	InputSchema map[string]any `json:"input_schema,omitempty"`
	// Defines the tool-level approval policy.
	ApprovalPolicy *types.McpToolApprovalPolicy `json:"approval_policy,omitempty" xml:"approval_policy,omitempty"`
}

func (m *McpToolAddApprovalRequestModel) GetMcpServerId() (value string) {
	if m == nil {
		return
	}
	return m.McpServerId
}

func (m *McpToolAddApprovalRequestModel) GetToolName() (value string) {
	if m == nil {
		return
	}
	return m.ToolName
}

func (m *McpToolAddApprovalRequestModel) GetToolDescription() (value string) {
	if m == nil {
		return
	}
	return m.ToolDescription
}

func (m *McpToolAddApprovalRequestModel) GetInputSchema() (value map[string]any) {
	if m == nil {
		return
	}
	return m.InputSchema
}

func (m *McpToolAddApprovalRequestModel) GetApprovalPolicy() (value *types.McpToolApprovalPolicy) {
	if m == nil {
		return
	}
	return m.ApprovalPolicy
}

func (m *McpToolAddApprovalRequestModel) SetMcpServerId(mcpServerId string) {
	m.McpServerId = mcpServerId
}

func (m *McpToolAddApprovalRequestModel) SetToolName(toolName string) {
	m.ToolName = toolName
}

func (m *McpToolAddApprovalRequestModel) SetToolDescription(toolDescription string) {
	m.ToolDescription = toolDescription
}

func (m *McpToolAddApprovalRequestModel) SetInputSchema(inputSchema map[string]any) {
	m.InputSchema = inputSchema
}

func (m *McpToolAddApprovalRequestModel) SetApprovalPolicy(approvalPolicy *types.McpToolApprovalPolicy) {
	m.ApprovalPolicy = approvalPolicy
}

func (m *McpToolAddApprovalRequestModel) GetExtraProperties() map[string]interface{} {
	return nil
}

func (m McpToolAddApprovalRequestModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolAddApprovalRequestModel to string"
	}
	return string(jsonData)
}

func (m *McpToolAddApprovalRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias McpToolAddApprovalRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = McpToolAddApprovalRequestModel(tmp)
	return nil
}

func (m *McpToolAddApprovalRequestModel) MarshalJSON() ([]byte, error) {
	type embed McpToolAddApprovalRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
