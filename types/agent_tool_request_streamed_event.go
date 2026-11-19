package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentToolRequestStreamedEvent struct {
	AgentToolRequest *AgentToolRequestClientEventData `json:"agent_tool_request" xml:"agent_tool_request" required:"true"`
	extraProperties  map[string]interface{}
}

func (a *AgentToolRequestStreamedEvent) GetAgentToolRequest() (value *AgentToolRequestClientEventData) {
	if a == nil {
		return
	}
	return a.AgentToolRequest
}

func (a *AgentToolRequestStreamedEvent) SetAgentToolRequest(agentToolRequest *AgentToolRequestClientEventData) {
	a.AgentToolRequest = agentToolRequest
}

func (a *AgentToolRequestStreamedEvent) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentToolRequestStreamedEvent) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentToolRequestStreamedEvent to string"
	}
	return string(jsonData)
}

func (a *AgentToolRequestStreamedEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentToolRequestStreamedEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentToolRequestStreamedEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentToolRequestStreamedEvent) MarshalJSON() ([]byte, error) {
	type embed AgentToolRequestStreamedEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
