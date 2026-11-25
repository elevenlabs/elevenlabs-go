package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentResponseMetadataAgentResponseMetadataEvent struct {
	Metadata        map[string]any `json:"metadata" xml:"metadata" required:"true"`
	EventId         int            `json:"event_id" xml:"event_id" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentResponseMetadataAgentResponseMetadataEvent) GetMetadata() (value map[string]any) {
	if a == nil {
		return
	}
	return a.Metadata
}

func (a *AgentResponseMetadataAgentResponseMetadataEvent) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentResponseMetadataAgentResponseMetadataEvent) SetMetadata(metadata map[string]any) {
	a.Metadata = metadata
}

func (a *AgentResponseMetadataAgentResponseMetadataEvent) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentResponseMetadataAgentResponseMetadataEvent) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentResponseMetadataAgentResponseMetadataEvent) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentResponseMetadataAgentResponseMetadataEvent to string"
	}
	return string(jsonData)
}

func (a *AgentResponseMetadataAgentResponseMetadataEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentResponseMetadataAgentResponseMetadataEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentResponseMetadataAgentResponseMetadataEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentResponseMetadataAgentResponseMetadataEvent) MarshalJSON() ([]byte, error) {
	type embed AgentResponseMetadataAgentResponseMetadataEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
