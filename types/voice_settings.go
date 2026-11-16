package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type VoiceSettings struct {
	// Determines how stable the voice is and the randomness between each generation. Lower values introduce broader emotional range for the voice. Higher values can result in a monotonous voice with limited emotion.
	Stability *float64 `json:"stability,omitempty" xml:"stability,omitempty" min:"0" max:"1"`
	// This setting boosts the similarity to the original speaker. Using this setting requires a slightly higher computational load, which in turn increases latency.
	UseSpeakerBoost *bool `json:"use_speaker_boost,omitempty" xml:"use_speaker_boost,omitempty"`
	// Determines how closely the AI should adhere to the original voice when attempting to replicate it.
	SimilarityBoost *float64 `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty" min:"0" max:"1"`
	// Determines the style exaggeration of the voice. This setting attempts to amplify the style of the original speaker. It does consume additional computational resources and might increase latency if set to anything other than 0.
	Style *float64 `json:"style,omitempty" xml:"style,omitempty"`
	// Adjusts the speed of the voice. A value of 1.0 is the default speed, while values less than 1.0 slow down the speech, and values greater than 1.0 speed it up.
	Speed *float64 `json:"speed,omitempty" xml:"speed,omitempty"`
	// Controls how closely the output matches your professional voice clone samples (0–1). Lower values sound more versatile, while higher values preserve more of the original voice characteristics. We recommend 0 for most use cases, but if you do run into issues 0.5 could be a good fallback. Applies only to professional voice clones fine-tuned for the selected v4 model.
	CloneStrength *float64 `json:"clone_strength,omitempty" xml:"clone_strength,omitempty" min:"0" max:"1"`
}

func (v *VoiceSettings) GetStability() (value *float64) {
	if v == nil {
		return
	}
	return v.Stability
}

func (v *VoiceSettings) GetUseSpeakerBoost() (value *bool) {
	if v == nil {
		return
	}
	return v.UseSpeakerBoost
}

func (v *VoiceSettings) GetSimilarityBoost() (value *float64) {
	if v == nil {
		return
	}
	return v.SimilarityBoost
}

func (v *VoiceSettings) GetStyle() (value *float64) {
	if v == nil {
		return
	}
	return v.Style
}

func (v *VoiceSettings) GetSpeed() (value *float64) {
	if v == nil {
		return
	}
	return v.Speed
}

func (v *VoiceSettings) GetCloneStrength() (value *float64) {
	if v == nil {
		return
	}
	return v.CloneStrength
}

func (v *VoiceSettings) SetStability(stability *float64) {
	v.Stability = stability
}

func (v *VoiceSettings) SetUseSpeakerBoost(useSpeakerBoost *bool) {
	v.UseSpeakerBoost = useSpeakerBoost
}

func (v *VoiceSettings) SetSimilarityBoost(similarityBoost *float64) {
	v.SimilarityBoost = similarityBoost
}

func (v *VoiceSettings) SetStyle(style *float64) {
	v.Style = style
}

func (v *VoiceSettings) SetSpeed(speed *float64) {
	v.Speed = speed
}

func (v *VoiceSettings) SetCloneStrength(cloneStrength *float64) {
	v.CloneStrength = cloneStrength
}

func (v *VoiceSettings) GetExtraProperties() map[string]interface{} {
	return nil
}

func (v VoiceSettings) String() string {
	jsonData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "error converting struct: VoiceSettings to string"
	}
	return string(jsonData)
}

func (v *VoiceSettings) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, v)
}

func (v *VoiceSettings) MarshalJSON() ([]byte, error) {
	type embed VoiceSettings
	var marshaler = struct {
		embed
	}{
		embed: embed(*v),
	}
	return json.Marshal(marshaler)
}
