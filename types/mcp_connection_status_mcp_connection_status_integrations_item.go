package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Status of a single MCP integration/server.
type McpConnectionStatusMcpConnectionStatusIntegrationsItem struct {
	IntegrationId   string                                                                `json:"integration_id" xml:"integration_id" required:"true"`
	IntegrationType McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationType `json:"integration_type" xml:"integration_type" required:"true"`
	IsConnected     bool                                                                  `json:"is_connected" xml:"is_connected" required:"true"`
	ToolCount       *int                                                                  `json:"tool_count,omitempty" xml:"tool_count,omitempty"`
	extraProperties map[string]interface{}
}

func (m *McpConnectionStatusMcpConnectionStatusIntegrationsItem) GetIntegrationId() (value string) {
	if m == nil {
		return
	}
	return m.IntegrationId
}

func (m *McpConnectionStatusMcpConnectionStatusIntegrationsItem) GetIntegrationType() (value McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationType) {
	if m == nil {
		return
	}
	return m.IntegrationType
}

func (m *McpConnectionStatusMcpConnectionStatusIntegrationsItem) GetIsConnected() (value bool) {
	if m == nil {
		return
	}
	return m.IsConnected
}

func (m *McpConnectionStatusMcpConnectionStatusIntegrationsItem) GetToolCount() (value *int) {
	if m == nil {
		return
	}
	return m.ToolCount
}

func (m *McpConnectionStatusMcpConnectionStatusIntegrationsItem) SetIntegrationId(integrationId string) {
	m.IntegrationId = integrationId
}

func (m *McpConnectionStatusMcpConnectionStatusIntegrationsItem) SetIntegrationType(integrationType McpConnectionStatusMcpConnectionStatusIntegrationsItemIntegrationType) {
	m.IntegrationType = integrationType
}

func (m *McpConnectionStatusMcpConnectionStatusIntegrationsItem) SetIsConnected(isConnected bool) {
	m.IsConnected = isConnected
}

func (m *McpConnectionStatusMcpConnectionStatusIntegrationsItem) SetToolCount(toolCount *int) {
	m.ToolCount = toolCount
}

func (m *McpConnectionStatusMcpConnectionStatusIntegrationsItem) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpConnectionStatusMcpConnectionStatusIntegrationsItem) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpConnectionStatusMcpConnectionStatusIntegrationsItem to string"
	}
	return string(jsonData)
}

func (m *McpConnectionStatusMcpConnectionStatusIntegrationsItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias McpConnectionStatusMcpConnectionStatusIntegrationsItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = McpConnectionStatusMcpConnectionStatusIntegrationsItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpConnectionStatusMcpConnectionStatusIntegrationsItem) MarshalJSON() ([]byte, error) {
	type embed McpConnectionStatusMcpConnectionStatusIntegrationsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
