package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentChatResponsePartTextResponsePart struct {
	Text            string                                    `json:"text" xml:"text" required:"true"`
	Type            AgentChatResponsePartTextResponsePartType `json:"type" xml:"type" required:"true"`
	EventId         int                                       `json:"event_id" xml:"event_id" required:"true"`
	ResponseId      string                                    `json:"response_id" xml:"response_id" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AgentChatResponsePartTextResponsePart) GetText() (value string) {
	if a == nil {
		return
	}
	return a.Text
}

func (a *AgentChatResponsePartTextResponsePart) GetType() (value AgentChatResponsePartTextResponsePartType) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentChatResponsePartTextResponsePart) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentChatResponsePartTextResponsePart) GetResponseId() (value string) {
	if a == nil {
		return
	}
	return a.ResponseId
}

func (a *AgentChatResponsePartTextResponsePart) SetText(text string) {
	a.Text = text
}

func (a *AgentChatResponsePartTextResponsePart) SetType(type_ AgentChatResponsePartTextResponsePartType) {
	a.Type = type_
}

func (a *AgentChatResponsePartTextResponsePart) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentChatResponsePartTextResponsePart) SetResponseId(responseId string) {
	a.ResponseId = responseId
}

func (a *AgentChatResponsePartTextResponsePart) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentChatResponsePartTextResponsePart) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentChatResponsePartTextResponsePart to string"
	}
	return string(jsonData)
}

func (a *AgentChatResponsePartTextResponsePart) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentChatResponsePartTextResponsePart
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentChatResponsePartTextResponsePart(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentChatResponsePartTextResponsePart) MarshalJSON() ([]byte, error) {
	type embed AgentChatResponsePartTextResponsePart
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
