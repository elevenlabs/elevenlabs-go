package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type McpToolCallMcpToolCallSuccess struct {
	ServiceId       string         `json:"service_id" xml:"service_id" required:"true"`
	ToolCallId      string         `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	ToolName        string         `json:"tool_name" xml:"tool_name" required:"true"`
	ToolDescription *string        `json:"tool_description,omitempty" xml:"tool_description,omitempty"`
	Parameters      map[string]any `json:"parameters" xml:"parameters" required:"true"`
	Timestamp       *string        `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	state           string
	Result          []map[string]any `json:"result" xml:"result" required:"true"`
	extraProperties map[string]interface{}
}

func (m *McpToolCallMcpToolCallSuccess) GetServiceId() (value string) {
	if m == nil {
		return
	}
	return m.ServiceId
}

func (m *McpToolCallMcpToolCallSuccess) GetToolCallId() (value string) {
	if m == nil {
		return
	}
	return m.ToolCallId
}

func (m *McpToolCallMcpToolCallSuccess) GetToolName() (value string) {
	if m == nil {
		return
	}
	return m.ToolName
}

func (m *McpToolCallMcpToolCallSuccess) GetToolDescription() (value *string) {
	if m == nil {
		return
	}
	return m.ToolDescription
}

func (m *McpToolCallMcpToolCallSuccess) GetParameters() (value map[string]any) {
	if m == nil {
		return
	}
	return m.Parameters
}

func (m *McpToolCallMcpToolCallSuccess) GetTimestamp() (value *string) {
	if m == nil {
		return
	}
	return m.Timestamp
}

func (m *McpToolCallMcpToolCallSuccess) GetResult() (value []map[string]any) {
	if m == nil {
		return
	}
	return m.Result
}

func (m *McpToolCallMcpToolCallSuccess) SetServiceId(serviceId string) {
	m.ServiceId = serviceId
}

func (m *McpToolCallMcpToolCallSuccess) SetToolCallId(toolCallId string) {
	m.ToolCallId = toolCallId
}

func (m *McpToolCallMcpToolCallSuccess) SetToolName(toolName string) {
	m.ToolName = toolName
}

func (m *McpToolCallMcpToolCallSuccess) SetToolDescription(toolDescription *string) {
	m.ToolDescription = toolDescription
}

func (m *McpToolCallMcpToolCallSuccess) SetParameters(parameters map[string]any) {
	m.Parameters = parameters
}

func (m *McpToolCallMcpToolCallSuccess) SetTimestamp(timestamp *string) {
	m.Timestamp = timestamp
}

func (m *McpToolCallMcpToolCallSuccess) SetResult(result []map[string]any) {
	m.Result = result
}

func (m *McpToolCallMcpToolCallSuccess) State() string {
	return m.state
}

func (m *McpToolCallMcpToolCallSuccess) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpToolCallMcpToolCallSuccess) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolCallMcpToolCallSuccess to string"
	}
	return string(jsonData)
}

func (m *McpToolCallMcpToolCallSuccess) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	var literals struct {
		State *string `json:"state"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.State != nil && *literals.State != "success" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", m, "success", *literals.State)
	}
	m.state = "success"
	extraProperties, err := internal.ExtractExtraProperties(data, *m, "state")
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpToolCallMcpToolCallSuccess) MarshalJSON() ([]byte, error) {
	type embed McpToolCallMcpToolCallSuccess
	var marshaler = struct {
		embed
		State string `json:"state"`
	}{
		embed: embed(*m),
		State: "success",
	}
	return json.Marshal(marshaler)
}
