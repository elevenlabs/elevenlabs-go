package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationPing struct {
	PingEvent       *ConversationPingPingEvent `json:"ping_event" xml:"ping_event" required:"true"`
	Type            *string                    `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ConversationPing) GetPingEvent() (value *ConversationPingPingEvent) {
	if c == nil {
		return
	}
	return c.PingEvent
}

func (c *ConversationPing) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationPing) SetPingEvent(pingEvent *ConversationPingPingEvent) {
	c.PingEvent = pingEvent
}

func (c *ConversationPing) SetType(type_ *string) {
	c.Type = type_
}

func (c *ConversationPing) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationPing) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationPing to string"
	}
	return string(jsonData)
}

func (c *ConversationPing) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationPing
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationPing(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationPing) MarshalJSON() ([]byte, error) {
	type embed ConversationPing
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
