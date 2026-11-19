package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeInvalidRequestError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeInvalidRequestError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeInvalidRequestError) SetError(error string) {
	s.Error = error
}

func (s *ScribeInvalidRequestError) MessageType() string {
	return s.messageType
}

func (s *ScribeInvalidRequestError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeInvalidRequestError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeInvalidRequestError to string"
	}
	return string(jsonData)
}

func (s *ScribeInvalidRequestError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeInvalidRequestError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeInvalidRequestError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "invalid_request" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "invalid_request", *literals.MessageType)
	}
	s.messageType = "invalid_request"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScribeInvalidRequestError) MarshalJSON() ([]byte, error) {
	type embed ScribeInvalidRequestError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "invalid_request",
	}
	return json.Marshal(marshaler)
}
