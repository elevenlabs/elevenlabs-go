package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingCommittedTranscript struct {
	messageType     string
	Text            string  `json:"text" xml:"text" required:"true" minLength:"1"`
	SegmentId       string  `json:"segment_id" xml:"segment_id" required:"true" minLength:"1"`
	StartS          float64 `json:"start_s" xml:"start_s" required:"true" min:"0"`
	EndS            float64 `json:"end_s" xml:"end_s" required:"true" min:"0"`
	SpeakerId       string  `json:"speaker_id" xml:"speaker_id" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DubbingCommittedTranscript) GetText() (value string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DubbingCommittedTranscript) GetSegmentId() (value string) {
	if d == nil {
		return
	}
	return d.SegmentId
}

func (d *DubbingCommittedTranscript) GetStartS() (value float64) {
	if d == nil {
		return
	}
	return d.StartS
}

func (d *DubbingCommittedTranscript) GetEndS() (value float64) {
	if d == nil {
		return
	}
	return d.EndS
}

func (d *DubbingCommittedTranscript) GetSpeakerId() (value string) {
	if d == nil {
		return
	}
	return d.SpeakerId
}

func (d *DubbingCommittedTranscript) SetText(text string) {
	d.Text = text
}

func (d *DubbingCommittedTranscript) SetSegmentId(segmentId string) {
	d.SegmentId = segmentId
}

func (d *DubbingCommittedTranscript) SetStartS(startS float64) {
	d.StartS = startS
}

func (d *DubbingCommittedTranscript) SetEndS(endS float64) {
	d.EndS = endS
}

func (d *DubbingCommittedTranscript) SetSpeakerId(speakerId string) {
	d.SpeakerId = speakerId
}

func (d *DubbingCommittedTranscript) MessageType() string {
	return d.messageType
}

func (d *DubbingCommittedTranscript) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingCommittedTranscript) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingCommittedTranscript to string"
	}
	return string(jsonData)
}

func (d *DubbingCommittedTranscript) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingCommittedTranscript
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingCommittedTranscript(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "committed_transcript" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "committed_transcript", *literals.MessageType)
	}
	d.messageType = "committed_transcript"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingCommittedTranscript) MarshalJSON() ([]byte, error) {
	type embed DubbingCommittedTranscript
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "committed_transcript",
	}
	return json.Marshal(marshaler)
}
