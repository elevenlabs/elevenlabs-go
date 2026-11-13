package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeAuthError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeAuthError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeAuthError) SetError(error string) {
	s.Error = error
}

func (s *ScribeAuthError) MessageType() string {
	return s.messageType
}

func (s *ScribeAuthError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeAuthError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeAuthError to string"
	}
	return string(jsonData)
}

func (s *ScribeAuthError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeAuthError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeAuthError(tmp)
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

func (s *ScribeAuthError) MarshalJSON() ([]byte, error) {
	type embed ScribeAuthError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "auth_error",
	}
	return json.Marshal(marshaler)
}
