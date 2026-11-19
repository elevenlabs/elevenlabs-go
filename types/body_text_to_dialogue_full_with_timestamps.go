package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyTextToDialogueFullWithTimestamps struct {
	// When enable_logging is set to false zero retention mode will be used for the request. This will mean history features are unavailable for this request, including request stitching. Zero retention mode may only be used by enterprise customers.
	EnableLogging *bool `json:"-" url:"enable_logging,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_logging"`
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM and WAV formats with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
	OutputFormat *string `json:"-" url:"output_format,omitempty" explode:"true" serializationStyle:"form" queryParam:"output_format"`
	// A list of dialogue inputs, each containing text and a voice ID which will be converted into speech. The maximum number of unique voice IDs is 10. For reliable generation, keep the total character count across all `inputs[].text` values at or below 2,000 characters per request. Longer requests can terminate early in streaming responses or return a validation error.
	Inputs []*DialogueInput `json:"inputs" xml:"inputs" required:"true"`
	// Identifier of the model that will be used, you can query them using GET /v1/models. The model needs to have support for text to speech, you can check this using the can_do_text_to_speech property.
	ModelId *string `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// Language code (ISO 639-1) used to enforce a language for the model and text normalization. If the model does not support the provided language code, it will be ignored. This parameter is not supported for multilingual_v2 models.
	LanguageCode *string `json:"language_code,omitempty" xml:"language_code,omitempty"`
	// Settings controlling the dialogue generation.
	Settings *ToDialogueSettingsResponseModel `json:"settings,omitempty" xml:"settings,omitempty"`
	// A list of pronunciation dictionary locators (id, version_id) to be applied to the text. They will be applied in order. You may have up to 3 locators per request
	PronunciationDictionaryLocators []*PronunciationDictionaryVersionLocatorRequestModel `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	// If specified, our system will make a best effort to sample deterministically, such that repeated requests with the same seed and parameters should return the same result. Determinism is not guaranteed. Must be integer between 0 and 4294967295.
	Seed *int `json:"seed,omitempty" xml:"seed,omitempty"`
	// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
	ApplyTextNormalization *BodyTextToDialogueFullWithTimestampsApplyTextNormalization `json:"apply_text_normalization,omitempty" xml:"apply_text_normalization,omitempty"`
	// Whether to use the IVC version of a professional voice. This may improve expressiveness and reduce latency.
	UsePvcAsIvc *bool `json:"use_pvc_as_ivc,omitempty" xml:"use_pvc_as_ivc,omitempty"`
	// A list of request_ids of dialogue generations that came before this one. Used to condition the model for continuity when splitting a large task into multiple requests. A maximum of 3 request_ids can be sent. The last request_id is the audio which is closest to the current request. Not supported by every model.
	PreviousRequestIds []string `json:"previous_request_ids,omitempty" xml:"previous_request_ids,omitempty"`
	// A list of request_ids of dialogue generations that come after this one. Useful for maintaining continuity when regenerating a clip in the middle of a sequence. A maximum of 3 request_ids can be sent. The first request_id is the audio which is closest to the current request. Not supported by every model.
	NextRequestIds []string `json:"next_request_ids,omitempty" xml:"next_request_ids,omitempty"`
	// The text that comes immediately before this generation, used to condition the model for prosodic continuity. A maximum of 100 characters can be sent. Not supported by every model.
	PreviousText *string `json:"previous_text,omitempty" xml:"previous_text,omitempty" maxLength:"100"`
	// The text that comes immediately after this generation, used to condition the model for prosodic continuity. A maximum of 100 characters can be sent. Not supported by every model.
	FutureText *string `json:"future_text,omitempty" xml:"future_text,omitempty" maxLength:"100"`
	// Audio effects applied to the generated audio. Unknown IDs return 404 with `param` naming the field. Reverb (`environment_id` with `send_level` above 0) and a non-zero `pan` produce stereo audio and are rejected with 400 for `ulaw_*`, `alaw_*`, and `pcm_*` output formats. Effects don't change the cost of the request.
	AudioEffects *EffectsSpecInput `json:"audio_effects,omitempty" xml:"audio_effects,omitempty"`
}

func (b *BodyTextToDialogueFullWithTimestamps) GetEnableLogging() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableLogging
}

func (b *BodyTextToDialogueFullWithTimestamps) GetOutputFormat() (value *string) {
	if b == nil {
		return
	}
	return b.OutputFormat
}

func (b *BodyTextToDialogueFullWithTimestamps) GetInputs() (value []*DialogueInput) {
	if b == nil {
		return
	}
	return b.Inputs
}

func (b *BodyTextToDialogueFullWithTimestamps) GetModelId() (value *string) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyTextToDialogueFullWithTimestamps) GetLanguageCode() (value *string) {
	if b == nil {
		return
	}
	return b.LanguageCode
}

