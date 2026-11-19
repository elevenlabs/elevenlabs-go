package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateFinalTranscript struct {
	messageType     string
	Text            string `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslateFinalTranscript) GetText() (value string) {
	if t == nil {
		return
	}
	return t.Text
}

func (t *TranslateFinalTranscript) SetText(text string) {
	t.Text = text
}

func (t *TranslateFinalTranscript) MessageType() string {
	return t.messageType
}

func (t *TranslateFinalTranscript) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateFinalTranscript) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateFinalTranscript to string"
	}
	return string(jsonData)
}

func (t *TranslateFinalTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateFinalTranscript
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateFinalTranscript(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "final_transcript" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "final_transcript", *literals.MessageType)
	}
	t.messageType = "final_transcript"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateFinalTranscript) MarshalJSON() ([]byte, error) {
	type embed TranslateFinalTranscript
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "final_transcript",
	}
	return json.Marshal(marshaler)
}
