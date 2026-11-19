package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingPartialTranscript struct {
	messageType     string
	Text            string `json:"text" xml:"text" required:"true"`
	extraProperties map[string]interface{}
}

func (d *DubbingPartialTranscript) GetText() (value string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DubbingPartialTranscript) SetText(text string) {
	d.Text = text
}

func (d *DubbingPartialTranscript) MessageType() string {
	return d.messageType
}

func (d *DubbingPartialTranscript) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingPartialTranscript) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingPartialTranscript to string"
	}
	return string(jsonData)
}

func (d *DubbingPartialTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingPartialTranscript
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingPartialTranscript(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "partial_transcript" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "partial_transcript", *literals.MessageType)
	}
	d.messageType = "partial_transcript"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingPartialTranscript) MarshalJSON() ([]byte, error) {
	type embed DubbingPartialTranscript
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "partial_transcript",
	}
	return json.Marshal(marshaler)
}
