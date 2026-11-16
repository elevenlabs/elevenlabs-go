package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentReasoningResponsePartClientEventData struct {
	Text            string                    `json:"text" xml:"text" required:"true"`
	Type            AgentChatResponsePartType `json:"type" xml:"type" required:"true"`
	EventId         int                       `json:"event_id" xml:"event_id" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentReasoningResponsePartClientEventData) GetText() (value string) {
	if a == nil {
		return
	}
	return a.Text
}

func (a *AgentReasoningResponsePartClientEventData) GetType() (value AgentChatResponsePartType) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentReasoningResponsePartClientEventData) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentReasoningResponsePartClientEventData) SetText(text string) {
	a.Text = text
}

func (a *AgentReasoningResponsePartClientEventData) SetType(type_ AgentChatResponsePartType) {
	a.Type = type_
}

func (a *AgentReasoningResponsePartClientEventData) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentReasoningResponsePartClientEventData) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentReasoningResponsePartClientEventData) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentReasoningResponsePartClientEventData to string"
	}
	return string(jsonData)
}

func (a *AgentReasoningResponsePartClientEventData) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentReasoningResponsePartClientEventData
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentReasoningResponsePartClientEventData(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentReasoningResponsePartClientEventData) MarshalJSON() ([]byte, error) {
	type embed AgentReasoningResponsePartClientEventData
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
