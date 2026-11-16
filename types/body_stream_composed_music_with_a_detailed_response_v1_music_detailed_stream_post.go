package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost struct {
	// When enable_logging is set to false zero retention mode will be used for the request. Zero retention mode may only be used by enterprise customers.
	EnableLogging *bool `json:"-" url:"enable_logging,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_logging"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. Use "auto" (the default) to let the API pick the best format for the selected model: mp3_44100_128 for v1 models and mp3_48000_192 for v2 models.
	OutputFormat *ComposeDetailedStreamMusicRequestOutputFormat `json:"-" url:"output_format,omitempty" explode:"true" serializationStyle:"form" queryParam:"output_format"`
	// A simple text prompt to generate a song from. Cannot be used in conjunction with `composition_plan`.
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty" maxLength:"4100"`
	// A detailed composition plan to guide music generation. Cannot be used in conjunction with `prompt`.
	CompositionPlan *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPostCompositionPlan `json:"composition_plan,omitempty" xml:"composition_plan,omitempty"`
	// The length of the song to generate in milliseconds. Used only in conjunction with `prompt`. Must be between 3000ms and 600000ms. Optional - if not provided, the model will choose a length based on the prompt.
	MusicLengthMs *int          `json:"music_length_ms,omitempty" xml:"music_length_ms,omitempty" min:"3000" max:"600000"`
	ModelId       *MusicModelId `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// Random seed to initialize the music generation process. Providing the same seed with the same parameters can help achieve more consistent results, but exact reproducibility is not guaranteed and outputs may change across system updates. Cannot be used in conjunction with prompt.
	Seed *int `json:"seed,omitempty" xml:"seed,omitempty" min:"0" max:"2147483647"`
	// If true, guarantees that the generated song will be instrumental. If false, the song may or may not be instrumental depending on the `prompt`. Can only be used with `prompt`.
	ForceInstrumental *bool `json:"force_instrumental,omitempty" xml:"force_instrumental,omitempty"`
	// The ID of the finetune to use for the generation
	FinetuneId *string `json:"finetune_id,omitempty" xml:"finetune_id,omitempty" maxLength:"100"`
	// Whether to store the generated song for inpainting.
	StoreForInpainting *bool `json:"store_for_inpainting,omitempty" xml:"store_for_inpainting,omitempty"`
	// Whether to return the timestamps of the words in the generated song.
	WithTimestamps *bool `json:"with_timestamps,omitempty" xml:"with_timestamps,omitempty"`
	// Whether to return the visual waveform of the generated song.
	WithWaveformVisual *bool `json:"with_waveform_visual,omitempty" xml:"with_waveform_visual,omitempty"`
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetEnableLogging() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableLogging
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetOutputFormat() (value *ComposeDetailedStreamMusicRequestOutputFormat) {
	if b == nil {
		return
	}
	return b.OutputFormat
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetPrompt() (value *string) {
	if b == nil {
		return
	}
	return b.Prompt
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetCompositionPlan() (value *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPostCompositionPlan) {
	if b == nil {
		return
	}
	return b.CompositionPlan
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetMusicLengthMs() (value *int) {
	if b == nil {
		return
	}
	return b.MusicLengthMs
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetModelId() (value *MusicModelId) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetSeed() (value *int) {
	if b == nil {
		return
	}
	return b.Seed
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetForceInstrumental() (value *bool) {
	if b == nil {
		return
	}
	return b.ForceInstrumental
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetFinetuneId() (value *string) {
	if b == nil {
		return
	}
	return b.FinetuneId
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetStoreForInpainting() (value *bool) {
	if b == nil {
		return
	}
	return b.StoreForInpainting
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetWithTimestamps() (value *bool) {
	if b == nil {
		return
	}
	return b.WithTimestamps
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetWithWaveformVisual() (value *bool) {
	if b == nil {
		return
	}
	return b.WithWaveformVisual
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) SetEnableLogging(enableLogging *bool) {
	b.EnableLogging = enableLogging
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) SetOutputFormat(outputFormat *ComposeDetailedStreamMusicRequestOutputFormat) {
	b.OutputFormat = outputFormat
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) SetPrompt(prompt *string) {
	b.Prompt = prompt
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) SetCompositionPlan(compositionPlan *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPostCompositionPlan) {
	b.CompositionPlan = compositionPlan
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) SetMusicLengthMs(musicLengthMs *int) {
	b.MusicLengthMs = musicLengthMs
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) SetModelId(modelId *MusicModelId) {
	b.ModelId = modelId
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) SetSeed(seed *int) {
	b.Seed = seed
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) SetForceInstrumental(forceInstrumental *bool) {
	b.ForceInstrumental = forceInstrumental
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) SetFinetuneId(finetuneId *string) {
	b.FinetuneId = finetuneId
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) SetStoreForInpainting(storeForInpainting *bool) {
	b.StoreForInpainting = storeForInpainting
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) SetWithTimestamps(withTimestamps *bool) {
	b.WithTimestamps = withTimestamps
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) SetWithWaveformVisual(withWaveformVisual *bool) {
	b.WithWaveformVisual = withWaveformVisual
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost to string"
	}
	return string(jsonData)
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost) MarshalJSON() ([]byte, error) {
	type embed BodyStreamComposedMusicWithADetailedResponseV1MusicDetailedStreamPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
