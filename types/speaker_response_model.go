package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeakerResponseModel struct {
	// The ID of the speaker.
	SpeakerId string `json:"speaker_id" xml:"speaker_id" required:"true"`
	// The duration of the speaker segment in seconds.
	DurationSecs float64 `json:"duration_secs" xml:"duration_secs" required:"true"`
	// The utterances of the speaker.
	Utterances      []*UtteranceResponseModel `json:"utterances,omitempty" xml:"utterances,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SpeakerResponseModel) GetSpeakerId() (value string) {
	if s == nil {
		return
	}
	return s.SpeakerId
}

func (s *SpeakerResponseModel) GetDurationSecs() (value float64) {
	if s == nil {
		return
	}
	return s.DurationSecs
}

func (s *SpeakerResponseModel) GetUtterances() (value []*UtteranceResponseModel) {
	if s == nil {
		return
	}
	return s.Utterances
}

func (s *SpeakerResponseModel) SetSpeakerId(speakerId string) {
	s.SpeakerId = speakerId
}

func (s *SpeakerResponseModel) SetDurationSecs(durationSecs float64) {
	s.DurationSecs = durationSecs
}

func (s *SpeakerResponseModel) SetUtterances(utterances []*UtteranceResponseModel) {
	s.Utterances = utterances
}

func (s *SpeakerResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeakerResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeakerResponseModel to string"
	}
	return string(jsonData)
}

func (s *SpeakerResponseModel) UnmarshalJSON(data []byte) error {
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

func (s *SpeakerResponseModel) MarshalJSON() ([]byte, error) {
	type embed SpeakerResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
