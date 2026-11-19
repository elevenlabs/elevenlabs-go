package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DubbingTranscriptCharacter struct {
	Text            *string  `json:"text,omitempty" xml:"text,omitempty"`
	StartS          *float64 `json:"start_s,omitempty" xml:"start_s,omitempty"`
	EndS            *float64 `json:"end_s,omitempty" xml:"end_s,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DubbingTranscriptCharacter) GetText() (value *string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DubbingTranscriptCharacter) GetStartS() (value *float64) {
	if d == nil {
		return
	}
	return d.StartS
}

func (d *DubbingTranscriptCharacter) GetEndS() (value *float64) {
	if d == nil {
		return
	}
	return d.EndS
}

func (d *DubbingTranscriptCharacter) SetText(text *string) {
	d.Text = text
}

func (d *DubbingTranscriptCharacter) SetStartS(startS *float64) {
	d.StartS = startS
}

func (d *DubbingTranscriptCharacter) SetEndS(endS *float64) {
	d.EndS = endS
}

func (d *DubbingTranscriptCharacter) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTranscriptCharacter) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTranscriptCharacter to string"
	}
	return string(jsonData)
}

func (d *DubbingTranscriptCharacter) UnmarshalJSON(data []byte) error {
	type unmarshaler DubbingTranscriptCharacter
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DubbingTranscriptCharacter(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingTranscriptCharacter) MarshalJSON() ([]byte, error) {
	type embed DubbingTranscriptCharacter
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
