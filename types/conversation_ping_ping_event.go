package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationPingPingEvent struct {
	EventId         int  `json:"event_id" xml:"event_id" required:"true"`
	PingMs          *int `json:"ping_ms,omitempty" xml:"ping_ms,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationPingPingEvent) GetEventId() (value int) {
	if c == nil {
		return
	}
	return c.EventId
}

func (c *ConversationPingPingEvent) GetPingMs() (value *int) {
	if c == nil {
		return
	}
	return c.PingMs
}

func (c *ConversationPingPingEvent) SetEventId(eventId int) {
	c.EventId = eventId
}

func (c *ConversationPingPingEvent) SetPingMs(pingMs *int) {
	c.PingMs = pingMs
}

func (c *ConversationPingPingEvent) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationPingPingEvent) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationPingPingEvent to string"
	}
	return string(jsonData)
}

func (c *ConversationPingPingEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationPingPingEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationPingPingEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationPingPingEvent) MarshalJSON() ([]byte, error) {
	type embed ConversationPingPingEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
