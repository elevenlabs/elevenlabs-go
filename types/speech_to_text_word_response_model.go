package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Word-level detail of the transcription with timing information.
type SpeechToTextWordResponseModel struct {
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
	Characters []*SpeechToTextCharacterResponseModel `json:"characters,omitempty" xml:"characters,omitempty"`
	// The channel this word was spoken on (for multichannel audio). Null for single-channel transcriptions.
	ChannelIndex    *int `json:"channel_index,omitempty" xml:"channel_index,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SpeechToTextWordResponseModel) GetText() (value string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SpeechToTextWordResponseModel) GetStart() (value *float64) {
	if s == nil {
		return
	}
	return s.Start
}

func (s *SpeechToTextWordResponseModel) GetEnd() (value *float64) {
	if s == nil {
		return
	}
	return s.End
}

func (s *SpeechToTextWordResponseModel) GetType() (value SpeechToTextWordResponseModelType) {
	if s == nil {
		return
	}
	return s.Type
}

func (s *SpeechToTextWordResponseModel) GetSpeakerId() (value *string) {
	if s == nil {
		return
	}
	return s.SpeakerId
}

func (s *SpeechToTextWordResponseModel) GetLogprob() (value float64) {
	if s == nil {
		return
	}
	return s.Logprob
}

func (s *SpeechToTextWordResponseModel) GetCharacters() (value []*SpeechToTextCharacterResponseModel) {
	if s == nil {
		return
	}
	return s.Characters
}

func (s *SpeechToTextWordResponseModel) GetChannelIndex() (value *int) {
	if s == nil {
		return
	}
	return s.ChannelIndex
}

func (s *SpeechToTextWordResponseModel) SetText(text string) {
	s.Text = text
}

func (s *SpeechToTextWordResponseModel) SetStart(start *float64) {
	s.Start = start
}

func (s *SpeechToTextWordResponseModel) SetEnd(end *float64) {
	s.End = end
}

func (s *SpeechToTextWordResponseModel) SetType(type_ SpeechToTextWordResponseModelType) {
	s.Type = type_
}

func (s *SpeechToTextWordResponseModel) SetSpeakerId(speakerId *string) {
	s.SpeakerId = speakerId
}

func (s *SpeechToTextWordResponseModel) SetLogprob(logprob float64) {
	s.Logprob = logprob
}

func (s *SpeechToTextWordResponseModel) SetCharacters(characters []*SpeechToTextCharacterResponseModel) {
	s.Characters = characters
}

func (s *SpeechToTextWordResponseModel) SetChannelIndex(channelIndex *int) {
	s.ChannelIndex = channelIndex
}

func (s *SpeechToTextWordResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeechToTextWordResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeechToTextWordResponseModel to string"
	}
	return string(jsonData)
}

func (s *SpeechToTextWordResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeechToTextWordResponseModel) MarshalJSON() ([]byte, error) {
	type embed SpeechToTextWordResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
