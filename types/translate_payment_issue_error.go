package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslatePaymentIssueError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslatePaymentIssueError) GetError() (value string) {
	if t == nil {
		return
	}
	return t.Error
}

func (t *TranslatePaymentIssueError) SetError(error string) {
	t.Error = error
}

func (t *TranslatePaymentIssueError) MessageType() string {
	return t.messageType
}

func (t *TranslatePaymentIssueError) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslatePaymentIssueError) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslatePaymentIssueError to string"
	}
	return string(jsonData)
}

func (t *TranslatePaymentIssueError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslatePaymentIssueError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslatePaymentIssueError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "payment_issue" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "payment_issue", *literals.MessageType)
	}
	t.messageType = "payment_issue"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslatePaymentIssueError) MarshalJSON() ([]byte, error) {
	type embed TranslatePaymentIssueError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "payment_issue",
	}
	return json.Marshal(marshaler)
}
