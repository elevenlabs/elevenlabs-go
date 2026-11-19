package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type SimilarVoicesForSpeakerResponse struct {
	Voices          []*SimilarVoice `json:"voices" xml:"voices" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SimilarVoicesForSpeakerResponse) GetVoices() (value []*SimilarVoice) {
	if s == nil {
		return
	}
	return s.Voices
}

func (s *SimilarVoicesForSpeakerResponse) SetVoices(voices []*SimilarVoice) {
	s.Voices = voices
}

func (s *SimilarVoicesForSpeakerResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SimilarVoicesForSpeakerResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SimilarVoicesForSpeakerResponse to string"
	}
	return string(jsonData)
}

func (s *SimilarVoicesForSpeakerResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SimilarVoicesForSpeakerResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SimilarVoicesForSpeakerResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SimilarVoicesForSpeakerResponse) MarshalJSON() ([]byte, error) {
	type embed SimilarVoicesForSpeakerResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
