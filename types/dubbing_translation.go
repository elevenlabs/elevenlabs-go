package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingTranslation struct {
	messageType     string
	Text            string  `json:"text" xml:"text" required:"true" minLength:"1"`
	SegmentId       string  `json:"segment_id" xml:"segment_id" required:"true" minLength:"1"`
	StartS          float64 `json:"start_s" xml:"start_s" required:"true" min:"0"`
	EndS            float64 `json:"end_s" xml:"end_s" required:"true" min:"0"`
	SpeakerId       string  `json:"speaker_id" xml:"speaker_id" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DubbingTranslation) GetText() (value string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DubbingTranslation) GetSegmentId() (value string) {
	if d == nil {
		return
	}
	return d.SegmentId
}

func (d *DubbingTranslation) GetStartS() (value float64) {
	if d == nil {
		return
	}
	return d.StartS
}

func (d *DubbingTranslation) GetEndS() (value float64) {
	if d == nil {
		return
	}
	return d.EndS
}

func (d *DubbingTranslation) GetSpeakerId() (value string) {
	if d == nil {
		return
	}
	return d.SpeakerId
}

func (d *DubbingTranslation) SetText(text string) {
	d.Text = text
}

func (d *DubbingTranslation) SetSegmentId(segmentId string) {
	d.SegmentId = segmentId
}

func (d *DubbingTranslation) SetStartS(startS float64) {
	d.StartS = startS
}

func (d *DubbingTranslation) SetEndS(endS float64) {
	d.EndS = endS
}

func (d *DubbingTranslation) SetSpeakerId(speakerId string) {
	d.SpeakerId = speakerId
}

func (d *DubbingTranslation) MessageType() string {
	return d.messageType
}

func (d *DubbingTranslation) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTranslation) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTranslation to string"
	}
	return string(jsonData)
}

func (d *DubbingTranslation) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingTranslation
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingTranslation(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "translation" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "translation", *literals.MessageType)
	}
	d.messageType = "translation"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingTranslation) MarshalJSON() ([]byte, error) {
	type embed DubbingTranslation
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "translation",
	}
	return json.Marshal(marshaler)
}
