package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateQueueOverflowError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslateQueueOverflowError) GetError() (value string) {
	if t == nil {
		return
	}
	return t.Error
}

func (t *TranslateQueueOverflowError) SetError(error string) {
	t.Error = error
}

func (t *TranslateQueueOverflowError) MessageType() string {
	return t.messageType
}

func (t *TranslateQueueOverflowError) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateQueueOverflowError) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateQueueOverflowError to string"
	}
	return string(jsonData)
}

func (t *TranslateQueueOverflowError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateQueueOverflowError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateQueueOverflowError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "queue_overflow" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "queue_overflow", *literals.MessageType)
	}
	t.messageType = "queue_overflow"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateQueueOverflowError) MarshalJSON() ([]byte, error) {
	type embed TranslateQueueOverflowError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "queue_overflow",
	}
	return json.Marshal(marshaler)
}
