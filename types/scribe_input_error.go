package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeInputError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeInputError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeInputError) SetError(error string) {
	s.Error = error
}

func (s *ScribeInputError) MessageType() string {
	return s.messageType
}

func (s *ScribeInputError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeInputError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeInputError to string"
	}
	return string(jsonData)
}

func (s *ScribeInputError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeInputError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeInputError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "input_error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "input_error", *literals.MessageType)
	}
	s.messageType = "input_error"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScribeInputError) MarshalJSON() ([]byte, error) {
	type embed ScribeInputError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "input_error",
	}
	return json.Marshal(marshaler)
}
