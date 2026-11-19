package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Word-level detail of the transcription with timing information.
type DiarizedTranscriptWordsItem struct {
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
	Characters []*DiarizedTranscriptWordsItemCharactersItem `json:"characters,omitempty" xml:"characters,omitempty"`
	// The channel this word was spoken on (for multichannel audio). Null for single-channel transcriptions.
	ChannelIndex    *int `json:"channel_index,omitempty" xml:"channel_index,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DiarizedTranscriptWordsItem) GetText() (value string) {
	if d == nil {
		return
	}
	return d.Text
}

func (d *DiarizedTranscriptWordsItem) GetStart() (value *float64) {
	if d == nil {
		return
	}
	return d.Start
}

func (d *DiarizedTranscriptWordsItem) GetEnd() (value *float64) {
	if d == nil {
		return
	}
	return d.End
}

func (d *DiarizedTranscriptWordsItem) GetType() (value SpeechToTextWordResponseModelType) {
	if d == nil {
		return
	}
	return d.Type
}

func (d *DiarizedTranscriptWordsItem) GetSpeakerId() (value *string) {
	if d == nil {
		return
	}
	return d.SpeakerId
}

func (d *DiarizedTranscriptWordsItem) GetLogprob() (value float64) {
	if d == nil {
		return
	}
	return d.Logprob
}

func (d *DiarizedTranscriptWordsItem) GetCharacters() (value []*DiarizedTranscriptWordsItemCharactersItem) {
	if d == nil {
		return
	}
	return d.Characters
}

func (d *DiarizedTranscriptWordsItem) GetChannelIndex() (value *int) {
	if d == nil {
		return
	}
	return d.ChannelIndex
}

func (d *DiarizedTranscriptWordsItem) SetText(text string) {
	d.Text = text
}

func (d *DiarizedTranscriptWordsItem) SetStart(start *float64) {
	d.Start = start
}

func (d *DiarizedTranscriptWordsItem) SetEnd(end *float64) {
	d.End = end
}

func (d *DiarizedTranscriptWordsItem) SetType(type_ SpeechToTextWordResponseModelType) {
	d.Type = type_
}

func (d *DiarizedTranscriptWordsItem) SetSpeakerId(speakerId *string) {
	d.SpeakerId = speakerId
}

func (d *DiarizedTranscriptWordsItem) SetLogprob(logprob float64) {
	d.Logprob = logprob
}

func (d *DiarizedTranscriptWordsItem) SetCharacters(characters []*DiarizedTranscriptWordsItemCharactersItem) {
	d.Characters = characters
}

func (d *DiarizedTranscriptWordsItem) SetChannelIndex(channelIndex *int) {
	d.ChannelIndex = channelIndex
}

func (d *DiarizedTranscriptWordsItem) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DiarizedTranscriptWordsItem) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DiarizedTranscriptWordsItem to string"
	}
	return string(jsonData)
}

func (d *DiarizedTranscriptWordsItem) UnmarshalJSON(data []byte) error {
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

func (d *DiarizedTranscriptWordsItem) MarshalJSON() ([]byte, error) {
	type embed DiarizedTranscriptWordsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
