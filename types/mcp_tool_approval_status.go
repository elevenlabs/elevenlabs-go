package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Derived approval state for a currently discovered MCP tool.
type McpToolApprovalStatus struct {
	// Canonical MCP tool identifier in the form mcp:<server_id>:<tool_name>
	ToolId string `json:"tool_id" xml:"tool_id" required:"true"`
	// Whether a discovered MCP tool has an approval and whether it still matches.
	State McpToolApprovalState `json:"state" xml:"state" required:"true"`
	// Stored execution policy. Set when the tool has an approval.
	ApprovalPolicy *McpToolApprovalPolicy `json:"approval_policy,omitempty" xml:"approval_policy,omitempty"`
	// Previously approved definition, included when the tool needs review and a snapshot exists.
	ApprovedDefinition *McpApprovedToolDefinition `json:"approved_definition,omitempty" xml:"approved_definition,omitempty"`
	extraProperties    map[string]interface{}
}

func (m *McpToolApprovalStatus) GetToolId() (value string) {
	if m == nil {
		return
	}
	return m.ToolId
}

func (m *McpToolApprovalStatus) GetState() (value McpToolApprovalState) {
	if m == nil {
		return
	}
	return m.State
}

func (m *McpToolApprovalStatus) GetApprovalPolicy() (value *McpToolApprovalPolicy) {
	if m == nil {
		return
	}
	return m.ApprovalPolicy
}

func (m *McpToolApprovalStatus) GetApprovedDefinition() (value *McpApprovedToolDefinition) {
	if m == nil {
		return
	}
	return m.ApprovedDefinition
}

func (m *McpToolApprovalStatus) SetToolId(toolId string) {
	m.ToolId = toolId
}

func (m *McpToolApprovalStatus) SetState(state McpToolApprovalState) {
	m.State = state
}

func (m *McpToolApprovalStatus) SetApprovalPolicy(approvalPolicy *McpToolApprovalPolicy) {
	m.ApprovalPolicy = approvalPolicy
}

func (m *McpToolApprovalStatus) SetApprovedDefinition(approvedDefinition *McpApprovedToolDefinition) {
	m.ApprovedDefinition = approvedDefinition
}

func (m *McpToolApprovalStatus) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpToolApprovalStatus) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolApprovalStatus to string"
	}
	return string(jsonData)
}

func (m *McpToolApprovalStatus) UnmarshalJSON(data []byte) error {
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

func (m *McpToolApprovalStatus) MarshalJSON() ([]byte, error) {
	type embed McpToolApprovalStatus
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
