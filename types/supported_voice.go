package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SupportedVoice struct {
	Label                    string          `json:"label" xml:"label" required:"true" minLength:"1"`
	VoiceId                  string          `json:"voice_id" xml:"voice_id" required:"true" minLength:"1"`
	Description              *string         `json:"description,omitempty" xml:"description,omitempty"`
	Language                 *string         `json:"language,omitempty" xml:"language,omitempty"`
	ModelFamily              *TtsModelFamily `json:"model_family,omitempty" xml:"model_family,omitempty"`
	OptimizeStreamingLatency *int            `json:"optimize_streaming_latency,omitempty" xml:"optimize_streaming_latency,omitempty"`
	Stability                *float64        `json:"stability,omitempty" xml:"stability,omitempty" min:"0" max:"1"`
	Speed                    *float64        `json:"speed,omitempty" xml:"speed,omitempty" min:"0.7" max:"1.2"`
	SimilarityBoost          *float64        `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty" min:"0" max:"1"`
	CloneStrength            *float64        `json:"clone_strength,omitempty" xml:"clone_strength,omitempty" min:"0" max:"1"`
	extraProperties          map[string]interface{}
}

func (s *SupportedVoice) GetLabel() (value string) {
	if s == nil {
		return
	}
	return s.Label
}

func (s *SupportedVoice) GetVoiceId() (value string) {
	if s == nil {
		return
	}
	return s.VoiceId
}

func (s *SupportedVoice) GetDescription() (value *string) {
	if s == nil {
		return
	}
	return s.Description
}

func (s *SupportedVoice) GetLanguage() (value *string) {
	if s == nil {
		return
	}
	return s.Language
}

func (s *SupportedVoice) GetModelFamily() (value *TtsModelFamily) {
	if s == nil {
		return
	}
	return s.ModelFamily
}

func (s *SupportedVoice) GetOptimizeStreamingLatency() (value *int) {
	if s == nil {
		return
	}
	return s.OptimizeStreamingLatency
}

func (s *SupportedVoice) GetStability() (value *float64) {
	if s == nil {
		return
	}
	return s.Stability
}

func (s *SupportedVoice) GetSpeed() (value *float64) {
	if s == nil {
		return
	}
	return s.Speed
}

func (s *SupportedVoice) GetSimilarityBoost() (value *float64) {
	if s == nil {
		return
	}
	return s.SimilarityBoost
}

func (s *SupportedVoice) GetCloneStrength() (value *float64) {
	if s == nil {
		return
	}
	return s.CloneStrength
}

func (s *SupportedVoice) SetLabel(label string) {
	s.Label = label
}

func (s *SupportedVoice) SetVoiceId(voiceId string) {
	s.VoiceId = voiceId
}

func (s *SupportedVoice) SetDescription(description *string) {
	s.Description = description
}

func (s *SupportedVoice) SetLanguage(language *string) {
	s.Language = language
}

func (s *SupportedVoice) SetModelFamily(modelFamily *TtsModelFamily) {
	s.ModelFamily = modelFamily
}

func (s *SupportedVoice) SetOptimizeStreamingLatency(optimizeStreamingLatency *int) {
	s.OptimizeStreamingLatency = optimizeStreamingLatency
}

func (s *SupportedVoice) SetStability(stability *float64) {
	s.Stability = stability
}

func (s *SupportedVoice) SetSpeed(speed *float64) {
	s.Speed = speed
}

func (s *SupportedVoice) SetSimilarityBoost(similarityBoost *float64) {
	s.SimilarityBoost = similarityBoost
}

func (s *SupportedVoice) SetCloneStrength(cloneStrength *float64) {
	s.CloneStrength = cloneStrength
}

func (s *SupportedVoice) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SupportedVoice) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SupportedVoice to string"
	}
	return string(jsonData)
}

func (s *SupportedVoice) UnmarshalJSON(data []byte) error {
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

func (s *SupportedVoice) MarshalJSON() ([]byte, error) {
	type embed SupportedVoice
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
