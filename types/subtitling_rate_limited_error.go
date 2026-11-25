package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubtitlingRateLimitedError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubtitlingRateLimitedError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *SubtitlingRateLimitedError) SetError(error string) {
	s.Error = error
}

func (s *SubtitlingRateLimitedError) MessageType() string {
	return s.messageType
}

func (s *SubtitlingRateLimitedError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingRateLimitedError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingRateLimitedError to string"
	}
	return string(jsonData)
}

func (s *SubtitlingRateLimitedError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubtitlingRateLimitedError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubtitlingRateLimitedError(tmp)
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

func (s *SubtitlingRateLimitedError) MarshalJSON() ([]byte, error) {
	type embed SubtitlingRateLimitedError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "rate_limited",
	}
	return json.Marshal(marshaler)
}
