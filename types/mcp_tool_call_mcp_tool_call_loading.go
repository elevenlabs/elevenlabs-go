package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type McpToolCallMcpToolCallLoading struct {
	ServiceId       string         `json:"service_id" xml:"service_id" required:"true"`
	ToolCallId      string         `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	ToolName        string         `json:"tool_name" xml:"tool_name" required:"true"`
	ToolDescription *string        `json:"tool_description,omitempty" xml:"tool_description,omitempty"`
	Parameters      map[string]any `json:"parameters" xml:"parameters" required:"true"`
	Timestamp       *string        `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	state           string
	extraProperties map[string]interface{}
}

func (m *McpToolCallMcpToolCallLoading) GetServiceId() (value string) {
	if m == nil {
		return
	}
	return m.ServiceId
}

func (m *McpToolCallMcpToolCallLoading) GetToolCallId() (value string) {
	if m == nil {
		return
	}
	return m.ToolCallId
}

func (m *McpToolCallMcpToolCallLoading) GetToolName() (value string) {
	if m == nil {
		return
	}
	return m.ToolName
}

func (m *McpToolCallMcpToolCallLoading) GetToolDescription() (value *string) {
	if m == nil {
		return
	}
	return m.ToolDescription
}

func (m *McpToolCallMcpToolCallLoading) GetParameters() (value map[string]any) {
	if m == nil {
		return
	}
	return m.Parameters
}

func (m *McpToolCallMcpToolCallLoading) GetTimestamp() (value *string) {
	if m == nil {
		return
	}
	return m.Timestamp
}

func (m *McpToolCallMcpToolCallLoading) SetServiceId(serviceId string) {
	m.ServiceId = serviceId
}

func (m *McpToolCallMcpToolCallLoading) SetToolCallId(toolCallId string) {
	m.ToolCallId = toolCallId
}

func (m *McpToolCallMcpToolCallLoading) SetToolName(toolName string) {
	m.ToolName = toolName
}

func (m *McpToolCallMcpToolCallLoading) SetToolDescription(toolDescription *string) {
	m.ToolDescription = toolDescription
}

func (m *McpToolCallMcpToolCallLoading) SetParameters(parameters map[string]any) {
	m.Parameters = parameters
}

func (m *McpToolCallMcpToolCallLoading) SetTimestamp(timestamp *string) {
	m.Timestamp = timestamp
}

func (m *McpToolCallMcpToolCallLoading) State() string {
	return m.state
}

func (m *McpToolCallMcpToolCallLoading) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpToolCallMcpToolCallLoading) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolCallMcpToolCallLoading to string"
	}
	return string(jsonData)
}

func (m *McpToolCallMcpToolCallLoading) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	var literals struct {
		State *string `json:"state"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.State != nil && *literals.State != "loading" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", m, "loading", *literals.State)
	}
	m.state = "loading"
	extraProperties, err := internal.ExtractExtraProperties(data, *m, "state")
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpToolCallMcpToolCallLoading) MarshalJSON() ([]byte, error) {
	type embed McpToolCallMcpToolCallLoading
	var marshaler = struct {
		embed
		State string `json:"state"`
	}{
		embed: embed(*m),
		State: "loading",
	}
	return json.Marshal(marshaler)
}
