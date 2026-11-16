package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeRateLimitedError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeRateLimitedError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeRateLimitedError) SetError(error string) {
	s.Error = error
}

func (s *ScribeRateLimitedError) MessageType() string {
	return s.messageType
}

func (s *ScribeRateLimitedError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeRateLimitedError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeRateLimitedError to string"
	}
	return string(jsonData)
}

func (s *ScribeRateLimitedError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeRateLimitedError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeRateLimitedError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "rate_limited" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "rate_limited", *literals.MessageType)
	}
	s.messageType = "rate_limited"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScribeRateLimitedError) MarshalJSON() ([]byte, error) {
	type embed ScribeRateLimitedError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "rate_limited",
	}
	return json.Marshal(marshaler)
}
