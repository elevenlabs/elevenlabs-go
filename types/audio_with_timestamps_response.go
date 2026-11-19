package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioWithTimestampsResponse struct {
	// Base64 encoded audio data
	AudioBase64 string `json:"audio_base64" xml:"audio_base64" required:"true"`
	// Timestamp information for each character in the original text
	Alignment *CharacterAlignmentResponseModel `json:"alignment,omitempty" xml:"alignment,omitempty"`
	// Timestamp information for each character in the normalized text
	NormalizedAlignment *CharacterAlignmentResponseModel `json:"normalized_alignment,omitempty" xml:"normalized_alignment,omitempty"`
	extraProperties     map[string]interface{}
}

func (a *AudioWithTimestampsResponse) GetAudioBase64() (value string) {
	if a == nil {
		return
	}
	return a.AudioBase64
}

func (a *AudioWithTimestampsResponse) GetAlignment() (value *CharacterAlignmentResponseModel) {
	if a == nil {
		return
	}
	return a.Alignment
}

func (a *AudioWithTimestampsResponse) GetNormalizedAlignment() (value *CharacterAlignmentResponseModel) {
	if a == nil {
		return
	}
	return a.NormalizedAlignment
}

func (a *AudioWithTimestampsResponse) SetAudioBase64(audioBase64 string) {
	a.AudioBase64 = audioBase64
}

func (a *AudioWithTimestampsResponse) SetAlignment(alignment *CharacterAlignmentResponseModel) {
	a.Alignment = alignment
}

func (a *AudioWithTimestampsResponse) SetNormalizedAlignment(normalizedAlignment *CharacterAlignmentResponseModel) {
	a.NormalizedAlignment = normalizedAlignment
}

func (a *AudioWithTimestampsResponse) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioWithTimestampsResponse) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioWithTimestampsResponse to string"
	}
	return string(jsonData)
}

func (a *AudioWithTimestampsResponse) UnmarshalJSON(data []byte) error {
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

func (a *AudioWithTimestampsResponse) MarshalJSON() ([]byte, error) {
	type embed AudioWithTimestampsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
