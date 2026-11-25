package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationAgentResponseAgentResponseEvent struct {
	AgentResponse   string                                                        `json:"agent_response" xml:"agent_response" required:"true"`
	EventId         int                                                           `json:"event_id" xml:"event_id" required:"true"`
	ResponseId      string                                                        `json:"response_id" xml:"response_id" required:"true"`
	InResponseToIds []string                                                      `json:"in_response_to_ids,omitempty" xml:"in_response_to_ids,omitempty"`
	ResponseType    *ConversationAgentResponseAgentResponseEventResponseType      `json:"response_type,omitempty" xml:"response_type,omitempty"`
	Attachments     []*ConversationAgentResponseAgentResponseEventAttachmentsItem `json:"attachments,omitempty" xml:"attachments,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationAgentResponseAgentResponseEvent) GetAgentResponse() (value string) {
	if c == nil {
		return
	}
	return c.AgentResponse
}

func (c *ConversationAgentResponseAgentResponseEvent) GetEventId() (value int) {
	if c == nil {
		return
	}
	return c.EventId
}

func (c *ConversationAgentResponseAgentResponseEvent) GetResponseId() (value string) {
	if c == nil {
		return
	}
	return c.ResponseId
}

func (c *ConversationAgentResponseAgentResponseEvent) GetInResponseToIds() (value []string) {
	if c == nil {
		return
	}
	return c.InResponseToIds
}

func (c *ConversationAgentResponseAgentResponseEvent) GetResponseType() (value *ConversationAgentResponseAgentResponseEventResponseType) {
	if c == nil {
		return
	}
	return c.ResponseType
}

func (c *ConversationAgentResponseAgentResponseEvent) GetAttachments() (value []*ConversationAgentResponseAgentResponseEventAttachmentsItem) {
	if c == nil {
		return
	}
	return c.Attachments
}

func (c *ConversationAgentResponseAgentResponseEvent) SetAgentResponse(agentResponse string) {
	c.AgentResponse = agentResponse
}

func (c *ConversationAgentResponseAgentResponseEvent) SetEventId(eventId int) {
	c.EventId = eventId
}

func (c *ConversationAgentResponseAgentResponseEvent) SetResponseId(responseId string) {
	c.ResponseId = responseId
}

func (c *ConversationAgentResponseAgentResponseEvent) SetInResponseToIds(inResponseToIds []string) {
	c.InResponseToIds = inResponseToIds
}

func (c *ConversationAgentResponseAgentResponseEvent) SetResponseType(responseType *ConversationAgentResponseAgentResponseEventResponseType) {
	c.ResponseType = responseType
}

func (c *ConversationAgentResponseAgentResponseEvent) SetAttachments(attachments []*ConversationAgentResponseAgentResponseEventAttachmentsItem) {
	c.Attachments = attachments
}

func (c *ConversationAgentResponseAgentResponseEvent) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationAgentResponseAgentResponseEvent) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationAgentResponseAgentResponseEvent to string"
	}
	return string(jsonData)
}

func (c *ConversationAgentResponseAgentResponseEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationAgentResponseAgentResponseEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationAgentResponseAgentResponseEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationAgentResponseAgentResponseEvent) MarshalJSON() ([]byte, error) {
	type embed ConversationAgentResponseAgentResponseEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
