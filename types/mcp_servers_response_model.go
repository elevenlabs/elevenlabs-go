package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Response model for a list of MCP Server configurations.
type McpServersResponseModel struct {
	McpServers      []*McpServerResponseModel `json:"mcp_servers" xml:"mcp_servers" required:"true"`
	extraProperties map[string]interface{}
}

func (m *McpServersResponseModel) GetMcpServers() (value []*McpServerResponseModel) {
	if m == nil {
		return
	}
	return m.McpServers
}

func (m *McpServersResponseModel) SetMcpServers(mcpServers []*McpServerResponseModel) {
	m.McpServers = mcpServers
}

func (m *McpServersResponseModel) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpServersResponseModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpServersResponseModel to string"
	}
	return string(jsonData)
}

func (m *McpServersResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler McpServersResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*m = McpServersResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpServersResponseModel) MarshalJSON() ([]byte, error) {
	type embed McpServersResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
