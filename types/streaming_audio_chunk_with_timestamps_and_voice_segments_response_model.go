package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel struct {
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

func (s *StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel) GetAudioBase64() (value string) {
	if s == nil {
		return
	}
	return s.AudioBase64
}

func (s *StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel) GetAlignment() (value *CharacterAlignmentResponseModel) {
	if s == nil {
		return
	}
	return s.Alignment
}

func (s *StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel) GetNormalizedAlignment() (value *CharacterAlignmentResponseModel) {
	if s == nil {
		return
	}
	return s.NormalizedAlignment
}

func (s *StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel) GetVoiceSegments() (value []*VoiceSegment) {
	if s == nil {
		return
	}
	return s.VoiceSegments
}

func (s *StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel) SetAudioBase64(audioBase64 string) {
	s.AudioBase64 = audioBase64
}

func (s *StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel) SetAlignment(alignment *CharacterAlignmentResponseModel) {
	s.Alignment = alignment
}

func (s *StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel) SetNormalizedAlignment(normalizedAlignment *CharacterAlignmentResponseModel) {
	s.NormalizedAlignment = normalizedAlignment
}

func (s *StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel) SetVoiceSegments(voiceSegments []*VoiceSegment) {
	s.VoiceSegments = voiceSegments
}

func (s *StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel to string"
	}
	return string(jsonData)
}

func (s *StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel) UnmarshalJSON(data []byte) error {
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

func (s *StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel) MarshalJSON() ([]byte, error) {
	type embed StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
