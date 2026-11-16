package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for structured music generation endpoint
type DetailedMusicResponse struct {
	// The composition plan used to generate the song
	CompositionPlan *DetailedMusicResponseCompositionPlan `json:"composition_plan" xml:"composition_plan" required:"true"`
	SongMetadata    *SongMetadata                         `json:"song_metadata" xml:"song_metadata" required:"true"`
	// The timestamps of the words in the generated song
	WordsTimestamps []*WordTimestamp `json:"words_timestamps" xml:"words_timestamps" required:"true"`
	// A low-resolution waveform of the generated song, for showing a preview of it. Holds 4 values per second of audio, from -1000 to 1000. Stereo is mixed down to a single channel. Only present if `with_waveform_visual` was True in the request body.
	WaveformVisual  []int `json:"waveform_visual,omitempty" xml:"waveform_visual,omitempty"`
	extraProperties map[string]interface{}
}

func (d *DetailedMusicResponse) GetCompositionPlan() (value *DetailedMusicResponseCompositionPlan) {
	if d == nil {
		return
	}
	return d.CompositionPlan
}

func (d *DetailedMusicResponse) GetSongMetadata() (value *SongMetadata) {
	if d == nil {
		return
	}
	return d.SongMetadata
}

func (d *DetailedMusicResponse) GetWordsTimestamps() (value []*WordTimestamp) {
	if d == nil {
		return
	}
	return d.WordsTimestamps
}

func (d *DetailedMusicResponse) GetWaveformVisual() (value []int) {
	if d == nil {
		return
	}
	return d.WaveformVisual
}

func (d *DetailedMusicResponse) SetCompositionPlan(compositionPlan *DetailedMusicResponseCompositionPlan) {
	d.CompositionPlan = compositionPlan
}

func (d *DetailedMusicResponse) SetSongMetadata(songMetadata *SongMetadata) {
	d.SongMetadata = songMetadata
}

func (d *DetailedMusicResponse) SetWordsTimestamps(wordsTimestamps []*WordTimestamp) {
	d.WordsTimestamps = wordsTimestamps
}

func (d *DetailedMusicResponse) SetWaveformVisual(waveformVisual []int) {
	d.WaveformVisual = waveformVisual
}

func (d *DetailedMusicResponse) GetExtraProperties() map[string]interface{} {
	if d == nil {
		return nil
	}
	return d.extraProperties
}

func (d DetailedMusicResponse) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "error converting struct: DetailedMusicResponse to string"
	}
	return string(jsonData)
}

func (d *DetailedMusicResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, d); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties
	return nil
}

func (d *DetailedMusicResponse) MarshalJSON() ([]byte, error) {
	type embed DetailedMusicResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*d),
	}
	return json.Marshal(marshaler)
}
