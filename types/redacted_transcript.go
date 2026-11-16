package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type RedactedTranscript struct {
	MessageType     *string `json:"message_type,omitempty" xml:"message_type,omitempty"`
	Text            string  `json:"text" xml:"text" required:"true"`
	RedactedText    string  `json:"redacted_text" xml:"redacted_text" required:"true"`
	extraProperties map[string]interface{}
}

func (r *RedactedTranscript) GetMessageType() (value *string) {
	if r == nil {
		return
	}
	return r.MessageType
}

func (r *RedactedTranscript) GetText() (value string) {
	if r == nil {
		return
	}
	return r.Text
}

func (r *RedactedTranscript) GetRedactedText() (value string) {
	if r == nil {
		return
	}
	return r.RedactedText
}

func (r *RedactedTranscript) SetMessageType(messageType *string) {
	r.MessageType = messageType
}

func (r *RedactedTranscript) SetText(text string) {
	r.Text = text
}

func (r *RedactedTranscript) SetRedactedText(redactedText string) {
	r.RedactedText = redactedText
}

func (r *RedactedTranscript) GetExtraProperties() map[string]interface{} {
	if r == nil {
		return nil
	}
	return r.extraProperties
}

func (r RedactedTranscript) String() string {
	jsonData, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return "error converting struct: RedactedTranscript to string"
	}
	return string(jsonData)
}

func (r *RedactedTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, r); err != nil {
		return err
	}
	type alias RedactedTranscript
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*r = RedactedTranscript(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties
	return nil
}

func (r *RedactedTranscript) MarshalJSON() ([]byte, error) {
	type embed RedactedTranscript
	var marshaler = struct {
		embed
	}{
		embed: embed(*r),
	}
	return json.Marshal(marshaler)
}
