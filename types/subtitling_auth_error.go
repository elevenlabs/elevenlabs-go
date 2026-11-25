package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubtitlingAuthError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubtitlingAuthError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *SubtitlingAuthError) SetError(error string) {
	s.Error = error
}

func (s *SubtitlingAuthError) MessageType() string {
	return s.messageType
}

func (s *SubtitlingAuthError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingAuthError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingAuthError to string"
	}
	return string(jsonData)
}

func (s *SubtitlingAuthError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubtitlingAuthError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubtitlingAuthError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "auth_error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "auth_error", *literals.MessageType)
	}
	s.messageType = "auth_error"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitlingAuthError) MarshalJSON() ([]byte, error) {
	type embed SubtitlingAuthError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "auth_error",
	}
	return json.Marshal(marshaler)
}
