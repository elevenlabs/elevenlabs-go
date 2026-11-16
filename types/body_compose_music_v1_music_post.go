package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyComposeMusicV1MusicPost struct {
	// When enable_logging is set to false zero retention mode will be used for the request. Zero retention mode may only be used by enterprise customers.
	EnableLogging *bool `json:"-" url:"enable_logging,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_logging"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. Use "auto" (the default) to let the API pick the best format for the selected model: mp3_44100_128 for v1 models and mp3_48000_192 for v2 models.
	OutputFormat *ComposeMusicRequestOutputFormat `json:"-" url:"output_format,omitempty" explode:"true" serializationStyle:"form" queryParam:"output_format"`
	// A simple text prompt to generate a song from. Cannot be used in conjunction with `composition_plan`.
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty" maxLength:"4100"`
	// A detailed composition plan to guide music generation. Cannot be used in conjunction with `prompt`.
	CompositionPlan *BodyComposeMusicV1MusicPostCompositionPlan `json:"composition_plan,omitempty" xml:"composition_plan,omitempty"`
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
	// Whether to sign the generated song with C2PA. Applicable only for mp3 files.
	SignWithC2Pa *bool `json:"sign_with_c2pa,omitempty" xml:"sign_with_c2pa,omitempty"`
}

func (b *BodyComposeMusicV1MusicPost) GetEnableLogging() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableLogging
}

func (b *BodyComposeMusicV1MusicPost) GetOutputFormat() (value *ComposeMusicRequestOutputFormat) {
	if b == nil {
		return
	}
	return b.OutputFormat
}

func (b *BodyComposeMusicV1MusicPost) GetPrompt() (value *string) {
	if b == nil {
		return
	}
	return b.Prompt
}

func (b *BodyComposeMusicV1MusicPost) GetCompositionPlan() (value *BodyComposeMusicV1MusicPostCompositionPlan) {
	if b == nil {
		return
	}
	return b.CompositionPlan
}

func (b *BodyComposeMusicV1MusicPost) GetMusicLengthMs() (value *int) {
	if b == nil {
		return
	}
	return b.MusicLengthMs
}

func (b *BodyComposeMusicV1MusicPost) GetModelId() (value *MusicModelId) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyComposeMusicV1MusicPost) GetSeed() (value *int) {
	if b == nil {
		return
	}
	return b.Seed
}

func (b *BodyComposeMusicV1MusicPost) GetForceInstrumental() (value *bool) {
	if b == nil {
		return
	}
	return b.ForceInstrumental
}

func (b *BodyComposeMusicV1MusicPost) GetFinetuneId() (value *string) {
	if b == nil {
		return
	}
	return b.FinetuneId
}

func (b *BodyComposeMusicV1MusicPost) GetRespectSectionsDurations() (value *bool) {
	if b == nil {
		return
	}
	return b.RespectSectionsDurations
}

func (b *BodyComposeMusicV1MusicPost) GetStoreForInpainting() (value *bool) {
	if b == nil {
		return
	}
	return b.StoreForInpainting
}

func (b *BodyComposeMusicV1MusicPost) GetSignWithC2Pa() (value *bool) {
	if b == nil {
		return
	}
	return b.SignWithC2Pa
}

func (b *BodyComposeMusicV1MusicPost) SetEnableLogging(enableLogging *bool) {
	b.EnableLogging = enableLogging
}

func (b *BodyComposeMusicV1MusicPost) SetOutputFormat(outputFormat *ComposeMusicRequestOutputFormat) {
	b.OutputFormat = outputFormat
}

func (b *BodyComposeMusicV1MusicPost) SetPrompt(prompt *string) {
	b.Prompt = prompt
}

func (b *BodyComposeMusicV1MusicPost) SetCompositionPlan(compositionPlan *BodyComposeMusicV1MusicPostCompositionPlan) {
	b.CompositionPlan = compositionPlan
}

func (b *BodyComposeMusicV1MusicPost) SetMusicLengthMs(musicLengthMs *int) {
	b.MusicLengthMs = musicLengthMs
}

func (b *BodyComposeMusicV1MusicPost) SetModelId(modelId *MusicModelId) {
	b.ModelId = modelId
}

func (b *BodyComposeMusicV1MusicPost) SetSeed(seed *int) {
	b.Seed = seed
}

func (b *BodyComposeMusicV1MusicPost) SetForceInstrumental(forceInstrumental *bool) {
	b.ForceInstrumental = forceInstrumental
}

func (b *BodyComposeMusicV1MusicPost) SetFinetuneId(finetuneId *string) {
	b.FinetuneId = finetuneId
}

func (b *BodyComposeMusicV1MusicPost) SetRespectSectionsDurations(respectSectionsDurations *bool) {
	b.RespectSectionsDurations = respectSectionsDurations
}

func (b *BodyComposeMusicV1MusicPost) SetStoreForInpainting(storeForInpainting *bool) {
	b.StoreForInpainting = storeForInpainting
}

func (b *BodyComposeMusicV1MusicPost) SetSignWithC2Pa(signWithC2Pa *bool) {
	b.SignWithC2Pa = signWithC2Pa
}

func (b *BodyComposeMusicV1MusicPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyComposeMusicV1MusicPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyComposeMusicV1MusicPost to string"
	}
	return string(jsonData)
}

func (b *BodyComposeMusicV1MusicPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyComposeMusicV1MusicPost) MarshalJSON() ([]byte, error) {
	type embed BodyComposeMusicV1MusicPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
