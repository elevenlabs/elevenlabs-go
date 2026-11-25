package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationUserTranscriptUserTranscriptionEvent struct {
	UserTranscript  string `json:"user_transcript" xml:"user_transcript" required:"true"`
	EventId         int    `json:"event_id" xml:"event_id" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ConversationUserTranscriptUserTranscriptionEvent) GetUserTranscript() (value string) {
	if c == nil {
		return
	}
	return c.UserTranscript
}

func (c *ConversationUserTranscriptUserTranscriptionEvent) GetEventId() (value int) {
	if c == nil {
		return
	}
	return c.EventId
}

func (c *ConversationUserTranscriptUserTranscriptionEvent) SetUserTranscript(userTranscript string) {
	c.UserTranscript = userTranscript
}

func (c *ConversationUserTranscriptUserTranscriptionEvent) SetEventId(eventId int) {
	c.EventId = eventId
}

func (c *ConversationUserTranscriptUserTranscriptionEvent) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationUserTranscriptUserTranscriptionEvent) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationUserTranscriptUserTranscriptionEvent to string"
	}
	return string(jsonData)
}

func (c *ConversationUserTranscriptUserTranscriptionEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationUserTranscriptUserTranscriptionEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationUserTranscriptUserTranscriptionEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationUserTranscriptUserTranscriptionEvent) MarshalJSON() ([]byte, error) {
	type embed ConversationUserTranscriptUserTranscriptionEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
