package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Word-level detail of the transcription with timing information.
type CommittedTranscriptWithTimestampsWordsItem struct {
	// The word or sound that was transcribed.
	Text string `json:"text" xml:"text" required:"true"`
	// The start time of the word or sound in seconds.
	Start *float64 `json:"start,omitempty" xml:"start,omitempty"`
	// The end time of the word or sound in seconds.
	End *float64 `json:"end,omitempty" xml:"end,omitempty"`
	// The type of the word or sound. 'audio_event' is used for non-word sounds like laughter or footsteps.
	Type SpeechToTextWordResponseModelType `json:"type" xml:"type" required:"true"`
	// Unique identifier for the speaker of this word.
	SpeakerId *string `json:"speaker_id,omitempty" xml:"speaker_id,omitempty"`
	// The log of the probability with which this word was predicted. Logprobs are in range [-infinity, 0], higher logprobs indicate a higher confidence the model has in its predictions.
	Logprob float64 `json:"logprob" xml:"logprob" required:"true"`
	// The characters that make up the word and their timing information.
	Characters []*CommittedTranscriptWithTimestampsWordsItemCharactersItem `json:"characters,omitempty" xml:"characters,omitempty"`
	// The channel this word was spoken on (for multichannel audio). Null for single-channel transcriptions.
	ChannelIndex    *int `json:"channel_index,omitempty" xml:"channel_index,omitempty"`
	extraProperties map[string]interface{}
}

func (c *CommittedTranscriptWithTimestampsWordsItem) GetText() (value string) {
	if c == nil {
		return
	}
	return c.Text
}

func (c *CommittedTranscriptWithTimestampsWordsItem) GetStart() (value *float64) {
	if c == nil {
		return
	}
	return c.Start
}

func (c *CommittedTranscriptWithTimestampsWordsItem) GetEnd() (value *float64) {
	if c == nil {
		return
	}
	return c.End
}

func (c *CommittedTranscriptWithTimestampsWordsItem) GetType() (value SpeechToTextWordResponseModelType) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *CommittedTranscriptWithTimestampsWordsItem) GetSpeakerId() (value *string) {
	if c == nil {
		return
	}
	return c.SpeakerId
}

func (c *CommittedTranscriptWithTimestampsWordsItem) GetLogprob() (value float64) {
	if c == nil {
		return
	}
	return c.Logprob
}

func (c *CommittedTranscriptWithTimestampsWordsItem) GetCharacters() (value []*CommittedTranscriptWithTimestampsWordsItemCharactersItem) {
	if c == nil {
		return
	}
	return c.Characters
}

func (c *CommittedTranscriptWithTimestampsWordsItem) GetChannelIndex() (value *int) {
	if c == nil {
		return
	}
	return c.ChannelIndex
}

func (c *CommittedTranscriptWithTimestampsWordsItem) SetText(text string) {
	c.Text = text
}

func (c *CommittedTranscriptWithTimestampsWordsItem) SetStart(start *float64) {
	c.Start = start
}

func (c *CommittedTranscriptWithTimestampsWordsItem) SetEnd(end *float64) {
	c.End = end
}

func (c *CommittedTranscriptWithTimestampsWordsItem) SetType(type_ SpeechToTextWordResponseModelType) {
	c.Type = type_
}

func (c *CommittedTranscriptWithTimestampsWordsItem) SetSpeakerId(speakerId *string) {
	c.SpeakerId = speakerId
}

func (c *CommittedTranscriptWithTimestampsWordsItem) SetLogprob(logprob float64) {
	c.Logprob = logprob
}

func (c *CommittedTranscriptWithTimestampsWordsItem) SetCharacters(characters []*CommittedTranscriptWithTimestampsWordsItemCharactersItem) {
	c.Characters = characters
}

func (c *CommittedTranscriptWithTimestampsWordsItem) SetChannelIndex(channelIndex *int) {
	c.ChannelIndex = channelIndex
}

func (c *CommittedTranscriptWithTimestampsWordsItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CommittedTranscriptWithTimestampsWordsItem) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CommittedTranscriptWithTimestampsWordsItem to string"
	}
	return string(jsonData)
}

func (c *CommittedTranscriptWithTimestampsWordsItem) UnmarshalJSON(data []byte) error {
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

func (c *CommittedTranscriptWithTimestampsWordsItem) MarshalJSON() ([]byte, error) {
	type embed CommittedTranscriptWithTimestampsWordsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
