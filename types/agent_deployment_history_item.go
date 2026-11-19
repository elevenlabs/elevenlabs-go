package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentDeploymentHistoryItem struct {
	// ID of the deployment
	Id string `json:"id" xml:"id" required:"true"`
	// Map of branch IDs to traffic percentages
	TrafficPercentageBranchIdMap map[string]float64 `json:"traffic_percentage_branch_id_map,omitempty"`
	// Unix timestamp of when the traffic split was applied
	DeployedAtUnixSecs int `json:"deployed_at_unix_secs" xml:"deployed_at_unix_secs" required:"true"`
	// Whether this is the live traffic split
	IsActive bool `json:"is_active" xml:"is_active" required:"true"`
	// What caused this traffic split change. Null for deployments created before this was recorded.
	Source *AgentDeploymentSource `json:"source,omitempty" xml:"source,omitempty"`
	// Access information for the deployment, including who created it
	AccessInfo      *ResourceAccessInfo `json:"access_info,omitempty" xml:"access_info,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentDeploymentHistoryItem) GetId() (value string) {
	if a == nil {
		return
	}
	return a.Id
}

func (a *AgentDeploymentHistoryItem) GetTrafficPercentageBranchIdMap() (value map[string]float64) {
	if a == nil {
		return
	}
	return a.TrafficPercentageBranchIdMap
}

func (a *AgentDeploymentHistoryItem) GetDeployedAtUnixSecs() (value int) {
	if a == nil {
		return
	}
	return a.DeployedAtUnixSecs
}

func (a *AgentDeploymentHistoryItem) GetIsActive() (value bool) {
	if a == nil {
		return
	}
	return a.IsActive
}

func (a *AgentDeploymentHistoryItem) GetSource() (value *AgentDeploymentSource) {
	if a == nil {
		return
	}
	return a.Source
}

func (a *AgentDeploymentHistoryItem) GetAccessInfo() (value *ResourceAccessInfo) {
	if a == nil {
		return
	}
	return a.AccessInfo
}

func (a *AgentDeploymentHistoryItem) SetId(id string) {
	a.Id = id
}

func (a *AgentDeploymentHistoryItem) SetTrafficPercentageBranchIdMap(trafficPercentageBranchIdMap map[string]float64) {
	a.TrafficPercentageBranchIdMap = trafficPercentageBranchIdMap
}

func (a *AgentDeploymentHistoryItem) SetDeployedAtUnixSecs(deployedAtUnixSecs int) {
	a.DeployedAtUnixSecs = deployedAtUnixSecs
}

func (a *AgentDeploymentHistoryItem) SetIsActive(isActive bool) {
	a.IsActive = isActive
}

func (a *AgentDeploymentHistoryItem) SetSource(source *AgentDeploymentSource) {
	a.Source = source
}

func (a *AgentDeploymentHistoryItem) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	a.AccessInfo = accessInfo
}

func (a *AgentDeploymentHistoryItem) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentDeploymentHistoryItem) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentDeploymentHistoryItem to string"
	}
	return string(jsonData)
}

func (a *AgentDeploymentHistoryItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentDeploymentHistoryItem) MarshalJSON() ([]byte, error) {
	type embed AgentDeploymentHistoryItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
