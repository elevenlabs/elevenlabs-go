package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type AudioIsolationRequest struct {
	// The audio file from which vocals/speech will be isolated from.
	Audio io.Reader `json:"audio" xml:"audio" required:"true"`
	// The format of input audio. Options are 'pcm_s16le_16' or 'other' For `pcm_s16le_16`, the input audio must be 16-bit PCM at a 16kHz sample rate, single channel (mono), and little-endian byte order. Latency will be lower than with passing an encoded waveform.
	FileFormat *ConvertAudioIsolationRequestFileFormat `json:"file_format,omitempty" xml:"file_format,omitempty"`
	// Optional preview image base64 for tracking this generation.
	PreviewB64 *string `json:"preview_b64,omitempty" xml:"preview_b64,omitempty"`
}

func (a *AudioIsolationRequest) GetAudio() (value io.Reader) {
	if a == nil {
		return
	}
	return a.Audio
}

func (a *AudioIsolationRequest) GetFileFormat() (value *ConvertAudioIsolationRequestFileFormat) {
	if a == nil {
		return
	}
	return a.FileFormat
}

func (a *AudioIsolationRequest) GetPreviewB64() (value *string) {
	if a == nil {
		return
	}
	return a.PreviewB64
}

func (a *AudioIsolationRequest) SetAudio(audio io.Reader) {
	a.Audio = audio
}

func (a *AudioIsolationRequest) SetFileFormat(fileFormat *ConvertAudioIsolationRequestFileFormat) {
	a.FileFormat = fileFormat
}

func (a *AudioIsolationRequest) SetPreviewB64(previewB64 *string) {
	a.PreviewB64 = previewB64
}

func (a *AudioIsolationRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (a AudioIsolationRequest) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioIsolationRequest to string"
	}
	return string(jsonData)
}

func (a *AudioIsolationRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, a)
}

func (a *AudioIsolationRequest) MarshalJSON() ([]byte, error) {
	type embed AudioIsolationRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
