package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentResponseMetadata struct {
	AgentResponseMetadataEvent *AgentResponseMetadataAgentResponseMetadataEvent `json:"agent_response_metadata_event" xml:"agent_response_metadata_event" required:"true"`
	Type                       *string                                          `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties            map[string]interface{}
}

func (a *AgentResponseMetadata) GetAgentResponseMetadataEvent() (value *AgentResponseMetadataAgentResponseMetadataEvent) {
	if a == nil {
		return
	}
	return a.AgentResponseMetadataEvent
}

func (a *AgentResponseMetadata) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentResponseMetadata) SetAgentResponseMetadataEvent(agentResponseMetadataEvent *AgentResponseMetadataAgentResponseMetadataEvent) {
	a.AgentResponseMetadataEvent = agentResponseMetadataEvent
}

func (a *AgentResponseMetadata) SetType(type_ *string) {
	a.Type = type_
}

func (a *AgentResponseMetadata) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentResponseMetadata) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentResponseMetadata to string"
	}
	return string(jsonData)
}

func (a *AgentResponseMetadata) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentResponseMetadata
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentResponseMetadata(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentResponseMetadata) MarshalJSON() ([]byte, error) {
	type embed AgentResponseMetadata
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
