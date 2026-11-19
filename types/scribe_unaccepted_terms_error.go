package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ScribeUnacceptedTermsError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *ScribeUnacceptedTermsError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *ScribeUnacceptedTermsError) SetError(error string) {
	s.Error = error
}

func (s *ScribeUnacceptedTermsError) MessageType() string {
	return s.messageType
}

func (s *ScribeUnacceptedTermsError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s ScribeUnacceptedTermsError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: ScribeUnacceptedTermsError to string"
	}
	return string(jsonData)
}

func (s *ScribeUnacceptedTermsError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias ScribeUnacceptedTermsError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = ScribeUnacceptedTermsError(tmp)
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

func (s *ScribeUnacceptedTermsError) MarshalJSON() ([]byte, error) {
	type embed ScribeUnacceptedTermsError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "unaccepted_terms",
	}
	return json.Marshal(marshaler)
}
