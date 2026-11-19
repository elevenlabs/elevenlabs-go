package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type StreamEnvelope struct {
	ConversationId  string               `json:"conversation_id" xml:"conversation_id" required:"true"`
	TurnId          string               `json:"turn_id" xml:"turn_id" required:"true"`
	CreatedAtUnixMs int                  `json:"created_at_unix_ms" xml:"created_at_unix_ms" required:"true"`
	Event           *StreamEnvelopeEvent `json:"event" xml:"event" required:"true"`
	extraProperties map[string]interface{}
}

func (s *StreamEnvelope) GetConversationId() (value string) {
	if s == nil {
		return
	}
	return s.ConversationId
}

func (s *StreamEnvelope) GetTurnId() (value string) {
	if s == nil {
		return
	}
	return s.TurnId
}

func (s *StreamEnvelope) GetCreatedAtUnixMs() (value int) {
	if s == nil {
		return
	}
	return s.CreatedAtUnixMs
}

func (s *StreamEnvelope) GetEvent() (value *StreamEnvelopeEvent) {
	if s == nil {
		return
	}
	return s.Event
}

func (s *StreamEnvelope) SetConversationId(conversationId string) {
	s.ConversationId = conversationId
}

func (s *StreamEnvelope) SetTurnId(turnId string) {
	s.TurnId = turnId
}

func (s *StreamEnvelope) SetCreatedAtUnixMs(createdAtUnixMs int) {
	s.CreatedAtUnixMs = createdAtUnixMs
}

func (s *StreamEnvelope) SetEvent(event *StreamEnvelopeEvent) {
	s.Event = event
}

func (s *StreamEnvelope) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StreamEnvelope) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StreamEnvelope to string"
	}
	return string(jsonData)
}

func (s *StreamEnvelope) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias StreamEnvelope
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = StreamEnvelope(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StreamEnvelope) MarshalJSON() ([]byte, error) {
	type embed StreamEnvelope
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
