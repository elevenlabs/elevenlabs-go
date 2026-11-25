package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubtitlingUnacceptedTermsError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubtitlingUnacceptedTermsError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *SubtitlingUnacceptedTermsError) SetError(error string) {
	s.Error = error
}

func (s *SubtitlingUnacceptedTermsError) MessageType() string {
	return s.messageType
}

func (s *SubtitlingUnacceptedTermsError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingUnacceptedTermsError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingUnacceptedTermsError to string"
	}
	return string(jsonData)
}

func (s *SubtitlingUnacceptedTermsError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubtitlingUnacceptedTermsError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubtitlingUnacceptedTermsError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "unaccepted_terms" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "unaccepted_terms", *literals.MessageType)
	}
	s.messageType = "unaccepted_terms"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitlingUnacceptedTermsError) MarshalJSON() ([]byte, error) {
	type embed SubtitlingUnacceptedTermsError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "unaccepted_terms",
	}
	return json.Marshal(marshaler)
}
