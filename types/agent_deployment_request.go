package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AgentDeploymentRequest struct {
	// List of deployment requests
	Requests        []*AgentDeploymentRequestItem `json:"requests" xml:"requests" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentDeploymentRequest) GetRequests() (value []*AgentDeploymentRequestItem) {
	if a == nil {
		return
	}
	return a.Requests
}

func (a *AgentDeploymentRequest) SetRequests(requests []*AgentDeploymentRequestItem) {
	a.Requests = requests
}

func (a *AgentDeploymentRequest) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentDeploymentRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentDeploymentRequest to string"
	}
	return string(jsonData)
}

func (a *AgentDeploymentRequest) UnmarshalJSON(data []byte) error {
	type unmarshaler AgentDeploymentRequest
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AgentDeploymentRequest(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentDeploymentRequest) MarshalJSON() ([]byte, error) {
	type embed AgentDeploymentRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
