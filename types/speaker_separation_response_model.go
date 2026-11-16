package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeakerSeparationResponseModel struct {
	// The ID of the voice.
	VoiceId string `json:"voice_id" xml:"voice_id" required:"true"`
	// The ID of the sample.
	SampleId string `json:"sample_id" xml:"sample_id" required:"true"`
	// The status of the speaker separation.
	Status SpeakerSeparationResponseModelStatus `json:"status" xml:"status" required:"true"`
	// The speakers of the sample.
	Speakers map[string]*SpeakerResponseModel `json:"speakers,omitempty"`
	// The IDs of the selected speakers.
	SelectedSpeakerIds []string `json:"selected_speaker_ids,omitempty" xml:"selected_speaker_ids,omitempty"`
	extraProperties    map[string]interface{}
}

func (s *SpeakerSeparationResponseModel) GetVoiceId() (value string) {
	if s == nil {
		return
	}
	return s.VoiceId
}

func (s *SpeakerSeparationResponseModel) GetSampleId() (value string) {
	if s == nil {
		return
	}
	return s.SampleId
}

func (s *SpeakerSeparationResponseModel) GetStatus() (value SpeakerSeparationResponseModelStatus) {
	if s == nil {
		return
	}
	return s.Status
}

func (s *SpeakerSeparationResponseModel) GetSpeakers() (value map[string]*SpeakerResponseModel) {
	if s == nil {
		return
	}
	return s.Speakers
}

func (s *SpeakerSeparationResponseModel) GetSelectedSpeakerIds() (value []string) {
	if s == nil {
		return
	}
	return s.SelectedSpeakerIds
}

func (s *SpeakerSeparationResponseModel) SetVoiceId(voiceId string) {
	s.VoiceId = voiceId
}

func (s *SpeakerSeparationResponseModel) SetSampleId(sampleId string) {
	s.SampleId = sampleId
}

func (s *SpeakerSeparationResponseModel) SetStatus(status SpeakerSeparationResponseModelStatus) {
	s.Status = status
}

func (s *SpeakerSeparationResponseModel) SetSpeakers(speakers map[string]*SpeakerResponseModel) {
	s.Speakers = speakers
}

func (s *SpeakerSeparationResponseModel) SetSelectedSpeakerIds(selectedSpeakerIds []string) {
	s.SelectedSpeakerIds = selectedSpeakerIds
}

func (s *SpeakerSeparationResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeakerSeparationResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeakerSeparationResponseModel to string"
	}
	return string(jsonData)
}

func (s *SpeakerSeparationResponseModel) UnmarshalJSON(data []byte) error {
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

func (s *SpeakerSeparationResponseModel) MarshalJSON() ([]byte, error) {
	type embed SpeakerSeparationResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
