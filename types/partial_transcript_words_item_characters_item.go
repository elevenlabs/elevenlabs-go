package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type PartialTranscriptWordsItemCharactersItem struct {
	// The character that was transcribed.
	Text string `json:"text" xml:"text" required:"true"`
	// The start time of the character in seconds.
	Start *float64 `json:"start,omitempty" xml:"start,omitempty"`
	// The end time of the character in seconds.
	End             *float64 `json:"end,omitempty" xml:"end,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PartialTranscriptWordsItemCharactersItem) GetText() (value string) {
	if p == nil {
		return
	}
	return p.Text
}

func (p *PartialTranscriptWordsItemCharactersItem) GetStart() (value *float64) {
	if p == nil {
		return
	}
	return p.Start
}

func (p *PartialTranscriptWordsItemCharactersItem) GetEnd() (value *float64) {
	if p == nil {
		return
	}
	return p.End
}

func (p *PartialTranscriptWordsItemCharactersItem) SetText(text string) {
	p.Text = text
}

func (p *PartialTranscriptWordsItemCharactersItem) SetStart(start *float64) {
	p.Start = start
}

func (p *PartialTranscriptWordsItemCharactersItem) SetEnd(end *float64) {
	p.End = end
}

func (p *PartialTranscriptWordsItemCharactersItem) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PartialTranscriptWordsItemCharactersItem) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PartialTranscriptWordsItemCharactersItem to string"
	}
	return string(jsonData)
}

func (p *PartialTranscriptWordsItemCharactersItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, p); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties
	return nil
}

func (p *PartialTranscriptWordsItemCharactersItem) MarshalJSON() ([]byte, error) {
	type embed PartialTranscriptWordsItemCharactersItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
