package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

// Multipart response structure with JSON metadata and binary audio
type MultipartMusicResponse struct {
	// Response model for structured music generation endpoint
	Metadata *DetailedMusicResponse `json:"metadata" xml:"metadata" required:"true"`
	// Binary audio data in the requested format
	Audio           io.Reader `json:"audio" xml:"audio" required:"true"`
	extraProperties map[string]interface{}
}

func (m *MultipartMusicResponse) GetMetadata() (value *DetailedMusicResponse) {
	if m == nil {
		return
	}
	return m.Metadata
}

func (m *MultipartMusicResponse) GetAudio() (value io.Reader) {
	if m == nil {
		return
	}
	return m.Audio
}

func (m *MultipartMusicResponse) SetMetadata(metadata *DetailedMusicResponse) {
	m.Metadata = metadata
}

func (m *MultipartMusicResponse) SetAudio(audio io.Reader) {
	m.Audio = audio
}

func (m *MultipartMusicResponse) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MultipartMusicResponse) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MultipartMusicResponse to string"
	}
	return string(jsonData)
}

func (m *MultipartMusicResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias MultipartMusicResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = MultipartMusicResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MultipartMusicResponse) MarshalJSON() ([]byte, error) {
	type embed MultipartMusicResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
