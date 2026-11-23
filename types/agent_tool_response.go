package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentToolResponse struct {
	AgentToolResponse *AgentToolResponseAgentToolResponse `json:"agent_tool_response" xml:"agent_tool_response" required:"true"`
	Type              *string                             `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties   map[string]interface{}
}

func (a *AgentToolResponse) GetAgentToolResponse() (value *AgentToolResponseAgentToolResponse) {
	if a == nil {
		return
	}
	return a.AgentToolResponse
}

func (a *AgentToolResponse) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentToolResponse) SetAgentToolResponse(agentToolResponse *AgentToolResponseAgentToolResponse) {
	a.AgentToolResponse = agentToolResponse
}

func (a *AgentToolResponse) SetType(type_ *string) {
	a.Type = type_
}

func (a *AgentToolResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentToolResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentToolResponse to string"
	}
	return string(jsonData)
}

func (a *AgentToolResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentToolResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentToolResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentToolResponse) MarshalJSON() ([]byte, error) {
	type embed AgentToolResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
