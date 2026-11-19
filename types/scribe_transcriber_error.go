package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeTranscriberError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeTranscriberError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeTranscriberError) SetError(error string) {
	s.Error = error
}

func (s *ScribeTranscriberError) MessageType() string {
	return s.messageType
}

func (s *ScribeTranscriberError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeTranscriberError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeTranscriberError to string"
	}
	return string(jsonData)
}

func (s *ScribeTranscriberError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeTranscriberError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeTranscriberError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "transcriber_error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "transcriber_error", *literals.MessageType)
	}
	s.messageType = "transcriber_error"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScribeTranscriberError) MarshalJSON() ([]byte, error) {
	type embed ScribeTranscriberError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "transcriber_error",
	}
	return json.Marshal(marshaler)
}
