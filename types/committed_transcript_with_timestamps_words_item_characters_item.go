package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CommittedTranscriptWithTimestampsWordsItemCharactersItem struct {
	// The character that was transcribed.
	Text string `json:"text" xml:"text" required:"true"`
	// The start time of the character in seconds.
	Start *float64 `json:"start,omitempty" xml:"start,omitempty"`
	// The end time of the character in seconds.
	End             *float64 `json:"end,omitempty" xml:"end,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CommittedTranscriptWithTimestampsWordsItemCharactersItem) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

func (c *CommittedTranscriptWithTimestampsWordsItemCharactersItem) GetStart() (value *float64) {
	if c == nil {
		return
	}
	return c.Start
}

func (c *CommittedTranscriptWithTimestampsWordsItemCharactersItem) GetEnd() (value *float64) {
	if c == nil {
		return
	}
	return c.End
}

func (c *CommittedTranscriptWithTimestampsWordsItemCharactersItem) SetText(text string) {
	c.Text = text
}

func (c *CommittedTranscriptWithTimestampsWordsItemCharactersItem) SetStart(start *float64) {
	c.Start = start
}

func (c *CommittedTranscriptWithTimestampsWordsItemCharactersItem) SetEnd(end *float64) {
	c.End = end
}

func (c *CommittedTranscriptWithTimestampsWordsItemCharactersItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CommittedTranscriptWithTimestampsWordsItemCharactersItem) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CommittedTranscriptWithTimestampsWordsItemCharactersItem to string"
	}
	return string(jsonData)
}

func (c *CommittedTranscriptWithTimestampsWordsItemCharactersItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CommittedTranscriptWithTimestampsWordsItemCharactersItem) MarshalJSON() ([]byte, error) {
	type embed CommittedTranscriptWithTimestampsWordsItemCharactersItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