func (b *BodyTextToDialogueFullWithTimestamps) GetSettings() (value *ToDialogueSettingsResponseModel) {
	if b == nil {
		return
	}
	return b.Settings
}

func (b *BodyTextToDialogueFullWithTimestamps) GetPronunciationDictionaryLocators() (value []*PronunciationDictionaryVersionLocatorRequestModel) {
	if b == nil {
		return
	}
	return b.PronunciationDictionaryLocators
}

func (b *BodyTextToDialogueFullWithTimestamps) GetSeed() (value *int) {
	if b == nil {
		return
	}
	return b.Seed
}

func (b *BodyTextToDialogueFullWithTimestamps) GetApplyTextNormalization() (value *BodyTextToDialogueFullWithTimestampsApplyTextNormalization) {
	if b == nil {
		return
	}
	return b.ApplyTextNormalization
}

func (b *BodyTextToDialogueFullWithTimestamps) GetUsePvcAsIvc() (value *bool) {
	if b == nil {
		return
	}
	return b.UsePvcAsIvc
}

func (b *BodyTextToDialogueFullWithTimestamps) GetPreviousRequestIds() (value []string) {
	if b == nil {
		return
	}
	return b.PreviousRequestIds
}

func (b *BodyTextToDialogueFullWithTimestamps) GetNextRequestIds() (value []string) {
	if b == nil {
		return
	}
	return b.NextRequestIds
}

func (b *BodyTextToDialogueFullWithTimestamps) GetPreviousText() (value *string) {
	if b == nil {
		return
	}
	return b.PreviousText
}

func (b *BodyTextToDialogueFullWithTimestamps) GetFutureText() (value *string) {
	if b == nil {
		return
	}
	return b.FutureText
}

func (b *BodyTextToDialogueFullWithTimestamps) GetAudioEffects() (value *EffectsSpecInput) {
	if b == nil {
		return
	}
	return b.AudioEffects
}

func (b *BodyTextToDialogueFullWithTimestamps) SetEnableLogging(enableLogging *bool) {
	b.EnableLogging = enableLogging
}

func (b *BodyTextToDialogueFullWithTimestamps) SetOutputFormat(outputFormat *string) {
	b.OutputFormat = outputFormat
}

func (b *BodyTextToDialogueFullWithTimestamps) SetInputs(inputs []*DialogueInput) {
	b.Inputs = inputs
}

func (b *BodyTextToDialogueFullWithTimestamps) SetModelId(modelId *string) {
	b.ModelId = modelId
}

func (b *BodyTextToDialogueFullWithTimestamps) SetLanguageCode(languageCode *string) {
	b.LanguageCode = languageCode
}

func (b *BodyTextToDialogueFullWithTimestamps) SetSettings(settings *ToDialogueSettingsResponseModel) {
	b.Settings = settings
}

func (b *BodyTextToDialogueFullWithTimestamps) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*PronunciationDictionaryVersionLocatorRequestModel) {
	b.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (b *BodyTextToDialogueFullWithTimestamps) SetSeed(seed *int) {
	b.Seed = seed
}

func (b *BodyTextToDialogueFullWithTimestamps) SetApplyTextNormalization(applyTextNormalization *BodyTextToDialogueFullWithTimestampsApplyTextNormalization) {
	b.ApplyTextNormalization = applyTextNormalization
}

func (b *BodyTextToDialogueFullWithTimestamps) SetUsePvcAsIvc(usePvcAsIvc *bool) {
	b.UsePvcAsIvc = usePvcAsIvc
}

func (b *BodyTextToDialogueFullWithTimestamps) SetPreviousRequestIds(previousRequestIds []string) {
	b.PreviousRequestIds = previousRequestIds
}

func (b *BodyTextToDialogueFullWithTimestamps) SetNextRequestIds(nextRequestIds []string) {
	b.NextRequestIds = nextRequestIds
}

func (b *BodyTextToDialogueFullWithTimestamps) SetPreviousText(previousText *string) {
	b.PreviousText = previousText
}

func (b *BodyTextToDialogueFullWithTimestamps) SetFutureText(futureText *string) {
	b.FutureText = futureText
}

func (b *BodyTextToDialogueFullWithTimestamps) SetAudioEffects(audioEffects *EffectsSpecInput) {
	b.AudioEffects = audioEffects
}

func (b *BodyTextToDialogueFullWithTimestamps) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyTextToDialogueFullWithTimestamps) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyTextToDialogueFullWithTimestamps to string"
	}
	return string(jsonData)
}

func (b *BodyTextToDialogueFullWithTimestamps) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyTextToDialogueFullWithTimestamps) MarshalJSON() ([]byte, error) {
	type embed BodyTextToDialogueFullWithTimestamps
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
