package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type DubbingInputAudioChunk struct {
	messageType     string
	AudioBase64     string `json:"audio_base_64" xml:"audio_base_64" required:"true" minLength:"1"`
	extraProperties map[string]interface{}
}

func (d *DubbingInputAudioChunk) GetAudioBase64() (value string) {
	if d == nil {
		return
	}
	return d.AudioBase64
}

func (d *DubbingInputAudioChunk) SetAudioBase64(audioBase64 string) {
	d.AudioBase64 = audioBase64
}

func (d *DubbingInputAudioChunk) MessageType() string {
	return d.messageType
}

func (d *DubbingInputAudioChunk) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DubbingInputAudioChunk) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DubbingInputAudioChunk to string"
	}
	return string(jsonData)
}

func (d *DubbingInputAudioChunk) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, d); err != nil {
		return err
	}
	type alias DubbingInputAudioChunk
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*d = DubbingInputAudioChunk(tmp)
	var literals struct {
		MessageType *string `json:"message_type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.MessageType != nil && *literals.MessageType != "input_audio_chunk" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", d, "input_audio_chunk", *literals.MessageType)
	}
	d.messageType = "input_audio_chunk"
	extraProperties, err := internal.ExtractExtraProperties(data, *d, "message_type")
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DubbingInputAudioChunk) MarshalJSON() ([]byte, error) {
	type embed DubbingInputAudioChunk
	var marshaler = struct {
		embed
		MessageType string `json:"message_type"`
	}{
		embed:       embed(*d),
		MessageType: "input_audio_chunk",
	}
	return json.Marshal(marshaler)
}
