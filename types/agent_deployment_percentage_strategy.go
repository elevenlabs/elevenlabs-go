package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentDeploymentPercentageStrategy struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// Traffic percentage to deploy
	TrafficPercentage float64 `json:"traffic_percentage" xml:"traffic_percentage" required:"true" min:"0" max:"100"`
	extraProperties   map[string]interface{}
}

func (a *AgentDeploymentPercentageStrategy) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentDeploymentPercentageStrategy) GetTrafficPercentage() (value float64) {
	if a == nil {
		return
	}
	return a.TrafficPercentage
}

func (a *AgentDeploymentPercentageStrategy) SetType(type_ *string) {
	a.Type = type_
}

func (a *AgentDeploymentPercentageStrategy) SetTrafficPercentage(trafficPercentage float64) {
	a.TrafficPercentage = trafficPercentage
}

func (a *AgentDeploymentPercentageStrategy) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentDeploymentPercentageStrategy) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentDeploymentPercentageStrategy to string"
	}
	return string(jsonData)
}

func (a *AgentDeploymentPercentageStrategy) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentDeploymentPercentageStrategy
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentDeploymentPercentageStrategy(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentDeploymentPercentageStrategy) MarshalJSON() ([]byte, error) {
	type embed AgentDeploymentPercentageStrategy
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
