package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeQueueOverflowError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeQueueOverflowError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeQueueOverflowError) SetError(error string) {
	s.Error = error
}

func (s *ScribeQueueOverflowError) MessageType() string {
	return s.messageType
}

func (s *ScribeQueueOverflowError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeQueueOverflowError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeQueueOverflowError to string"
	}
	return string(jsonData)
}

func (s *ScribeQueueOverflowError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeQueueOverflowError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeQueueOverflowError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "queue_overflow" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "queue_overflow", *literals.MessageType)
	}
	s.messageType = "queue_overflow"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScribeQueueOverflowError) MarshalJSON() ([]byte, error) {
	type embed ScribeQueueOverflowError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "queue_overflow",
	}
	return json.Marshal(marshaler)
}
