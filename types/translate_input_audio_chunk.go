package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TranslateInputAudioChunk struct {
	messageType string
	// Base64-encoded audio data. Recommended chunk size is around 100 ms of audio.
	AudioBase64     string `json:"audio_base_64" xml:"audio_base_64" required:"true"`
	extraProperties map[string]interface{}
}

func (t *TranslateInputAudioChunk) GetAudioBase64() (value string) {
	if t == nil {
		return
	}
	return t.AudioBase64
}

func (t *TranslateInputAudioChunk) SetAudioBase64(audioBase64 string) {
	t.AudioBase64 = audioBase64
}

func (t *TranslateInputAudioChunk) MessageType() string {
	return t.messageType
}

func (t *TranslateInputAudioChunk) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TranslateInputAudioChunk) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TranslateInputAudioChunk to string"
	}
	return string(jsonData)
}

func (t *TranslateInputAudioChunk) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, t); err != nil {
		return err
	}
	type alias TranslateInputAudioChunk
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*t = TranslateInputAudioChunk(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "input_audio_chunk" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", t, "input_audio_chunk", *literals.MessageType)
	}
	t.messageType = "input_audio_chunk"
	extraProperties, err := internal.ExtractExtraProperties(data, *t, "message_type")
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TranslateInputAudioChunk) MarshalJSON() ([]byte, error) {
	type embed TranslateInputAudioChunk
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*t),
		MessageType: "input_audio_chunk",
	}
	return json.Marshal(marshaler)
}
