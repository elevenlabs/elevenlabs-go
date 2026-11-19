package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateTranslation struct {
	messageType     string
	Text            string `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslateTranslation) GetText() (value string) {
	if t == nil {
		return
	}
	return t.Text
}

func (t *TranslateTranslation) SetText(text string) {
	t.Text = text
}

func (t *TranslateTranslation) MessageType() string {
	return t.messageType
}

func (t *TranslateTranslation) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateTranslation) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateTranslation to string"
	}
	return string(jsonData)
}

func (t *TranslateTranslation) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateTranslation
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateTranslation(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "translation" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "translation", *literals.MessageType)
	}
	t.messageType = "translation"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateTranslation) MarshalJSON() ([]byte, error) {
	type embed TranslateTranslation
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "translation",
	}
	return json.Marshal(marshaler)
}
