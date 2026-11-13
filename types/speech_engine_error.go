package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeechEngineError struct {
	type_ string
	// Human-readable description of the error.
	Message         string `json:"message" xml:"message" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SpeechEngineError) GetMessage() (value string) {
	if s == nil {
		return
	}
	return s.Message
}

func (s *SpeechEngineError) SetMessage(message string) {
	s.Message = message
}

func (s *SpeechEngineError) Type() string {
	return s.type_
}

func (s *SpeechEngineError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeechEngineError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeechEngineError to string"
	}
	return string(jsonData)
}

func (s *SpeechEngineError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SpeechEngineError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SpeechEngineError(tmp)
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "error", *literals.Type)
	}
	s.type_ = "error"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeechEngineError) MarshalJSON() ([]byte, error) {
	type embed SpeechEngineError
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*s),
		Type:  "error",
	}
	return json.Marshal(marshaler)
}
