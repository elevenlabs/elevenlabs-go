package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationPong struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// ID of the ping event this pong responds to.
	EventId         int `json:"event_id" xml:"event_id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationPong) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationPong) GetEventId() (value int) {
	if c == nil {
		return
	}
	return c.EventId
}

func (c *ConversationPong) SetType(type_ *string) {
	c.Type = type_
}

func (c *ConversationPong) SetEventId(eventId int) {
	c.EventId = eventId
}

func (c *ConversationPong) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationPong) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationPong to string"
	}
	return string(jsonData)
}

func (c *ConversationPong) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationPong
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationPong(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationPong) MarshalJSON() ([]byte, error) {
	type embed ConversationPong
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
