package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeThrottledError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeThrottledError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeThrottledError) SetError(error string) {
	s.Error = error
}

func (s *ScribeThrottledError) MessageType() string {
	return s.messageType
}

func (s *ScribeThrottledError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeThrottledError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeThrottledError to string"
	}
	return string(jsonData)
}

func (s *ScribeThrottledError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeThrottledError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeThrottledError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "commit_throttled" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "commit_throttled", *literals.MessageType)
	}
	s.messageType = "commit_throttled"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *ScribeThrottledError) MarshalJSON() ([]byte, error) {
	type embed ScribeThrottledError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "commit_throttled",
	}
	return json.Marshal(marshaler)
}
