package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeQuotaExceededError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeQuotaExceededError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeQuotaExceededError) SetError(error string) {
	s.Error = error
}

func (s *ScribeQuotaExceededError) MessageType() string {
	return s.messageType
}

func (s *ScribeQuotaExceededError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeQuotaExceededError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeQuotaExceededError to string"
	}
	return string(jsonData)
}

func (s *ScribeQuotaExceededError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeQuotaExceededError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeQuotaExceededError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "quota_exceeded" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "quota_exceeded", *literals.MessageType)
	}
	s.messageType = "quota_exceeded"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScribeQuotaExceededError) MarshalJSON() ([]byte, error) {
	type embed ScribeQuotaExceededError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "quota_exceeded",
	}
	return json.Marshal(marshaler)
}
