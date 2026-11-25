package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentToolResponseFullPayload struct {
	AgentToolResponseFullPayload *AgentToolResponseFullPayloadAgentToolResponseFullPayload `json:"agent_tool_response_full_payload" xml:"agent_tool_response_full_payload" required:"true"`
	Type                         *string                                                   `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties              map[string]interface{}
}

func (a *AgentToolResponseFullPayload) GetAgentToolResponseFullPayload() (value *AgentToolResponseFullPayloadAgentToolResponseFullPayload) {
	if a == nil {
		return
	}
	return a.AgentToolResponseFullPayload
}

func (a *AgentToolResponseFullPayload) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentToolResponseFullPayload) SetAgentToolResponseFullPayload(agentToolResponseFullPayload *AgentToolResponseFullPayloadAgentToolResponseFullPayload) {
	a.AgentToolResponseFullPayload = agentToolResponseFullPayload
}

func (a *AgentToolResponseFullPayload) SetType(type_ *string) {
	a.Type = type_
}

func (a *AgentToolResponseFullPayload) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentToolResponseFullPayload) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentToolResponseFullPayload to string"
	}
	return string(jsonData)
}

func (a *AgentToolResponseFullPayload) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentToolResponseFullPayload
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentToolResponseFullPayload(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentToolResponseFullPayload) MarshalJSON() ([]byte, error) {
	type embed AgentToolResponseFullPayload
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
