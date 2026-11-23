package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentResponseCompleteAgentResponseCompleteEvent struct {
	EventId         int `json:"event_id" xml:"event_id" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentResponseCompleteAgentResponseCompleteEvent) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentResponseCompleteAgentResponseCompleteEvent) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentResponseCompleteAgentResponseCompleteEvent) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentResponseCompleteAgentResponseCompleteEvent) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentResponseCompleteAgentResponseCompleteEvent to string"
	}
	return string(jsonData)
}

func (a *AgentResponseCompleteAgentResponseCompleteEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentResponseCompleteAgentResponseCompleteEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentResponseCompleteAgentResponseCompleteEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentResponseCompleteAgentResponseCompleteEvent) MarshalJSON() ([]byte, error) {
	type embed AgentResponseCompleteAgentResponseCompleteEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
