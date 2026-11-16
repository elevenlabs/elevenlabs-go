package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentDeploymentResponse struct {
	// Map of branch IDs to traffic percentages
	TrafficPercentageBranchIdMap map[string]float64 `json:"traffic_percentage_branch_id_map,omitempty"`
	extraProperties              map[string]interface{}
}

func (a *AgentDeploymentResponse) GetTrafficPercentageBranchIdMap() (value map[string]float64) {
	if a == nil {
		return
	}
	return a.TrafficPercentageBranchIdMap
}

func (a *AgentDeploymentResponse) SetTrafficPercentageBranchIdMap(trafficPercentageBranchIdMap map[string]float64) {
	a.TrafficPercentageBranchIdMap = trafficPercentageBranchIdMap
}

func (a *AgentDeploymentResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentDeploymentResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentDeploymentResponse to string"
	}
	return string(jsonData)
}

func (a *AgentDeploymentResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler AgentDeploymentResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AgentDeploymentResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentDeploymentResponse) MarshalJSON() ([]byte, error) {
	type embed AgentDeploymentResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
