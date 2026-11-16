package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyComposeMusicWithADetailedResponseV1MusicDetailedPost struct {
	// When enable_logging is set to false zero retention mode will be used for the request. Zero retention mode may only be used by enterprise customers.
	EnableLogging *bool `json:"-" url:"enable_logging,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_logging"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. Use "auto" (the default) to let the API pick the best format for the selected model: mp3_44100_128 for v1 models and mp3_48000_192 for v2 models.
	OutputFormat *ComposeDetailedMusicRequestOutputFormat `json:"-" url:"output_format,omitempty" explode:"true" serializationStyle:"form" queryParam:"output_format"`
	// A simple text prompt to generate a song from. Cannot be used in conjunction with `composition_plan`.
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty" maxLength:"4100"`
	// A detailed composition plan to guide music generation. Cannot be used in conjunction with `prompt`.
	CompositionPlan *BodyComposeMusicWithADetailedResponseV1MusicDetailedPostCompositionPlan `json:"composition_plan,omitempty" xml:"composition_plan,omitempty"`
	// The length of the song to generate in milliseconds. Used only in conjunction with `prompt`. Must be between 3000ms and 600000ms. Optional - if not provided, the model will choose a length based on the prompt.
	MusicLengthMs *int          `json:"music_length_ms,omitempty" xml:"music_length_ms,omitempty" min:"3000" max:"600000"`
	ModelId       *MusicModelId `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// Random seed to initialize the music generation process. Providing the same seed with the same parameters can help achieve more consistent results, but exact reproducibility is not guaranteed and outputs may change across system updates. Cannot be used in conjunction with prompt.
	Seed *int `json:"seed,omitempty" xml:"seed,omitempty" min:"0" max:"2147483647"`
	// If true, guarantees that the generated song will be instrumental. If false, the song may or may not be instrumental depending on the `prompt`. Can only be used with `prompt`.
	ForceInstrumental *bool `json:"force_instrumental,omitempty" xml:"force_instrumental,omitempty"`
	// The ID of the finetune to use for the generation
	FinetuneId *string `json:"finetune_id,omitempty" xml:"finetune_id,omitempty" maxLength:"100"`
	// Controls how strictly section durations in the `composition_plan` are enforced. Only used with `composition_plan` and only applies to `music_v1`; for `music_v2` and `music_v2_5` section durations are always enforced and this is ignored. When false for `music_v1`, the model may adjust individual section durations for better quality and latency, while preserving the total song duration from the plan.
	RespectSectionsDurations *bool `json:"respect_sections_durations,omitempty" xml:"respect_sections_durations,omitempty"`
	// Whether to store the generated song for inpainting.
	StoreForInpainting *bool `json:"store_for_inpainting,omitempty" xml:"store_for_inpainting,omitempty"`
	// Whether to return the timestamps of the words in the generated song.
	WithTimestamps *bool `json:"with_timestamps,omitempty" xml:"with_timestamps,omitempty"`
	// Whether to return the visual waveform of the generated song.
	WithWaveformVisual *bool `json:"with_waveform_visual,omitempty" xml:"with_waveform_visual,omitempty"`
	// Whether to sign the generated song with C2PA. Applicable only for mp3 files.
	SignWithC2Pa *bool `json:"sign_with_c2pa,omitempty" xml:"sign_with_c2pa,omitempty"`
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetEnableLogging() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableLogging
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetOutputFormat() (value *ComposeDetailedMusicRequestOutputFormat) {
	if b == nil {
		return
	}
	return b.OutputFormat
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetPrompt() (value *string) {
	if b == nil {
		return
	}
	return b.Prompt
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetCompositionPlan() (value *BodyComposeMusicWithADetailedResponseV1MusicDetailedPostCompositionPlan) {
	if b == nil {
		return
	}
	return b.CompositionPlan
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetMusicLengthMs() (value *int) {
	if b == nil {
		return
	}
	return b.MusicLengthMs
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetModelId() (value *MusicModelId) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetSeed() (value *int) {
	if b == nil {
		return
	}
	return b.Seed
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetForceInstrumental() (value *bool) {
	if b == nil {
		return
	}
	return b.ForceInstrumental
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetFinetuneId() (value *string) {
	if b == nil {
		return
	}
	return b.FinetuneId
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetRespectSectionsDurations() (value *bool) {
	if b == nil {
		return
	}
	return b.RespectSectionsDurations
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetStoreForInpainting() (value *bool) {
	if b == nil {
		return
	}
	return b.StoreForInpainting
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetWithTimestamps() (value *bool) {
	if b == nil {
		return
	}
	return b.WithTimestamps
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetWithWaveformVisual() (value *bool) {
	if b == nil {
		return
	}
	return b.WithWaveformVisual
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetSignWithC2Pa() (value *bool) {
	if b == nil {
		return
	}
	return b.SignWithC2Pa
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetEnableLogging(enableLogging *bool) {
	b.EnableLogging = enableLogging
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetOutputFormat(outputFormat *ComposeDetailedMusicRequestOutputFormat) {
	b.OutputFormat = outputFormat
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetPrompt(prompt *string) {
	b.Prompt = prompt
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetCompositionPlan(compositionPlan *BodyComposeMusicWithADetailedResponseV1MusicDetailedPostCompositionPlan) {
	b.CompositionPlan = compositionPlan
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetMusicLengthMs(musicLengthMs *int) {
	b.MusicLengthMs = musicLengthMs
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetModelId(modelId *MusicModelId) {
	b.ModelId = modelId
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetSeed(seed *int) {
	b.Seed = seed
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetForceInstrumental(forceInstrumental *bool) {
	b.ForceInstrumental = forceInstrumental
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetFinetuneId(finetuneId *string) {
	b.FinetuneId = finetuneId
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetRespectSectionsDurations(respectSectionsDurations *bool) {
	b.RespectSectionsDurations = respectSectionsDurations
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetStoreForInpainting(storeForInpainting *bool) {
	b.StoreForInpainting = storeForInpainting
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetWithTimestamps(withTimestamps *bool) {
	b.WithTimestamps = withTimestamps
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetWithWaveformVisual(withWaveformVisual *bool) {
	b.WithWaveformVisual = withWaveformVisual
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) SetSignWithC2Pa(signWithC2Pa *bool) {
	b.SignWithC2Pa = signWithC2Pa
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyComposeMusicWithADetailedResponseV1MusicDetailedPost to string"
	}
	return string(jsonData)
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyComposeMusicWithADetailedResponseV1MusicDetailedPost) MarshalJSON() ([]byte, error) {
	type embed BodyComposeMusicWithADetailedResponseV1MusicDetailedPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
