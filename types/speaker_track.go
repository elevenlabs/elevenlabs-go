package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeakerTrack struct {
	Id              string                 `json:"id" xml:"id" required:"true"`
	MediaRef        *DubbingMediaReference `json:"media_ref" xml:"media_ref" required:"true"`
	SpeakerName     string                 `json:"speaker_name" xml:"speaker_name" required:"true"`
	Voices          map[string]string      `json:"voices" required:"true"`
	Segments        []string               `json:"segments" xml:"segments" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SpeakerTrack) GetId() (value string) {
	if s == nil {
		return
	}
	return s.Id
}

func (s *SpeakerTrack) GetMediaRef() (value *DubbingMediaReference) {
	if s == nil {
		return
	}
	return s.MediaRef
}

func (s *SpeakerTrack) GetSpeakerName() (value string) {
	if s == nil {
		return
	}
	return s.SpeakerName
}

func (s *SpeakerTrack) GetVoices() (value map[string]string) {
	if s == nil {
		return
	}
	return s.Voices
}

func (s *SpeakerTrack) GetSegments() (value []string) {
	if s == nil {
		return
	}
	return s.Segments
}

func (s *SpeakerTrack) SetId(id string) {
	s.Id = id
}

func (s *SpeakerTrack) SetMediaRef(mediaRef *DubbingMediaReference) {
	s.MediaRef = mediaRef
}

func (s *SpeakerTrack) SetSpeakerName(speakerName string) {
	s.SpeakerName = speakerName
}

func (s *SpeakerTrack) SetVoices(voices map[string]string) {
	s.Voices = voices
}

func (s *SpeakerTrack) SetSegments(segments []string) {
	s.Segments = segments
}

func (s *SpeakerTrack) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeakerTrack) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeakerTrack to string"
	}
	return string(jsonData)
}

func (s *SpeakerTrack) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SpeakerTrack
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SpeakerTrack(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeakerTrack) MarshalJSON() ([]byte, error) {
	type embed SpeakerTrack
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
