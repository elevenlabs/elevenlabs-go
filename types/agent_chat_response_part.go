package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentChatResponsePart struct {
	TextResponsePart *AgentChatResponsePartTextResponsePart `json:"text_response_part" xml:"text_response_part" required:"true"`
	Type             *string                                `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties  map[string]interface{}
}

func (a *AgentChatResponsePart) GetTextResponsePart() (value *AgentChatResponsePartTextResponsePart) {
	if a == nil {
		return
	}
	return a.TextResponsePart
}

func (a *AgentChatResponsePart) GetType() (value *string) {
	if a == nil {
		return
	}
	return a.Type
}

func (a *AgentChatResponsePart) SetTextResponsePart(textResponsePart *AgentChatResponsePartTextResponsePart) {
	a.TextResponsePart = textResponsePart
}

func (a *AgentChatResponsePart) SetType(type_ *string) {
	a.Type = type_
}

func (a *AgentChatResponsePart) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentChatResponsePart) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentChatResponsePart to string"
	}
	return string(jsonData)
}

func (a *AgentChatResponsePart) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, a); err != nil {
		return err
	}
	type alias AgentChatResponsePart
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*a = AgentChatResponsePart(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AgentChatResponsePart) MarshalJSON() ([]byte, error) {
	type embed AgentChatResponsePart
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
