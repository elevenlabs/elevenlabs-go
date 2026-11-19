package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentChatResponsePartClientEvent struct {
	TextResponsePart *AgentChatResponsePartClientEventData `json:"text_response_part" xml:"text_response_part" required:"true"`
	extraProperties  map[string]interface{}
}

func (a *AgentChatResponsePartClientEvent) GetTextResponsePart() (value *AgentChatResponsePartClientEventData) {
	if a == nil {
		return
	}
	return a.TextResponsePart
}

func (a *AgentChatResponsePartClientEvent) SetTextResponsePart(textResponsePart *AgentChatResponsePartClientEventData) {
	a.TextResponsePart = textResponsePart
}

func (a *AgentChatResponsePartClientEvent) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentChatResponsePartClientEvent) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentChatResponsePartClientEvent to string"
	}
	return string(jsonData)
}

func (a *AgentChatResponsePartClientEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentChatResponsePartClientEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentChatResponsePartClientEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentChatResponsePartClientEvent) MarshalJSON() ([]byte, error) {
	type embed AgentChatResponsePartClientEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
