package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Request model for updating MCP Server approval mode.
type McpApprovalPolicyUpdateRequestModel struct {
	// Defines the MCP server-level approval policy for tool execution.
	ApprovalPolicy  McpApprovalPolicy `json:"approval_policy" xml:"approval_policy" required:"true"`
	extraProperties map[string]interface{}
}

func (m *McpApprovalPolicyUpdateRequestModel) GetApprovalPolicy() (value McpApprovalPolicy) {
	if m == nil {
		return
	}
	return m.ApprovalPolicy
}

func (m *McpApprovalPolicyUpdateRequestModel) SetApprovalPolicy(approvalPolicy McpApprovalPolicy) {
	m.ApprovalPolicy = approvalPolicy
}

func (m *McpApprovalPolicyUpdateRequestModel) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpApprovalPolicyUpdateRequestModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpApprovalPolicyUpdateRequestModel to string"
	}
	return string(jsonData)
}

func (m *McpApprovalPolicyUpdateRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias McpApprovalPolicyUpdateRequestModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = McpApprovalPolicyUpdateRequestModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpApprovalPolicyUpdateRequestModel) MarshalJSON() ([]byte, error) {
	type embed McpApprovalPolicyUpdateRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
