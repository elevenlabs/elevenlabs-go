package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubtitlingSessionStarted struct {
	messageType     string
	SessionId       string  `json:"session_id" xml:"session_id" required:"true"`
	ClientSessionId *string `json:"client_session_id,omitempty" xml:"client_session_id,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SubtitlingSessionStarted) GetSessionId() (value string) {
	if s == nil {
		return
	}
	return s.SessionId
}

func (s *SubtitlingSessionStarted) GetClientSessionId() (value *string) {
	if s == nil {
		return
	}
	return s.ClientSessionId
}

func (s *SubtitlingSessionStarted) SetSessionId(sessionId string) {
	s.SessionId = sessionId
}

func (s *SubtitlingSessionStarted) SetClientSessionId(clientSessionId *string) {
	s.ClientSessionId = clientSessionId
}

func (s *SubtitlingSessionStarted) MessageType() string {
	return s.messageType
}

func (s *SubtitlingSessionStarted) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingSessionStarted) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingSessionStarted to string"
	}
	return string(jsonData)
}

func (s *SubtitlingSessionStarted) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "session_started" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "session_started", *literals.MessageType)
	}
	s.messageType = "session_started"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitlingSessionStarted) MarshalJSON() ([]byte, error) {
	type embed SubtitlingSessionStarted
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "session_started",
	}
	return json.Marshal(marshaler)
}
