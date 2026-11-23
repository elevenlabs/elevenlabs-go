package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type McpToolCall struct {
	McpToolCall     *McpToolCallMcpToolCall `json:"mcp_tool_call" xml:"mcp_tool_call" required:"true"`
	Type            *string                 `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (m *McpToolCall) GetMcpToolCall() (value *McpToolCallMcpToolCall) {
	if m == nil {
		return
	}
	return m.McpToolCall
}

func (m *McpToolCall) GetType() (value *string) {
	if m == nil {
		return
	}
	return m.Type
}

func (m *McpToolCall) SetMcpToolCall(mcpToolCall *McpToolCallMcpToolCall) {
	m.McpToolCall = mcpToolCall
}

func (m *McpToolCall) SetType(type_ *string) {
	m.Type = type_
}

func (m *McpToolCall) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpToolCall) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolCall to string"
	}
	return string(jsonData)
}

func (m *McpToolCall) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias McpToolCall
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = McpToolCall(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpToolCall) MarshalJSON() ([]byte, error) {
	type embed McpToolCall
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
