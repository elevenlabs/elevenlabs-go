package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SubtitlingInputAudioChunk struct {
	messageType string
	// Base64-encoded audio data. Recommended chunk size is around 100 ms of audio.
	AudioBase64     string `json:"audio_base_64" xml:"audio_base_64" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SubtitlingInputAudioChunk) GetAudioBase64() (value string) {
	if s == nil {
		return
	}
	return s.AudioBase64
}

func (s *SubtitlingInputAudioChunk) SetAudioBase64(audioBase64 string) {
	s.AudioBase64 = audioBase64
}

func (s *SubtitlingInputAudioChunk) MessageType() string {
	return s.messageType
}

func (s *SubtitlingInputAudioChunk) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SubtitlingInputAudioChunk) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SubtitlingInputAudioChunk to string"
	}
	return string(jsonData)
}

func (s *SubtitlingInputAudioChunk) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SubtitlingInputAudioChunk
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SubtitlingInputAudioChunk(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "input_audio_chunk" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", s, "input_audio_chunk", *literals.MessageType)
	}
	s.messageType = "input_audio_chunk"
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "message_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SubtitlingInputAudioChunk) MarshalJSON() ([]byte, error) {
	type embed SubtitlingInputAudioChunk
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*s),
		MessageType: "input_audio_chunk",
	}
	return json.Marshal(marshaler)
}
