package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeResourceExhaustedError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeResourceExhaustedError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeResourceExhaustedError) SetError(error string) {
	s.Error = error
}

func (s *ScribeResourceExhaustedError) MessageType() string {
	return s.messageType
}

func (s *ScribeResourceExhaustedError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeResourceExhaustedError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeResourceExhaustedError to string"
	}
	return string(jsonData)
}

func (s *ScribeResourceExhaustedError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeResourceExhaustedError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeResourceExhaustedError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "resource_exhausted" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "resource_exhausted", *literals.MessageType)
	}
	s.messageType = "resource_exhausted"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScribeResourceExhaustedError) MarshalJSON() ([]byte, error) {
	type embed ScribeResourceExhaustedError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "resource_exhausted",
	}
	return json.Marshal(marshaler)
}
