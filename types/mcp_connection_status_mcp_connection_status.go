package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// MCP connection status with array of integration statuses.
type McpConnectionStatusMcpConnectionStatus struct {
	Integrations    []*McpConnectionStatusMcpConnectionStatusIntegrationsItem `json:"integrations,omitempty" xml:"integrations,omitempty"`
	extraProperties map[string]interface{}
}

func (m *McpConnectionStatusMcpConnectionStatus) GetIntegrations() (value []*McpConnectionStatusMcpConnectionStatusIntegrationsItem) {
	if m == nil {
		return
	}
	return m.Integrations
}

func (m *McpConnectionStatusMcpConnectionStatus) SetIntegrations(integrations []*McpConnectionStatusMcpConnectionStatusIntegrationsItem) {
	m.Integrations = integrations
}

func (m *McpConnectionStatusMcpConnectionStatus) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpConnectionStatusMcpConnectionStatus) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpConnectionStatusMcpConnectionStatus to string"
	}
	return string(jsonData)
}

func (m *McpConnectionStatusMcpConnectionStatus) UnmarshalJSON(data []byte) error {
	type unmarshaler McpConnectionStatusMcpConnectionStatus
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = McpConnectionStatusMcpConnectionStatus(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpConnectionStatusMcpConnectionStatus) MarshalJSON() ([]byte, error) {
	type embed McpConnectionStatusMcpConnectionStatus
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
