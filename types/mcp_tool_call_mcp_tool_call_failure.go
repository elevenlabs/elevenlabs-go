package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type McpToolCallMcpToolCallFailure struct {
	ServiceId       string         `json:"service_id" xml:"service_id" required:"true"`
	ToolCallId      string         `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	ToolName        string         `json:"tool_name" xml:"tool_name" required:"true"`
	ToolDescription *string        `json:"tool_description,omitempty" xml:"tool_description,omitempty"`
	Parameters      map[string]any `json:"parameters" xml:"parameters" required:"true"`
	Timestamp       *string        `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	state           string
	ErrorMessage    string `json:"error_message" xml:"error_message" required:"true"`
	extraProperties map[string]interface{}
}

func (m *McpToolCallMcpToolCallFailure) GetServiceId() (value string) {
	if m == nil {
		return
	}
	return m.ServiceId
}

func (m *McpToolCallMcpToolCallFailure) GetToolCallId() (value string) {
	if m == nil {
		return
	}
	return m.ToolCallId
}

func (m *McpToolCallMcpToolCallFailure) GetToolName() (value string) {
	if m == nil {
		return
	}
	return m.ToolName
}

func (m *McpToolCallMcpToolCallFailure) GetToolDescription() (value *string) {
	if m == nil {
		return
	}
	return m.ToolDescription
}

func (m *McpToolCallMcpToolCallFailure) GetParameters() (value map[string]any) {
	if m == nil {
		return
	}
	return m.Parameters
}

func (m *McpToolCallMcpToolCallFailure) GetTimestamp() (value *string) {
	if m == nil {
		return
	}
	return m.Timestamp
}

func (m *McpToolCallMcpToolCallFailure) GetErrorMessage() (value string) {
	if m == nil {
		return
	}
	return m.ErrorMessage
}

func (m *McpToolCallMcpToolCallFailure) SetServiceId(serviceId string) {
	m.ServiceId = serviceId
}

func (m *McpToolCallMcpToolCallFailure) SetToolCallId(toolCallId string) {
	m.ToolCallId = toolCallId
}

func (m *McpToolCallMcpToolCallFailure) SetToolName(toolName string) {
	m.ToolName = toolName
}

func (m *McpToolCallMcpToolCallFailure) SetToolDescription(toolDescription *string) {
	m.ToolDescription = toolDescription
}

func (m *McpToolCallMcpToolCallFailure) SetParameters(parameters map[string]any) {
	m.Parameters = parameters
}

func (m *McpToolCallMcpToolCallFailure) SetTimestamp(timestamp *string) {
	m.Timestamp = timestamp
}

func (m *McpToolCallMcpToolCallFailure) SetErrorMessage(errorMessage string) {
	m.ErrorMessage = errorMessage
}

func (m *McpToolCallMcpToolCallFailure) State() string {
	return m.state
}

func (m *McpToolCallMcpToolCallFailure) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpToolCallMcpToolCallFailure) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolCallMcpToolCallFailure to string"
	}
	return string(jsonData)
}

func (m *McpToolCallMcpToolCallFailure) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	var literals struct {
		State *string `json:"state"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.State != nil && *literals.State != "failure" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", m, "failure", *literals.State)
	}
	m.state = "failure"
	extraProperties, err := internal.ExtractExtraProperties(data, *m, "state")
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpToolCallMcpToolCallFailure) MarshalJSON() ([]byte, error) {
	type embed McpToolCallMcpToolCallFailure
	var marshaler = struct {
		embed
		State string `json:"state"`
	}{
		embed: embed(*m),
		State: "failure",
	}
	return json.Marshal(marshaler)
}
