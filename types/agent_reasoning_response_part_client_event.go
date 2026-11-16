package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentReasoningResponsePartClientEvent struct {
	ReasoningResponsePart *AgentReasoningResponsePartClientEventData `json:"reasoning_response_part" xml:"reasoning_response_part" required:"true"`
	extraProperties       map[string]interface{}
}

func (a *AgentReasoningResponsePartClientEvent) GetReasoningResponsePart() (value *AgentReasoningResponsePartClientEventData) {
	if a == nil {
		return
	}
	return a.ReasoningResponsePart
}

func (a *AgentReasoningResponsePartClientEvent) SetReasoningResponsePart(reasoningResponsePart *AgentReasoningResponsePartClientEventData) {
	a.ReasoningResponsePart = reasoningResponsePart
}

func (a *AgentReasoningResponsePartClientEvent) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentReasoningResponsePartClientEvent) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentReasoningResponsePartClientEvent to string"
	}
	return string(jsonData)
}

func (a *AgentReasoningResponsePartClientEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentReasoningResponsePartClientEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentReasoningResponsePartClientEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentReasoningResponsePartClientEvent) MarshalJSON() ([]byte, error) {
	type embed AgentReasoningResponsePartClientEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
