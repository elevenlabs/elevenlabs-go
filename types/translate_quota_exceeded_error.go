package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateQuotaExceededError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslateQuotaExceededError) GetError() (value string) {
	if t == nil {
		return
	}
	return t.Error
}

func (t *TranslateQuotaExceededError) SetError(error string) {
	t.Error = error
}

func (t *TranslateQuotaExceededError) MessageType() string {
	return t.messageType
}

func (t *TranslateQuotaExceededError) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateQuotaExceededError) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateQuotaExceededError to string"
	}
	return string(jsonData)
}

func (t *TranslateQuotaExceededError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateQuotaExceededError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateQuotaExceededError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "quota_exceeded" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "quota_exceeded", *literals.MessageType)
	}
	t.messageType = "quota_exceeded"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateQuotaExceededError) MarshalJSON() ([]byte, error) {
	type embed TranslateQuotaExceededError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "quota_exceeded",
	}
	return json.Marshal(marshaler)
}
