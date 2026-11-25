package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationUserTranscript struct {
	UserTranscriptionEvent *ConversationUserTranscriptUserTranscriptionEvent `json:"user_transcription_event" xml:"user_transcription_event" required:"true"`
	Type                   *string                                           `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties        map[string]interface{}
}

func (c *ConversationUserTranscript) GetUserTranscriptionEvent() (value *ConversationUserTranscriptUserTranscriptionEvent) {
	if c == nil {
		return
	}
	return c.UserTranscriptionEvent
}

func (c *ConversationUserTranscript) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationUserTranscript) SetUserTranscriptionEvent(userTranscriptionEvent *ConversationUserTranscriptUserTranscriptionEvent) {
	c.UserTranscriptionEvent = userTranscriptionEvent
}

func (c *ConversationUserTranscript) SetType(type_ *string) {
	c.Type = type_
}

func (c *ConversationUserTranscript) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationUserTranscript) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationUserTranscript to string"
	}
	return string(jsonData)
}

func (c *ConversationUserTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationUserTranscript
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationUserTranscript(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationUserTranscript) MarshalJSON() ([]byte, error) {
	type embed ConversationUserTranscript
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
