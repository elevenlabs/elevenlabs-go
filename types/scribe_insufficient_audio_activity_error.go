package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeInsufficientAudioActivityError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeInsufficientAudioActivityError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeInsufficientAudioActivityError) SetError(error string) {
	s.Error = error
}

func (s *ScribeInsufficientAudioActivityError) MessageType() string {
	return s.messageType
}

func (s *ScribeInsufficientAudioActivityError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeInsufficientAudioActivityError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeInsufficientAudioActivityError to string"
	}
	return string(jsonData)
}

func (s *ScribeInsufficientAudioActivityError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeInsufficientAudioActivityError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeInsufficientAudioActivityError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "insufficient_audio_activity" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "insufficient_audio_activity", *literals.MessageType)
	}
	s.messageType = "insufficient_audio_activity"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScribeInsufficientAudioActivityError) MarshalJSON() ([]byte, error) {
	type embed ScribeInsufficientAudioActivityError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "insufficient_audio_activity",
	}
	return json.Marshal(marshaler)
}
