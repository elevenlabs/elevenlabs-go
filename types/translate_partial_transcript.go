package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslatePartialTranscript struct {
	messageType     string
	Text            string `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslatePartialTranscript) GetText() (value string) {
	if t == nil {
		return
	}
	return t.Text
}

func (t *TranslatePartialTranscript) SetText(text string) {
	t.Text = text
}

func (t *TranslatePartialTranscript) MessageType() string {
	return t.messageType
}

func (t *TranslatePartialTranscript) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslatePartialTranscript) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslatePartialTranscript to string"
	}
	return string(jsonData)
}

func (t *TranslatePartialTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslatePartialTranscript
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslatePartialTranscript(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "partial_transcript" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "partial_transcript", *literals.MessageType)
	}
	t.messageType = "partial_transcript"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslatePartialTranscript) MarshalJSON() ([]byte, error) {
	type embed TranslatePartialTranscript
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "partial_transcript",
	}
	return json.Marshal(marshaler)
}
