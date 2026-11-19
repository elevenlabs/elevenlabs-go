package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslateError) GetError() (value string) {
	if t == nil {
		return
	}
	return t.Error
}

func (t *TranslateError) SetError(error string) {
	t.Error = error
}

func (t *TranslateError) MessageType() string {
	return t.messageType
}

func (t *TranslateError) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateError) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateError to string"
	}
	return string(jsonData)
}

func (t *TranslateError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "error", *literals.MessageType)
	}
	t.messageType = "error"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateError) MarshalJSON() ([]byte, error) {
	type embed TranslateError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "error",
	}
	return json.Marshal(marshaler)
}
