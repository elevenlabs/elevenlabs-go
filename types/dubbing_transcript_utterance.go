package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DubbingTranscriptUtterance struct {
	Text            *string                  `json:"text,omitempty" xml:"text,omitempty"`
	SpeakerId       *string                  `json:"speaker_id,omitempty" xml:"speaker_id,omitempty"`
	StartS          *float64                 `json:"start_s,omitempty" xml:"start_s,omitempty"`
	EndS            *float64                 `json:"end_s,omitempty" xml:"end_s,omitempty"`
	Words           []*DubbingTranscriptWord `json:"words,omitempty" xml:"words,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DubbingTranscriptUtterance) GetText() (value *string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DubbingTranscriptUtterance) GetSpeakerId() (value *string) {
	if d == nil {
		return
	}
	return d.SpeakerId
}

func (d *DubbingTranscriptUtterance) GetStartS() (value *float64) {
	if d == nil {
		return
	}
	return d.StartS
}

func (d *DubbingTranscriptUtterance) GetEndS() (value *float64) {
	if d == nil {
		return
	}
	return d.EndS
}

func (d *DubbingTranscriptUtterance) GetWords() (value []*DubbingTranscriptWord) {
	if d == nil {
		return
	}
	return d.Words
}

func (d *DubbingTranscriptUtterance) SetText(text *string) {
	d.Text = text
}

func (d *DubbingTranscriptUtterance) SetSpeakerId(speakerId *string) {
	d.SpeakerId = speakerId
}

func (d *DubbingTranscriptUtterance) SetStartS(startS *float64) {
	d.StartS = startS
}

func (d *DubbingTranscriptUtterance) SetEndS(endS *float64) {
	d.EndS = endS
}

func (d *DubbingTranscriptUtterance) SetWords(words []*DubbingTranscriptWord) {
	d.Words = words
}

func (d *DubbingTranscriptUtterance) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTranscriptUtterance) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTranscriptUtterance to string"
	}
	return string(jsonData)
}

func (d *DubbingTranscriptUtterance) UnmarshalJSON(data []byte) error {
	type unmarshaler DubbingTranscriptUtterance
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DubbingTranscriptUtterance(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingTranscriptUtterance) MarshalJSON() ([]byte, error) {
	type embed DubbingTranscriptUtterance
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
