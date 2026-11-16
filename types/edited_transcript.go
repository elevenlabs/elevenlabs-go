package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type EditedTranscript struct {
	MessageType *string `json:"message_type,omitempty" xml:"message_type,omitempty"`
	// The committed transcript text the edit instruction was applied to.
	Text string `json:"text" xml:"text" required:"true"`
	// The edited transcript text. If no edits were made it will be identical to the `text` field.
	EditedText      string `json:"edited_text" xml:"edited_text" required:"true"`
	extraProperties map[string]interface{}
}

func (e *EditedTranscript) GetMessageType() (value *string) {
	if e == nil {
		return
	}
	return e.MessageType
}

func (e *EditedTranscript) GetText() (value string) {
	if e == nil {
		return
	}
	return e.Text
}

func (e *EditedTranscript) GetEditedText() (value string) {
	if e == nil {
		return
	}
	return e.EditedText
}

func (e *EditedTranscript) SetMessageType(messageType *string) {
	e.MessageType = messageType
}

func (e *EditedTranscript) SetText(text string) {
	e.Text = text
}

func (e *EditedTranscript) SetEditedText(editedText string) {
	e.EditedText = editedText
}

func (e *EditedTranscript) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EditedTranscript) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EditedTranscript to string"
	}
	return string(jsonData)
}

func (e *EditedTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, e); err != nil {
		return err
	}
	type alias EditedTranscript
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*e = EditedTranscript(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EditedTranscript) MarshalJSON() ([]byte, error) {
	type embed EditedTranscript
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
