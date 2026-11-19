package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentToolResponseStreamedEvent struct {
	AgentToolResponse *AgentToolResponseClientEventData `json:"agent_tool_response" xml:"agent_tool_response" required:"true"`
	extraProperties   map[string]interface{}
}

func (a *AgentToolResponseStreamedEvent) GetAgentToolResponse() (value *AgentToolResponseClientEventData) {
	if a == nil {
		return
	}
	return a.AgentToolResponse
}

func (a *AgentToolResponseStreamedEvent) SetAgentToolResponse(agentToolResponse *AgentToolResponseClientEventData) {
	a.AgentToolResponse = agentToolResponse
}

func (a *AgentToolResponseStreamedEvent) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentToolResponseStreamedEvent) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentToolResponseStreamedEvent to string"
	}
	return string(jsonData)
}

func (a *AgentToolResponseStreamedEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentToolResponseStreamedEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentToolResponseStreamedEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentToolResponseStreamedEvent) MarshalJSON() ([]byte, error) {
	type embed AgentToolResponseStreamedEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
