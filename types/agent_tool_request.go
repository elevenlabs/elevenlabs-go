package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentToolRequest struct {
	AgentToolRequest *AgentToolRequestAgentToolRequest `json:"agent_tool_request" xml:"agent_tool_request" required:"true"`
	Type             *string                           `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties  map[string]interface{}
}

func (a *AgentToolRequest) GetAgentToolRequest() (value *AgentToolRequestAgentToolRequest) {
	if a == nil {
		return
	}
	return a.AgentToolRequest
}

func (a *AgentToolRequest) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentToolRequest) SetAgentToolRequest(agentToolRequest *AgentToolRequestAgentToolRequest) {
	a.AgentToolRequest = agentToolRequest
}

func (a *AgentToolRequest) SetType(type_ *string) {
	a.Type = type_
}

func (a *AgentToolRequest) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentToolRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentToolRequest to string"
	}
	return string(jsonData)
}

func (a *AgentToolRequest) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentToolRequest
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentToolRequest(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentToolRequest) MarshalJSON() ([]byte, error) {
	type embed AgentToolRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
