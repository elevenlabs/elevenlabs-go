package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"io"
)

type BodyUploadMusicV1MusicUploadPost struct {
	// The audio file to upload.
	File io.Reader `json:"file" xml:"file" required:"true"`
	// Whether to generate and return the composition plan for the uploaded song. Pass a model id (`music_v1`, `music_v2` or `music_v2_5`) to control which composition plan format is returned. Passing `true`/`false` is deprecated; `true` defaults to the `music_v1` plan format. Enabling this will increase the latency.
	ExtractCompositionPlan *BodyUploadMusicV1MusicUploadPostExtractCompositionPlan `json:"extract_composition_plan,omitempty" xml:"extract_composition_plan,omitempty"`
	// Whether to transcribe the uploaded song and return word-level timestamps. If True, the response will include words_timestamps but will increase the latency.
	WithTimestamps *bool `json:"with_timestamps,omitempty" xml:"with_timestamps,omitempty"`
	// Whether to return the visual waveform of the uploaded song.
	WithWaveformVisual *bool `json:"with_waveform_visual,omitempty" xml:"with_waveform_visual,omitempty"`
}

func (b *BodyUploadMusicV1MusicUploadPost) GetFile() (value io.Reader) {
	if b == nil {
		return
	}
	return b.File
}

func (b *BodyUploadMusicV1MusicUploadPost) GetExtractCompositionPlan() (value *BodyUploadMusicV1MusicUploadPostExtractCompositionPlan) {
	if b == nil {
		return
	}
	return b.ExtractCompositionPlan
}

func (b *BodyUploadMusicV1MusicUploadPost) GetWithTimestamps() (value *bool) {
	if b == nil {
		return
	}
	return b.WithTimestamps
}

func (b *BodyUploadMusicV1MusicUploadPost) GetWithWaveformVisual() (value *bool) {
	if b == nil {
		return
	}
	return b.WithWaveformVisual
}

func (b *BodyUploadMusicV1MusicUploadPost) SetFile(file io.Reader) {
	b.File = file
}

func (b *BodyUploadMusicV1MusicUploadPost) SetExtractCompositionPlan(extractCompositionPlan *BodyUploadMusicV1MusicUploadPostExtractCompositionPlan) {
	b.ExtractCompositionPlan = extractCompositionPlan
}

func (b *BodyUploadMusicV1MusicUploadPost) SetWithTimestamps(withTimestamps *bool) {
	b.WithTimestamps = withTimestamps
}

func (b *BodyUploadMusicV1MusicUploadPost) SetWithWaveformVisual(withWaveformVisual *bool) {
	b.WithWaveformVisual = withWaveformVisual
}

func (b *BodyUploadMusicV1MusicUploadPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUploadMusicV1MusicUploadPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUploadMusicV1MusicUploadPost to string"
	}
	return string(jsonData)
}

func (b *BodyUploadMusicV1MusicUploadPost) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, b); err != nil {
		return err
	}
	type alias BodyUploadMusicV1MusicUploadPost
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*b = BodyUploadMusicV1MusicUploadPost(tmp)
	return nil
}

func (b *BodyUploadMusicV1MusicUploadPost) MarshalJSON() ([]byte, error) {
	type embed BodyUploadMusicV1MusicUploadPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
