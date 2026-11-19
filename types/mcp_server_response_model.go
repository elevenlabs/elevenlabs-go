package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model representing an MCP Server configuration.
type McpServerResponseModel struct {
	Id     string                 `json:"id" xml:"id" required:"true"`
	Config *McpServerConfigOutput `json:"config" xml:"config" required:"true"`
	// The access information of the MCP Server
	AccessInfo *ResourceAccessInfo `json:"access_info,omitempty" xml:"access_info,omitempty"`
	// List of agents that depend on this MCP Server.
	DependentAgents []*McpServerResponseModelDependentAgentsItem `json:"dependent_agents,omitempty" xml:"dependent_agents,omitempty"`
	Metadata        *McpServerMetadataResponseModel              `json:"metadata" xml:"metadata" required:"true"`
	extraProperties map[string]interface{}
}

func (m *McpServerResponseModel) GetId() (value string) {
	if m == nil {
		return
	}
	return m.Id
}

func (m *McpServerResponseModel) GetConfig() (value *McpServerConfigOutput) {
	if m == nil {
		return
	}
	return m.Config
}

func (m *McpServerResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if m == nil {
		return
	}
	return m.AccessInfo
}

func (m *McpServerResponseModel) GetDependentAgents() (value []*McpServerResponseModelDependentAgentsItem) {
	if m == nil {
		return
	}
	return m.DependentAgents
}

func (m *McpServerResponseModel) GetMetadata() (value *McpServerMetadataResponseModel) {
	if m == nil {
		return
	}
	return m.Metadata
}

func (m *McpServerResponseModel) SetId(id string) {
	m.Id = id
}

func (m *McpServerResponseModel) SetConfig(config *McpServerConfigOutput) {
	m.Config = config
}

func (m *McpServerResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	m.AccessInfo = accessInfo
}

func (m *McpServerResponseModel) SetDependentAgents(dependentAgents []*McpServerResponseModelDependentAgentsItem) {
	m.DependentAgents = dependentAgents
}

func (m *McpServerResponseModel) SetMetadata(metadata *McpServerMetadataResponseModel) {
	m.Metadata = metadata
}

func (m *McpServerResponseModel) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpServerResponseModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpServerResponseModel to string"
	}
	return string(jsonData)
}

func (m *McpServerResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpServerResponseModel) MarshalJSON() ([]byte, error) {
	type embed McpServerResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
