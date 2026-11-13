package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeSessionTimeLimitExceededError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeSessionTimeLimitExceededError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeSessionTimeLimitExceededError) SetError(error string) {
	s.Error = error
}

func (s *ScribeSessionTimeLimitExceededError) MessageType() string {
	return s.messageType
}

func (s *ScribeSessionTimeLimitExceededError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeSessionTimeLimitExceededError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeSessionTimeLimitExceededError to string"
	}
	return string(jsonData)
}

func (s *ScribeSessionTimeLimitExceededError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeSessionTimeLimitExceededError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeSessionTimeLimitExceededError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "session_time_limit_exceeded" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "session_time_limit_exceeded", *literals.MessageType)
	}
	s.messageType = "session_time_limit_exceeded"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScribeSessionTimeLimitExceededError) MarshalJSON() ([]byte, error) {
	type embed ScribeSessionTimeLimitExceededError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "session_time_limit_exceeded",
	}
	return json.Marshal(marshaler)
}
