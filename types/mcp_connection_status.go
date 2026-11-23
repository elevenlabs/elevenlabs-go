package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type McpConnectionStatus struct {
	// MCP connection status with array of integration statuses.
	McpConnectionStatus *McpConnectionStatusMcpConnectionStatus `json:"mcp_connection_status" xml:"mcp_connection_status" required:"true"`
	Type                *string                                 `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties     map[string]interface{}
}

func (m *McpConnectionStatus) GetMcpConnectionStatus() (value *McpConnectionStatusMcpConnectionStatus) {
	if m == nil {
		return
	}
	return m.McpConnectionStatus
}

func (m *McpConnectionStatus) GetType() (value *string) {
	if m == nil {
		return
	}
	return m.Type
}

func (m *McpConnectionStatus) SetMcpConnectionStatus(mcpConnectionStatus *McpConnectionStatusMcpConnectionStatus) {
	m.McpConnectionStatus = mcpConnectionStatus
}

func (m *McpConnectionStatus) SetType(type_ *string) {
	m.Type = type_
}

func (m *McpConnectionStatus) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpConnectionStatus) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpConnectionStatus to string"
	}
	return string(jsonData)
}

func (m *McpConnectionStatus) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias McpConnectionStatus
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = McpConnectionStatus(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpConnectionStatus) MarshalJSON() ([]byte, error) {
	type embed McpConnectionStatus
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
