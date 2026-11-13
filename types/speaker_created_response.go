package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeakerCreatedResponse struct {
	Version         int    `json:"version" xml:"version" required:"true"`
	SpeakerId       string `json:"speaker_id" xml:"speaker_id" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SpeakerCreatedResponse) GetVersion() (value int) {
	if s == nil {
		return
	}
	return s.Version
}

func (s *SpeakerCreatedResponse) GetSpeakerId() (value string) {
	if s == nil {
		return
	}
	return s.SpeakerId
}

func (s *SpeakerCreatedResponse) SetVersion(version int) {
	s.Version = version
}

func (s *SpeakerCreatedResponse) SetSpeakerId(speakerId string) {
	s.SpeakerId = speakerId
}

func (s *SpeakerCreatedResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeakerCreatedResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeakerCreatedResponse to string"
	}
	return string(jsonData)
}

func (s *SpeakerCreatedResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SpeakerCreatedResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SpeakerCreatedResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeakerCreatedResponse) MarshalJSON() ([]byte, error) {
	type embed SpeakerCreatedResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
