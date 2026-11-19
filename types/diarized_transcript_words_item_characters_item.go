package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DiarizedTranscriptWordsItemCharactersItem struct {
	// The character that was transcribed.
	Text string `json:"text" xml:"text" required:"true"`
	// The start time of the character in seconds.
	Start *float64 `json:"start,omitempty" xml:"start,omitempty"`
	// The end time of the character in seconds.
	End             *float64 `json:"end,omitempty" xml:"end,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DiarizedTranscriptWordsItemCharactersItem) GetText() (value string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DiarizedTranscriptWordsItemCharactersItem) GetStart() (value *float64) {
	if d == nil {
		return
	}
	return d.Start
}

func (d *DiarizedTranscriptWordsItemCharactersItem) GetEnd() (value *float64) {
	if d == nil {
		return
	}
	return d.End
}

func (d *DiarizedTranscriptWordsItemCharactersItem) SetText(text string) {
	d.Text = text
}

func (d *DiarizedTranscriptWordsItemCharactersItem) SetStart(start *float64) {
	d.Start = start
}

func (d *DiarizedTranscriptWordsItemCharactersItem) SetEnd(end *float64) {
	d.End = end
}

func (d *DiarizedTranscriptWordsItemCharactersItem) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DiarizedTranscriptWordsItemCharactersItem) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DiarizedTranscriptWordsItemCharactersItem to string"
	}
	return string(jsonData)
}

func (d *DiarizedTranscriptWordsItemCharactersItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DiarizedTranscriptWordsItemCharactersItem) MarshalJSON() ([]byte, error) {
	type embed DiarizedTranscriptWordsItemCharactersItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
