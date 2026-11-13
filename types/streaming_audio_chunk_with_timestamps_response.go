package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type StreamingAudioChunkWithTimestampsResponse struct {
	// Base64 encoded audio data
	AudioBase64 string `json:"audio_base64" xml:"audio_base64" required:"true"`
	// Timestamp information for each character in the original text
	Alignment *CharacterAlignmentResponseModel `json:"alignment,omitempty" xml:"alignment,omitempty"`
	// Timestamp information for each character in the normalized text
	NormalizedAlignment *CharacterAlignmentResponseModel `json:"normalized_alignment,omitempty" xml:"normalized_alignment,omitempty"`
	extraProperties     map[string]interface{}
}

func (s *StreamingAudioChunkWithTimestampsResponse) GetAudioBase64() (value string) {
	if s == nil {
		return
	}
	return s.AudioBase64
}

func (s *StreamingAudioChunkWithTimestampsResponse) GetAlignment() (value *CharacterAlignmentResponseModel) {
	if s == nil {
		return
	}
	return s.Alignment
}

func (s *StreamingAudioChunkWithTimestampsResponse) GetNormalizedAlignment() (value *CharacterAlignmentResponseModel) {
	if s == nil {
		return
	}
	return s.NormalizedAlignment
}

func (s *StreamingAudioChunkWithTimestampsResponse) SetAudioBase64(audioBase64 string) {
	s.AudioBase64 = audioBase64
}

func (s *StreamingAudioChunkWithTimestampsResponse) SetAlignment(alignment *CharacterAlignmentResponseModel) {
	s.Alignment = alignment
}

func (s *StreamingAudioChunkWithTimestampsResponse) SetNormalizedAlignment(normalizedAlignment *CharacterAlignmentResponseModel) {
	s.NormalizedAlignment = normalizedAlignment
}

func (s *StreamingAudioChunkWithTimestampsResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StreamingAudioChunkWithTimestampsResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StreamingAudioChunkWithTimestampsResponse to string"
	}
	return string(jsonData)
}

func (s *StreamingAudioChunkWithTimestampsResponse) UnmarshalJSON(data []byte) error {
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

func (s *StreamingAudioChunkWithTimestampsResponse) MarshalJSON() ([]byte, error) {
	type embed StreamingAudioChunkWithTimestampsResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
