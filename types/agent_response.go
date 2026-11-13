package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentResponse struct {
	type_           string
	EventId         int     `json:"event_id" xml:"event_id" required:"true"`
	Content         *string `json:"content,omitempty" xml:"content,omitempty"`
	IsFinal         *bool   `json:"is_final,omitempty" xml:"is_final,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AgentResponse) GetEventId() (value int) {
	if a == nil {
		return
	}
	return a.EventId
}

func (a *AgentResponse) GetContent() (value *string) {
	if a == nil {
		return
	}
	return a.Content
}

func (a *AgentResponse) GetIsFinal() (value *bool) {
	if a == nil {
		return
	}
	return a.IsFinal
}

func (a *AgentResponse) SetEventId(eventId int) {
	a.EventId = eventId
}

func (a *AgentResponse) SetContent(content *string) {
	a.Content = content
}

func (a *AgentResponse) SetIsFinal(isFinal *bool) {
	a.IsFinal = isFinal
}

func (a *AgentResponse) Type() string {
	return a.type_
}

func (a *AgentResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentResponse to string"
	}
	return string(jsonData)
}

func (a *AgentResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentResponse(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "agent_response" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", a, "agent_response", *literals.Type)
	}
	a.type_ = "agent_response"
	extraProperties, err := internal.ExtractExtraProperties(data, *a, "type")
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentResponse) MarshalJSON() ([]byte, error) {
	type embed AgentResponse
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*a),
		Type:  "agent_response",
	}
	return json.Marshal(marshaler)
}
