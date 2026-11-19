package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type AudioIsolationStreamRequest struct {
	// The audio file from which vocals/speech will be isolated from.
	Audio io.Reader `json:"audio" xml:"audio" required:"true"`
	// The format of input audio. Options are 'pcm_s16le_16' or 'other' For `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate, single channel (mono), and little-endian byte order. Latency will be lower than with passing an encoded waveform.
	FileFormat *StreamAudioIsolationRequestFileFormat `json:"file_format,omitempty" xml:"file_format,omitempty"`
}

func (a *AudioIsolationStreamRequest) GetAudio() (value io.Reader) {
	if a == nil {
		return
	}
	return a.Audio
}

func (a *AudioIsolationStreamRequest) GetFileFormat() (value *StreamAudioIsolationRequestFileFormat) {
	if a == nil {
		return
	}
	return a.FileFormat
}

func (a *AudioIsolationStreamRequest) SetAudio(audio io.Reader) {
	a.Audio = audio
}

func (a *AudioIsolationStreamRequest) SetFileFormat(fileFormat *StreamAudioIsolationRequestFileFormat) {
	a.FileFormat = fileFormat
}

func (a *AudioIsolationStreamRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (a AudioIsolationStreamRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioIsolationStreamRequest to string"
	}
	return string(jsonData)
}

func (a *AudioIsolationStreamRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}

func (a *AudioIsolationStreamRequest) MarshalJSON() ([]byte, error) {
	type embed AudioIsolationStreamRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
