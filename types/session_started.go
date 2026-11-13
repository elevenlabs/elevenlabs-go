package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SessionStarted struct {
	MessageType     *string               `json:"message_type,omitempty" xml:"message_type,omitempty"`
	SessionId       string                `json:"session_id" xml:"session_id" required:"true"`
	Config          *SessionStartedConfig `json:"config" xml:"config" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SessionStarted) GetMessageType() (value *string) {
	if s == nil {
		return
	}
	return s.MessageType
}

func (s *SessionStarted) GetSessionId() (value string) {
	if s == nil {
		return
	}
	return s.SessionId
}

func (s *SessionStarted) GetConfig() (value *SessionStartedConfig) {
	if s == nil {
		return
	}
	return s.Config
}

func (s *SessionStarted) SetMessageType(messageType *string) {
	s.MessageType = messageType
}

func (s *SessionStarted) SetSessionId(sessionId string) {
	s.SessionId = sessionId
}

func (s *SessionStarted) SetConfig(config *SessionStartedConfig) {
	s.Config = config
}

func (s *SessionStarted) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SessionStarted) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SessionStarted to string"
	}
	return string(jsonData)
}

func (s *SessionStarted) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SessionStarted
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SessionStarted(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SessionStarted) MarshalJSON() ([]byte, error) {
	type embed SessionStarted
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
