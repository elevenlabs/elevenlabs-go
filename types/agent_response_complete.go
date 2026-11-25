package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentResponseComplete struct {
	AgentResponseCompleteEvent *AgentResponseCompleteAgentResponseCompleteEvent `json:"agent_response_complete_event" xml:"agent_response_complete_event" required:"true"`
	Type                       *string                                          `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties            map[string]interface{}
}

func (a *AgentResponseComplete) GetAgentResponseCompleteEvent() (value *AgentResponseCompleteAgentResponseCompleteEvent) {
	if a == nil {
		return
	}
	return a.AgentResponseCompleteEvent
}

func (a *AgentResponseComplete) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentResponseComplete) SetAgentResponseCompleteEvent(agentResponseCompleteEvent *AgentResponseCompleteAgentResponseCompleteEvent) {
	a.AgentResponseCompleteEvent = agentResponseCompleteEvent
}

func (a *AgentResponseComplete) SetType(type_ *string) {
	a.Type = type_
}

func (a *AgentResponseComplete) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentResponseComplete) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentResponseComplete to string"
	}
	return string(jsonData)
}

func (a *AgentResponseComplete) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentResponseComplete
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentResponseComplete(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentResponseComplete) MarshalJSON() ([]byte, error) {
	type embed AgentResponseComplete
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
