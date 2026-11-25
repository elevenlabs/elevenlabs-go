package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubtitlingTranslationError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubtitlingTranslationError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *SubtitlingTranslationError) SetError(error string) {
	s.Error = error
}

func (s *SubtitlingTranslationError) MessageType() string {
	return s.messageType
}

func (s *SubtitlingTranslationError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingTranslationError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingTranslationError to string"
	}
	return string(jsonData)
}

func (s *SubtitlingTranslationError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubtitlingTranslationError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubtitlingTranslationError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "translation_error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "translation_error", *literals.MessageType)
	}
	s.messageType = "translation_error"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitlingTranslationError) MarshalJSON() ([]byte, error) {
	type embed SubtitlingTranslationError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "translation_error",
	}
	return json.Marshal(marshaler)
}
