package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentChatResponsePartClientEventData struct {
	Text            string                    `json:"text" xml:"text" required:"true"`
	Type            AgentChatResponsePartType `json:"type" xml:"type" required:"true"`
	EventId         int                       `json:"event_id" xml:"event_id" required:"true"`
	ResponseId      string                    `json:"response_id" xml:"response_id" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentChatResponsePartClientEventData) GetText() (value string) {
	if a == nil {
		return
	}
	return a.Text
}

func (a *AgentChatResponsePartClientEventData) GetType() (value AgentChatResponsePartType) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentChatResponsePartClientEventData) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentChatResponsePartClientEventData) GetResponseId() (value string) {
	if a == nil {
		return
	}
	return a.ResponseId
}

func (a *AgentChatResponsePartClientEventData) SetText(text string) {
	a.Text = text
}

func (a *AgentChatResponsePartClientEventData) SetType(type_ AgentChatResponsePartType) {
	a.Type = type_
}

func (a *AgentChatResponsePartClientEventData) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentChatResponsePartClientEventData) SetResponseId(responseId string) {
	a.ResponseId = responseId
}

func (a *AgentChatResponsePartClientEventData) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentChatResponsePartClientEventData) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentChatResponsePartClientEventData to string"
	}
	return string(jsonData)
}

func (a *AgentChatResponsePartClientEventData) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentChatResponsePartClientEventData
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentChatResponsePartClientEventData(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentChatResponsePartClientEventData) MarshalJSON() ([]byte, error) {
	type embed AgentChatResponsePartClientEventData
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
