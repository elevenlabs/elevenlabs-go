package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SimilarVoice struct {
	VoiceId         string        `json:"voice_id" xml:"voice_id" required:"true"`
	Name            string        `json:"name" xml:"name" required:"true"`
	Category        VoiceCategory `json:"category" xml:"category" required:"true"`
	Description     *string       `json:"description,omitempty" xml:"description,omitempty"`
	PreviewUrl      *string       `json:"preview_url,omitempty" xml:"preview_url,omitempty"`
	extraProperties map[string]interface{}
}

func (s *SimilarVoice) GetVoiceId() (value string) {
	if s == nil {
		return
	}
	return s.VoiceId
}

func (s *SimilarVoice) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

func (s *SimilarVoice) GetCategory() (value VoiceCategory) {
	if s == nil {
		return
	}
	return s.Category
}

func (s *SimilarVoice) GetDescription() (value *string) {
	if s == nil {
		return
	}
	return s.Description
}

func (s *SimilarVoice) GetPreviewUrl() (value *string) {
	if s == nil {
		return
	}
	return s.PreviewUrl
}

func (s *SimilarVoice) SetVoiceId(voiceId string) {
	s.VoiceId = voiceId
}

func (s *SimilarVoice) SetName(name string) {
	s.Name = name
}

func (s *SimilarVoice) SetCategory(category VoiceCategory) {
	s.Category = category
}

func (s *SimilarVoice) SetDescription(description *string) {
	s.Description = description
}

func (s *SimilarVoice) SetPreviewUrl(previewUrl *string) {
	s.PreviewUrl = previewUrl
}

func (s *SimilarVoice) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SimilarVoice) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SimilarVoice to string"
	}
	return string(jsonData)
}

func (s *SimilarVoice) UnmarshalJSON(data []byte) error {
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

func (s *SimilarVoice) MarshalJSON() ([]byte, error) {
	type embed SimilarVoice
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
