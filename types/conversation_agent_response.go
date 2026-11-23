package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationAgentResponse struct {
	AgentResponseEvent *ConversationAgentResponseAgentResponseEvent `json:"agent_response_event" xml:"agent_response_event" required:"true"`
	Type               *string                                      `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties    map[string]interface{}
}

func (c *ConversationAgentResponse) GetAgentResponseEvent() (value *ConversationAgentResponseAgentResponseEvent) {
	if c == nil {
		return
	}
	return c.AgentResponseEvent
}

func (c *ConversationAgentResponse) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationAgentResponse) SetAgentResponseEvent(agentResponseEvent *ConversationAgentResponseAgentResponseEvent) {
	c.AgentResponseEvent = agentResponseEvent
}

func (c *ConversationAgentResponse) SetType(type_ *string) {
	c.Type = type_
}

func (c *ConversationAgentResponse) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationAgentResponse) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationAgentResponse to string"
	}
	return string(jsonData)
}

func (c *ConversationAgentResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationAgentResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationAgentResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationAgentResponse) MarshalJSON() ([]byte, error) {
	type embed ConversationAgentResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
