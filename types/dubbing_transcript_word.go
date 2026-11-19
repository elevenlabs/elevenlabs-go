package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type DubbingTranscriptWord struct {
	Text            *string                       `json:"text,omitempty" xml:"text,omitempty"`
	WordType        *string                       `json:"word_type,omitempty" xml:"word_type,omitempty"`
	StartS          *float64                      `json:"start_s,omitempty" xml:"start_s,omitempty"`
	EndS            *float64                      `json:"end_s,omitempty" xml:"end_s,omitempty"`
	Characters      []*DubbingTranscriptCharacter `json:"characters,omitempty" xml:"characters,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DubbingTranscriptWord) GetText() (value *string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DubbingTranscriptWord) GetWordType() (value *string) {
	if d == nil {
		return
	}
	return d.WordType
}

func (d *DubbingTranscriptWord) GetStartS() (value *float64) {
	if d == nil {
		return
	}
	return d.StartS
}

func (d *DubbingTranscriptWord) GetEndS() (value *float64) {
	if d == nil {
		return
	}
	return d.EndS
}

func (d *DubbingTranscriptWord) GetCharacters() (value []*DubbingTranscriptCharacter) {
	if d == nil {
		return
	}
	return d.Characters
}

func (d *DubbingTranscriptWord) SetText(text *string) {
	d.Text = text
}

func (d *DubbingTranscriptWord) SetWordType(wordType *string) {
	d.WordType = wordType
}

func (d *DubbingTranscriptWord) SetStartS(startS *float64) {
	d.StartS = startS
}

func (d *DubbingTranscriptWord) SetEndS(endS *float64) {
	d.EndS = endS
}

func (d *DubbingTranscriptWord) SetCharacters(characters []*DubbingTranscriptCharacter) {
	d.Characters = characters
}

func (d *DubbingTranscriptWord) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingTranscriptWord) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingTranscriptWord to string"
	}
	return string(jsonData)
}

func (d *DubbingTranscriptWord) UnmarshalJSON(data []byte) error {
	type unmarshaler DubbingTranscriptWord
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DubbingTranscriptWord(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingTranscriptWord) MarshalJSON() ([]byte, error) {
	type embed DubbingTranscriptWord
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
