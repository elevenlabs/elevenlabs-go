package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentResponseClientEvent struct {
	AgentResponseEvent *AgentResponseClientEventData `json:"agent_response_event" xml:"agent_response_event" required:"true"`
	extraProperties    map[string]interface{}
}

func (a *AgentResponseClientEvent) GetAgentResponseEvent() (value *AgentResponseClientEventData) {
	if a == nil {
		return
	}
	return a.AgentResponseEvent
}

func (a *AgentResponseClientEvent) SetAgentResponseEvent(agentResponseEvent *AgentResponseClientEventData) {
	a.AgentResponseEvent = agentResponseEvent
}

func (a *AgentResponseClientEvent) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentResponseClientEvent) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentResponseClientEvent to string"
	}
	return string(jsonData)
}

func (a *AgentResponseClientEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentResponseClientEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentResponseClientEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentResponseClientEvent) MarshalJSON() ([]byte, error) {
	type embed AgentResponseClientEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
