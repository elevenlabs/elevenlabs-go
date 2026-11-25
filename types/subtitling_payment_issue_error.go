package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubtitlingPaymentIssueError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubtitlingPaymentIssueError) GetError() (value string) {
	if s == nil {
		return
	}
	return s.Error
}

func (s *SubtitlingPaymentIssueError) SetError(error string) {
	s.Error = error
}

func (s *SubtitlingPaymentIssueError) MessageType() string {
	return s.messageType
}

func (s *SubtitlingPaymentIssueError) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingPaymentIssueError) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingPaymentIssueError to string"
	}
	return string(jsonData)
}

func (s *SubtitlingPaymentIssueError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubtitlingPaymentIssueError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubtitlingPaymentIssueError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "payment_issue" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "payment_issue", *literals.MessageType)
	}
	s.messageType = "payment_issue"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitlingPaymentIssueError) MarshalJSON() ([]byte, error) {
	type embed SubtitlingPaymentIssueError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "payment_issue",
	}
	return json.Marshal(marshaler)
}
