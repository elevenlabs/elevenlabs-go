package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyStreamComposedMusicV1MusicStreamPost struct {
	// When enable_logging is set to false zero retention mode will be used for the request. Zero retention mode may only be used by enterprise customers.
	EnableLogging *bool `json:"-" url:"enable_logging,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_logging"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. Use "auto" (the default) to let the API pick the best format for the selected model: mp3_44100_128 for v1 models and mp3_48000_192 for v2 models.
	OutputFormat *StreamMusicRequestOutputFormat `json:"-" url:"output_format,omitempty" explode:"true" serializationStyle:"form" queryParam:"output_format"`
	// A simple text prompt to generate a song from. Cannot be used in conjunction with `composition_plan`.
	Prompt *string `json:"prompt,omitempty" xml:"prompt,omitempty" maxLength:"4100"`
	// A detailed composition plan to guide music generation. Cannot be used in conjunction with `prompt`.
	CompositionPlan *BodyStreamComposedMusicV1MusicStreamPostCompositionPlan `json:"composition_plan,omitempty" xml:"composition_plan,omitempty"`
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
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) GetEnableLogging() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableLogging
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) GetOutputFormat() (value *StreamMusicRequestOutputFormat) {
	if b == nil {
		return
	}
	return b.OutputFormat
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) GetPrompt() (value *string) {
	if b == nil {
		return
	}
	return b.Prompt
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) GetCompositionPlan() (value *BodyStreamComposedMusicV1MusicStreamPostCompositionPlan) {
	if b == nil {
		return
	}
	return b.CompositionPlan
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) GetMusicLengthMs() (value *int) {
	if b == nil {
		return
	}
	return b.MusicLengthMs
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) GetModelId() (value *MusicModelId) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) GetSeed() (value *int) {
	if b == nil {
		return
	}
	return b.Seed
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) GetForceInstrumental() (value *bool) {
	if b == nil {
		return
	}
	return b.ForceInstrumental
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) GetFinetuneId() (value *string) {
	if b == nil {
		return
	}
	return b.FinetuneId
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) GetStoreForInpainting() (value *bool) {
	if b == nil {
		return
	}
	return b.StoreForInpainting
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) SetEnableLogging(enableLogging *bool) {
	b.EnableLogging = enableLogging
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) SetOutputFormat(outputFormat *StreamMusicRequestOutputFormat) {
	b.OutputFormat = outputFormat
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) SetPrompt(prompt *string) {
	b.Prompt = prompt
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) SetCompositionPlan(compositionPlan *BodyStreamComposedMusicV1MusicStreamPostCompositionPlan) {
	b.CompositionPlan = compositionPlan
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) SetMusicLengthMs(musicLengthMs *int) {
	b.MusicLengthMs = musicLengthMs
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) SetModelId(modelId *MusicModelId) {
	b.ModelId = modelId
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) SetSeed(seed *int) {
	b.Seed = seed
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) SetForceInstrumental(forceInstrumental *bool) {
	b.ForceInstrumental = forceInstrumental
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) SetFinetuneId(finetuneId *string) {
	b.FinetuneId = finetuneId
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) SetStoreForInpainting(storeForInpainting *bool) {
	b.StoreForInpainting = storeForInpainting
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyStreamComposedMusicV1MusicStreamPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyStreamComposedMusicV1MusicStreamPost to string"
	}
	return string(jsonData)
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyStreamComposedMusicV1MusicStreamPost) MarshalJSON() ([]byte, error) {
	type embed BodyStreamComposedMusicV1MusicStreamPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
