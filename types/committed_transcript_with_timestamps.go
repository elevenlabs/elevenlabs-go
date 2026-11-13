package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type CommittedTranscriptWithTimestamps struct {
	MessageType     *string                                       `json:"message_type,omitempty" xml:"message_type,omitempty"`
	Text            string                                        `json:"text" xml:"text" required:"true"`
	LanguageCode    *string                                       `json:"language_code,omitempty" xml:"language_code,omitempty"`
	Words           []*CommittedTranscriptWithTimestampsWordsItem `json:"words,omitempty" xml:"words,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CommittedTranscriptWithTimestamps) GetMessageType() (value *string) {
	if c == nil {
		return
	}
	return c.MessageType
}

func (c *CommittedTranscriptWithTimestamps) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

func (c *CommittedTranscriptWithTimestamps) GetLanguageCode() (value *string) {
	if c == nil {
		return
	}
	return c.LanguageCode
}

func (c *CommittedTranscriptWithTimestamps) GetWords() (value []*CommittedTranscriptWithTimestampsWordsItem) {
	if c == nil {
		return
	}
	return c.Words
}

func (c *CommittedTranscriptWithTimestamps) SetMessageType(messageType *string) {
	c.MessageType = messageType
}

func (c *CommittedTranscriptWithTimestamps) SetText(text string) {
	c.Text = text
}

func (c *CommittedTranscriptWithTimestamps) SetLanguageCode(languageCode *string) {
	c.LanguageCode = languageCode
}

func (c *CommittedTranscriptWithTimestamps) SetWords(words []*CommittedTranscriptWithTimestampsWordsItem) {
	c.Words = words
}

func (c *CommittedTranscriptWithTimestamps) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CommittedTranscriptWithTimestamps) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CommittedTranscriptWithTimestamps to string"
	}
	return string(jsonData)
}

func (c *CommittedTranscriptWithTimestamps) UnmarshalJSON(data []byte) error {
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

func (c *CommittedTranscriptWithTimestamps) MarshalJSON() ([]byte, error) {
	type embed CommittedTranscriptWithTimestamps
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
