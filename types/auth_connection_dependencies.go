package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Dependencies that use an auth connection
type AuthConnectionDependencies struct {
	Tools                  []*AuthConnectionDependenciesToolsItem      `json:"tools,omitempty" xml:"tools,omitempty"`
	McpServers             []*AuthConnectionDependenciesMcpServersItem `json:"mcp_servers,omitempty" xml:"mcp_servers,omitempty"`
	IntegrationConnections []*DependentIntegrationConnectionIdentifier `json:"integration_connections,omitempty" xml:"integration_connections,omitempty"`
	extraProperties        map[string]interface{}
}

func (a *AuthConnectionDependencies) GetTools() (value []*AuthConnectionDependenciesToolsItem) {
	if a == nil {
		return
	}
	return a.Tools
}

func (a *AuthConnectionDependencies) GetMcpServers() (value []*AuthConnectionDependenciesMcpServersItem) {
	if a == nil {
		return
	}
	return a.McpServers
}

func (a *AuthConnectionDependencies) GetIntegrationConnections() (value []*DependentIntegrationConnectionIdentifier) {
	if a == nil {
		return
	}
	return a.IntegrationConnections
}

func (a *AuthConnectionDependencies) SetTools(tools []*AuthConnectionDependenciesToolsItem) {
	a.Tools = tools
}

func (a *AuthConnectionDependencies) SetMcpServers(mcpServers []*AuthConnectionDependenciesMcpServersItem) {
	a.McpServers = mcpServers
}

func (a *AuthConnectionDependencies) SetIntegrationConnections(integrationConnections []*DependentIntegrationConnectionIdentifier) {
	a.IntegrationConnections = integrationConnections
}

func (a *AuthConnectionDependencies) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AuthConnectionDependencies) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AuthConnectionDependencies to string"
	}
	return string(jsonData)
}

func (a *AuthConnectionDependencies) UnmarshalJSON(data []byte) error {
	type unmarshaler AuthConnectionDependencies
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AuthConnectionDependencies(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AuthConnectionDependencies) MarshalJSON() ([]byte, error) {
	type embed AuthConnectionDependencies
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
