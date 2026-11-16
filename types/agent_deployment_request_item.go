package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentDeploymentRequestItem struct {
	// ID of the branch to deploy
	BranchId           string                             `json:"branch_id" xml:"branch_id" required:"true"`
	DeploymentStrategy *AgentDeploymentPercentageStrategy `json:"deployment_strategy" xml:"deployment_strategy" required:"true"`
	extraProperties    map[string]interface{}
}

func (a *AgentDeploymentRequestItem) GetBranchId() (value string) {
	if a == nil {
		return
	}
	return a.BranchId
}

func (a *AgentDeploymentRequestItem) GetDeploymentStrategy() (value *AgentDeploymentPercentageStrategy) {
	if a == nil {
		return
	}
	return a.DeploymentStrategy
}

func (a *AgentDeploymentRequestItem) SetBranchId(branchId string) {
	a.BranchId = branchId
}

func (a *AgentDeploymentRequestItem) SetDeploymentStrategy(deploymentStrategy *AgentDeploymentPercentageStrategy) {
	a.DeploymentStrategy = deploymentStrategy
}

func (a *AgentDeploymentRequestItem) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentDeploymentRequestItem) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentDeploymentRequestItem to string"
	}
	return string(jsonData)
}

func (a *AgentDeploymentRequestItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentDeploymentRequestItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentDeploymentRequestItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentDeploymentRequestItem) MarshalJSON() ([]byte, error) {
	type embed AgentDeploymentRequestItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
