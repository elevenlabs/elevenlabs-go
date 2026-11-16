package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateTranslationError struct {
	messageType     string
	Error           string `json:"error" xml:"error" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslateTranslationError) GetError() (value string) {
	if t == nil {
		return
	}
	return t.Error
}

func (t *TranslateTranslationError) SetError(error string) {
	t.Error = error
}

func (t *TranslateTranslationError) MessageType() string {
	return t.messageType
}

func (t *TranslateTranslationError) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateTranslationError) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateTranslationError to string"
	}
	return string(jsonData)
}

func (t *TranslateTranslationError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateTranslationError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateTranslationError(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "translation_error" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "translation_error", *literals.MessageType)
	}
	t.messageType = "translation_error"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateTranslationError) MarshalJSON() ([]byte, error) {
	type embed TranslateTranslationError
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "translation_error",
	}
	return json.Marshal(marshaler)
}
