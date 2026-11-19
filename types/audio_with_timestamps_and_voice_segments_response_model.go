package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioWithTimestampsAndVoiceSegmentsResponseModel struct {
	// Base64 encoded audio data
	AudioBase64 string `json:"audio_base64" xml:"audio_base64" required:"true"`
	// Timestamp information for each character in the original text
	Alignment *CharacterAlignmentResponseModel `json:"alignment,omitempty" xml:"alignment,omitempty"`
	// Timestamp information for each character in the normalized text
	NormalizedAlignment *CharacterAlignmentResponseModel `json:"normalized_alignment,omitempty" xml:"normalized_alignment,omitempty"`
	// Voice segments for the audio
	VoiceSegments   []*VoiceSegment `json:"voice_segments" xml:"voice_segments" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AudioWithTimestampsAndVoiceSegmentsResponseModel) GetAudioBase64() (value string) {
	if a == nil {
		return
	}
	return a.AudioBase64
}

func (a *AudioWithTimestampsAndVoiceSegmentsResponseModel) GetAlignment() (value *CharacterAlignmentResponseModel) {
	if a == nil {
		return
	}
	return a.Alignment
}

func (a *AudioWithTimestampsAndVoiceSegmentsResponseModel) GetNormalizedAlignment() (value *CharacterAlignmentResponseModel) {
	if a == nil {
		return
	}
	return a.NormalizedAlignment
}

func (a *AudioWithTimestampsAndVoiceSegmentsResponseModel) GetVoiceSegments() (value []*VoiceSegment) {
	if a == nil {
		return
	}
	return a.VoiceSegments
}

func (a *AudioWithTimestampsAndVoiceSegmentsResponseModel) SetAudioBase64(audioBase64 string) {
	a.AudioBase64 = audioBase64
}

func (a *AudioWithTimestampsAndVoiceSegmentsResponseModel) SetAlignment(alignment *CharacterAlignmentResponseModel) {
	a.Alignment = alignment
}

func (a *AudioWithTimestampsAndVoiceSegmentsResponseModel) SetNormalizedAlignment(normalizedAlignment *CharacterAlignmentResponseModel) {
	a.NormalizedAlignment = normalizedAlignment
}

func (a *AudioWithTimestampsAndVoiceSegmentsResponseModel) SetVoiceSegments(voiceSegments []*VoiceSegment) {
	a.VoiceSegments = voiceSegments
}

func (a *AudioWithTimestampsAndVoiceSegmentsResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioWithTimestampsAndVoiceSegmentsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioWithTimestampsAndVoiceSegmentsResponseModel to string"
	}
	return string(jsonData)
}

func (a *AudioWithTimestampsAndVoiceSegmentsResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioWithTimestampsAndVoiceSegmentsResponseModel) MarshalJSON() ([]byte, error) {
	type embed AudioWithTimestampsAndVoiceSegmentsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
