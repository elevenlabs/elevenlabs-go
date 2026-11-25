package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type McpToolCallMcpToolCallAwaitingApproval struct {
	ServiceId       string         `json:"service_id" xml:"service_id" required:"true"`
	ToolCallId      string         `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	ToolName        string         `json:"tool_name" xml:"tool_name" required:"true"`
	ToolDescription *string        `json:"tool_description,omitempty" xml:"tool_description,omitempty"`
	Parameters      map[string]any `json:"parameters" xml:"parameters" required:"true"`
	Timestamp       *string        `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	state           string
	// Timeout in seconds for user approval
	ApprovalTimeoutSecs *int `json:"approval_timeout_secs,omitempty" xml:"approval_timeout_secs,omitempty"`
	extraProperties     map[string]interface{}
}

func (m *McpToolCallMcpToolCallAwaitingApproval) GetServiceId() (value string) {
	if m == nil {
		return
	}
	return m.ServiceId
}

func (m *McpToolCallMcpToolCallAwaitingApproval) GetToolCallId() (value string) {
	if m == nil {
		return
	}
	return m.ToolCallId
}

func (m *McpToolCallMcpToolCallAwaitingApproval) GetToolName() (value string) {
	if m == nil {
		return
	}
	return m.ToolName
}

func (m *McpToolCallMcpToolCallAwaitingApproval) GetToolDescription() (value *string) {
	if m == nil {
		return
	}
	return m.ToolDescription
}

func (m *McpToolCallMcpToolCallAwaitingApproval) GetParameters() (value map[string]any) {
	if m == nil {
		return
	}
	return m.Parameters
}

func (m *McpToolCallMcpToolCallAwaitingApproval) GetTimestamp() (value *string) {
	if m == nil {
		return
	}
	return m.Timestamp
}

func (m *McpToolCallMcpToolCallAwaitingApproval) GetApprovalTimeoutSecs() (value *int) {
	if m == nil {
		return
	}
	return m.ApprovalTimeoutSecs
}

func (m *McpToolCallMcpToolCallAwaitingApproval) SetServiceId(serviceId string) {
	m.ServiceId = serviceId
}

func (m *McpToolCallMcpToolCallAwaitingApproval) SetToolCallId(toolCallId string) {
	m.ToolCallId = toolCallId
}

func (m *McpToolCallMcpToolCallAwaitingApproval) SetToolName(toolName string) {
	m.ToolName = toolName
}

func (m *McpToolCallMcpToolCallAwaitingApproval) SetToolDescription(toolDescription *string) {
	m.ToolDescription = toolDescription
}

func (m *McpToolCallMcpToolCallAwaitingApproval) SetParameters(parameters map[string]any) {
	m.Parameters = parameters
}

func (m *McpToolCallMcpToolCallAwaitingApproval) SetTimestamp(timestamp *string) {
	m.Timestamp = timestamp
}

func (m *McpToolCallMcpToolCallAwaitingApproval) SetApprovalTimeoutSecs(approvalTimeoutSecs *int) {
	m.ApprovalTimeoutSecs = approvalTimeoutSecs
}

func (m *McpToolCallMcpToolCallAwaitingApproval) State() string {
	return m.state
}

func (m *McpToolCallMcpToolCallAwaitingApproval) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpToolCallMcpToolCallAwaitingApproval) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolCallMcpToolCallAwaitingApproval to string"
	}
	return string(jsonData)
}

func (m *McpToolCallMcpToolCallAwaitingApproval) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	var literals struct {
		State *string `json:"state"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.State != nil && *literals.State != "awaiting_approval" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", m, "awaiting_approval", *literals.State)
	}
	m.state = "awaiting_approval"
	extraProperties, err := internal.ExtractExtraProperties(data, *m, "state")
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpToolCallMcpToolCallAwaitingApproval) MarshalJSON() ([]byte, error) {
	type embed McpToolCallMcpToolCallAwaitingApproval
	var marshaler = struct {
		embed
		State string `json:"state"`
	}{
		embed: embed(*m),
		State: "awaiting_approval",
	}
	return json.Marshal(marshaler)
}
