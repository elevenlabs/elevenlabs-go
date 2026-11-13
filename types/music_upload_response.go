package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for music upload endpoint.
type MusicUploadResponse struct {
	// Unique identifier for the uploaded song
	SongId string `json:"song_id" xml:"song_id" required:"true"`
	// The composition plan extracted from the uploaded song. Only present if `extract_composition_plan` was provided in the request body.
	CompositionPlan *MusicUploadResponseCompositionPlan `json:"composition_plan,omitempty" xml:"composition_plan,omitempty"`
	// Word-level timestamps transcribed from the uploaded song. Only present if `with_timestamps` was True in the request body
	WordsTimestamps []*WordTimestamp `json:"words_timestamps,omitempty" xml:"words_timestamps,omitempty"`
	// A low-resolution waveform of the uploaded song, for showing a preview of it. Holds 4 values per second of audio, from -1000 to 1000. Stereo is mixed down to a single channel. Only present if `with_waveform_visual` was True in the request body.
	WaveformVisual  []int `json:"waveform_visual,omitempty" xml:"waveform_visual,omitempty"`
	extraProperties map[string]interface{}
}

func (m *MusicUploadResponse) GetSongId() (value string) {
	if m == nil {
		return
	}
	return m.SongId
}

func (m *MusicUploadResponse) GetCompositionPlan() (value *MusicUploadResponseCompositionPlan) {
	if m == nil {
		return
	}
	return m.CompositionPlan
}

func (m *MusicUploadResponse) GetWordsTimestamps() (value []*WordTimestamp) {
	if m == nil {
		return
	}
	return m.WordsTimestamps
}

func (m *MusicUploadResponse) GetWaveformVisual() (value []int) {
	if m == nil {
		return
	}
	return m.WaveformVisual
}

func (m *MusicUploadResponse) SetSongId(songId string) {
	m.SongId = songId
}

func (m *MusicUploadResponse) SetCompositionPlan(compositionPlan *MusicUploadResponseCompositionPlan) {
	m.CompositionPlan = compositionPlan
}

func (m *MusicUploadResponse) SetWordsTimestamps(wordsTimestamps []*WordTimestamp) {
	m.WordsTimestamps = wordsTimestamps
}

func (m *MusicUploadResponse) SetWaveformVisual(waveformVisual []int) {
	m.WaveformVisual = waveformVisual
}

func (m *MusicUploadResponse) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MusicUploadResponse) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MusicUploadResponse to string"
	}
	return string(jsonData)
}

func (m *MusicUploadResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MusicUploadResponse) MarshalJSON() ([]byte, error) {
	type embed MusicUploadResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
