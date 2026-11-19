package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Word-level detail of the transcription with timing information.
type PartialTranscriptWordsItem struct {
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
	Characters []*PartialTranscriptWordsItemCharactersItem `json:"characters,omitempty" xml:"characters,omitempty"`
	// The channel this word was spoken on (for multichannel audio). Null for single-channel transcriptions.
	ChannelIndex    *int `json:"channel_index,omitempty" xml:"channel_index,omitempty"`
	extraProperties map[string]interface{}
}

func (p *PartialTranscriptWordsItem) GetText() (value string) {
	if p == nil {
		return
	}
	return p.Text
}

func (p *PartialTranscriptWordsItem) GetStart() (value *float64) {
	if p == nil {
		return
	}
	return p.Start
}

func (p *PartialTranscriptWordsItem) GetEnd() (value *float64) {
	if p == nil {
		return
	}
	return p.End
}

func (p *PartialTranscriptWordsItem) GetType() (value SpeechToTextWordResponseModelType) {
	if p == nil {
		return
	}
	return p.Type
}

func (p *PartialTranscriptWordsItem) GetSpeakerId() (value *string) {
	if p == nil {
		return
	}
	return p.SpeakerId
}

func (p *PartialTranscriptWordsItem) GetLogprob() (value float64) {
	if p == nil {
		return
	}
	return p.Logprob
}

func (p *PartialTranscriptWordsItem) GetCharacters() (value []*PartialTranscriptWordsItemCharactersItem) {
	if p == nil {
		return
	}
	return p.Characters
}

func (p *PartialTranscriptWordsItem) GetChannelIndex() (value *int) {
	if p == nil {
		return
	}
	return p.ChannelIndex
}

func (p *PartialTranscriptWordsItem) SetText(text string) {
	p.Text = text
}

func (p *PartialTranscriptWordsItem) SetStart(start *float64) {
	p.Start = start
}

func (p *PartialTranscriptWordsItem) SetEnd(end *float64) {
	p.End = end
}

func (p *PartialTranscriptWordsItem) SetType(type_ SpeechToTextWordResponseModelType) {
	p.Type = type_
}

func (p *PartialTranscriptWordsItem) SetSpeakerId(speakerId *string) {
	p.SpeakerId = speakerId
}

func (p *PartialTranscriptWordsItem) SetLogprob(logprob float64) {
	p.Logprob = logprob
}

func (p *PartialTranscriptWordsItem) SetCharacters(characters []*PartialTranscriptWordsItemCharactersItem) {
	p.Characters = characters
}

func (p *PartialTranscriptWordsItem) SetChannelIndex(channelIndex *int) {
	p.ChannelIndex = channelIndex
}

func (p *PartialTranscriptWordsItem) GetExtraProperties() map[string]interface{} {
	if p == nil {
		return nil
	}
	return p.extraProperties
}

func (p PartialTranscriptWordsItem) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PartialTranscriptWordsItem to string"
	}
	return string(jsonData)
}

func (p *PartialTranscriptWordsItem) UnmarshalJSON(data []byte) error {
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

func (p *PartialTranscriptWordsItem) MarshalJSON() ([]byte, error) {
	type embed PartialTranscriptWordsItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
