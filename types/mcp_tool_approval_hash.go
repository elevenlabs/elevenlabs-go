package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Model for storing tool approval hashes for per-tool approval.
type McpToolApprovalHash struct {
	// The name of the MCP tool
	ToolName string `json:"tool_name" xml:"tool_name" required:"true"`
	// SHA256 hash of the tool's parameters and description
	ToolHash string `json:"tool_hash" xml:"tool_hash" required:"true"`
	// Defines the tool-level approval policy.
	ApprovalPolicy  *McpToolApprovalPolicy `json:"approval_policy,omitempty" xml:"approval_policy,omitempty"`
	extraProperties map[string]interface{}
}

func (m *McpToolApprovalHash) GetToolName() (value string) {
	if m == nil {
		return
	}
	return m.ToolName
}

func (m *McpToolApprovalHash) GetToolHash() (value string) {
	if m == nil {
		return
	}
	return m.ToolHash
}

func (m *McpToolApprovalHash) GetApprovalPolicy() (value *McpToolApprovalPolicy) {
	if m == nil {
		return
	}
	return m.ApprovalPolicy
}

func (m *McpToolApprovalHash) SetToolName(toolName string) {
	m.ToolName = toolName
}

func (m *McpToolApprovalHash) SetToolHash(toolHash string) {
	m.ToolHash = toolHash
}

func (m *McpToolApprovalHash) SetApprovalPolicy(approvalPolicy *McpToolApprovalPolicy) {
	m.ApprovalPolicy = approvalPolicy
}

func (m *McpToolApprovalHash) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpToolApprovalHash) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolApprovalHash to string"
	}
	return string(jsonData)
}

func (m *McpToolApprovalHash) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias McpToolApprovalHash
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = McpToolApprovalHash(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpToolApprovalHash) MarshalJSON() ([]byte, error) {
	type embed McpToolApprovalHash
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
