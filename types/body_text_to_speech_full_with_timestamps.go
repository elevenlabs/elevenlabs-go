package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyTextToSpeechFullWithTimestamps struct {
	// Output format of the generated audio. Formatted as codec_sample_rate_bitrate. So an mp3 with 22.05kHz sample rate at 32kbs is represented as mp3_22050_32. MP3 with 192kbps bitrate requires you to be subscribed to Creator tier or above. PCM and WAV formats with 44.1kHz sample rate requires you to be subscribed to Pro tier or above. Note that the μ-law format (sometimes written mu-law, often approximated as u-law) is commonly used for Twilio audio inputs.
	OutputFormat *ConvertWithTimestampsTextToSpeechRequestOutputFormat `json:"-" url:"output_format,omitempty" explode:"true" serializationStyle:"form" queryParam:"output_format"`
	// You can turn on latency optimizations at some cost of quality. The best possible final latency varies by model. Possible values:
	// 0 - default mode (no latency optimizations)
	// 1 - normal latency optimizations (about 50% of possible latency improvement of option 3)
	// 2 - strong latency optimizations (about 75% of possible latency improvement of option 3)
	// 3 - max latency optimizations
	// 4 - max latency optimizations, but also with text normalizer turned off for even more latency savings (best latency, but can mispronounce eg numbers and dates).
	//
	// Defaults to None.
	//
	OptimizeStreamingLatency *int `json:"-" url:"optimize_streaming_latency,omitempty" explode:"true" serializationStyle:"form" queryParam:"optimize_streaming_latency"`
	// When enable_logging is set to false zero retention mode will be used for the request. This will mean history features are unavailable for this request, including request stitching. Zero retention mode may only be used by enterprise customers.
	EnableLogging *bool `json:"-" url:"enable_logging,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_logging"`
	// Voice ID to be used, you can use https://api.elevenlabs.io/v1/voices to list all the available voices.
	VoiceId string `json:"-" url:"-"`
	// The text that will get converted into speech.
	Text string `json:"text" xml:"text" required:"true"`
	// Identifier of the model that will be used, you can query them using GET /v1/models. The model needs to have support for text to speech, you can check this using the can_do_text_to_speech property.
	ModelId *string `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// Language code (ISO 639-1) used to enforce a language for the model and text normalization. If the model does not support the provided language code, it will be ignored. This parameter is not supported for multilingual_v2 models.
	LanguageCode *string `json:"language_code,omitempty" xml:"language_code,omitempty"`
	// Voice settings overriding stored settings for the given voice. They are applied only on the given request.
	VoiceSettings VoiceSettings `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	// A list of pronunciation dictionary locators (id, version_id) to be applied to the text. They will be applied in order. You may have up to 3 locators per request
	PronunciationDictionaryLocators []*PronunciationDictionaryVersionLocatorRequestModel `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	// If specified, our system will make a best effort to sample deterministically, such that repeated requests with the same seed and parameters should return the same result. Determinism is not guaranteed. Must be integer between 0 and 4294967295.
	Seed *int `json:"seed,omitempty" xml:"seed,omitempty"`
	// The text that came before the text of the current request. Can be used to improve the speech's continuity when concatenating together multiple generations or to influence the speech's continuity in the current generation.
	PreviousText *string `json:"previous_text,omitempty" xml:"previous_text,omitempty"`
	// The text that comes after the text of the current request. Can be used to improve the speech's continuity when concatenating together multiple generations or to influence the speech's continuity in the current generation.
	NextText *string `json:"next_text,omitempty" xml:"next_text,omitempty"`
	// A list of request_id of the samples that were generated before this generation. Can be used to improve the speech's continuity when splitting up a large task into multiple requests. The results will be best when the same model is used across the generations. In case both previous_text and previous_request_ids is send, previous_text will be ignored. A maximum of 3 request_ids can be sent.
	PreviousRequestIds []string `json:"previous_request_ids,omitempty" xml:"previous_request_ids,omitempty"`
	// A list of request_id of the samples that come after this generation. next_request_ids is especially useful for maintaining the speech's continuity when regenerating a sample that has had some audio quality issues. For example, if you have generated 3 speech clips, and you want to improve clip 2, passing the request id of clip 3 as a next_request_id (and that of clip 1 as a previous_request_id) will help maintain natural flow in the combined speech. The results will be best when the same model is used across the generations. In case both next_text and next_request_ids is send, next_text will be ignored. A maximum of 3 request_ids can be sent.
	NextRequestIds []string `json:"next_request_ids,omitempty" xml:"next_request_ids,omitempty"`
	// Whether to use the IVC version of a professional voice. This may improve expressiveness and reduce latency.
	UsePvcAsIvc *bool `json:"use_pvc_as_ivc,omitempty" xml:"use_pvc_as_ivc,omitempty"`
	// This parameter controls text normalization with three modes: 'auto', 'on', and 'off'. When set to 'auto', the system will automatically decide whether to apply text normalization (e.g., spelling out numbers). With 'on', text normalization will always be applied, while with 'off', it will be skipped.
	ApplyTextNormalization *BodyTextToSpeechFullWithTimestampsApplyTextNormalization `json:"apply_text_normalization,omitempty" xml:"apply_text_normalization,omitempty"`
	// This parameter controls language text normalization. This helps with proper pronunciation of text in some supported languages. WARNING: This parameter can heavily increase the latency of the request. Currently only supported for Japanese.
	ApplyLanguageTextNormalization *bool `json:"apply_language_text_normalization,omitempty" xml:"apply_language_text_normalization,omitempty"`
	// Audio effects applied to the generated audio. Unknown IDs return 404 with `param` naming the field. Reverb (`environment_id` with `send_level` above 0) and a non-zero `pan` produce stereo audio and are rejected with 400 for `ulaw_*`, `alaw_*`, and `pcm_*` output formats. Effects don't change the cost of the request.
	AudioEffects *EffectsSpecInput `json:"audio_effects,omitempty" xml:"audio_effects,omitempty"`
}

func (b *BodyTextToSpeechFullWithTimestamps) GetOutputFormat() (value *ConvertWithTimestampsTextToSpeechRequestOutputFormat) {
	if b == nil {
		return
	}
	return b.OutputFormat
}

func (b *BodyTextToSpeechFullWithTimestamps) GetOptimizeStreamingLatency() (value *int) {
	if b == nil {
		return
	}
	return b.OptimizeStreamingLatency
}

func (b *BodyTextToSpeechFullWithTimestamps) GetEnableLogging() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableLogging
}

func (b *BodyTextToSpeechFullWithTimestamps) GetVoiceId() (value string) {
	if b == nil {
		return
	}
	return b.VoiceId
}

func (b *BodyTextToSpeechFullWithTimestamps) GetText() (value string) {
	if b == nil {
		return
	}
	return b.Text
}

func (b *BodyTextToSpeechFullWithTimestamps) GetModelId() (value *string) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyTextToSpeechFullWithTimestamps) GetLanguageCode() (value *string) {
	if b == nil {
		return
	}
	return b.LanguageCode
}

func (b *BodyTextToSpeechFullWithTimestamps) GetVoiceSettings() (value VoiceSettings) {
	if b == nil {
		return
	}
	return b.VoiceSettings
}

func (b *BodyTextToSpeechFullWithTimestamps) GetPronunciationDictionaryLocators() (value []*PronunciationDictionaryVersionLocatorRequestModel) {
	if b == nil {
		return
	}
	return b.PronunciationDictionaryLocators
}

func (b *BodyTextToSpeechFullWithTimestamps) GetSeed() (value *int) {
	if b == nil {
		return
	}
	return b.Seed
}

func (b *BodyTextToSpeechFullWithTimestamps) GetPreviousText() (value *string) {
	if b == nil {
		return
	}
	return b.PreviousText
}

func (b *BodyTextToSpeechFullWithTimestamps) GetNextText() (value *string) {
	if b == nil {
		return
	}
	return b.NextText
}

func (b *BodyTextToSpeechFullWithTimestamps) GetPreviousRequestIds() (value []string) {
	if b == nil {
		return
	}
	return b.PreviousRequestIds
}

func (b *BodyTextToSpeechFullWithTimestamps) GetNextRequestIds() (value []string) {
	if b == nil {
		return
	}
	return b.NextRequestIds
}

func (b *BodyTextToSpeechFullWithTimestamps) GetUsePvcAsIvc() (value *bool) {
	if b == nil {
		return
	}
	return b.UsePvcAsIvc
}

func (b *BodyTextToSpeechFullWithTimestamps) GetApplyTextNormalization() (value *BodyTextToSpeechFullWithTimestampsApplyTextNormalization) {
	if b == nil {
		return
	}
	return b.ApplyTextNormalization
}

func (b *BodyTextToSpeechFullWithTimestamps) GetApplyLanguageTextNormalization() (value *bool) {
	if b == nil {
		return
	}
	return b.ApplyLanguageTextNormalization
}

func (b *BodyTextToSpeechFullWithTimestamps) GetAudioEffects() (value *EffectsSpecInput) {
	if b == nil {
		return
	}
	return b.AudioEffects
}

func (b *BodyTextToSpeechFullWithTimestamps) SetOutputFormat(outputFormat *ConvertWithTimestampsTextToSpeechRequestOutputFormat) {
	b.OutputFormat = outputFormat
}

func (b *BodyTextToSpeechFullWithTimestamps) SetOptimizeStreamingLatency(optimizeStreamingLatency *int) {
	b.OptimizeStreamingLatency = optimizeStreamingLatency
}

func (b *BodyTextToSpeechFullWithTimestamps) SetEnableLogging(enableLogging *bool) {
	b.EnableLogging = enableLogging
}

func (b *BodyTextToSpeechFullWithTimestamps) SetVoiceId(voiceId string) {
	b.VoiceId = voiceId
}

func (b *BodyTextToSpeechFullWithTimestamps) SetText(text string) {
	b.Text = text
}

func (b *BodyTextToSpeechFullWithTimestamps) SetModelId(modelId *string) {
	b.ModelId = modelId
}

func (b *BodyTextToSpeechFullWithTimestamps) SetLanguageCode(languageCode *string) {
	b.LanguageCode = languageCode
}

func (b *BodyTextToSpeechFullWithTimestamps) SetVoiceSettings(voiceSettings VoiceSettings) {
	b.VoiceSettings = voiceSettings
}

func (b *BodyTextToSpeechFullWithTimestamps) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*PronunciationDictionaryVersionLocatorRequestModel) {
	b.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (b *BodyTextToSpeechFullWithTimestamps) SetSeed(seed *int) {
	b.Seed = seed
}

func (b *BodyTextToSpeechFullWithTimestamps) SetPreviousText(previousText *string) {
	b.PreviousText = previousText
}

func (b *BodyTextToSpeechFullWithTimestamps) SetNextText(nextText *string) {
	b.NextText = nextText
}

func (b *BodyTextToSpeechFullWithTimestamps) SetPreviousRequestIds(previousRequestIds []string) {
	b.PreviousRequestIds = previousRequestIds
}

func (b *BodyTextToSpeechFullWithTimestamps) SetNextRequestIds(nextRequestIds []string) {
	b.NextRequestIds = nextRequestIds
}

func (b *BodyTextToSpeechFullWithTimestamps) SetUsePvcAsIvc(usePvcAsIvc *bool) {
	b.UsePvcAsIvc = usePvcAsIvc
}

func (b *BodyTextToSpeechFullWithTimestamps) SetApplyTextNormalization(applyTextNormalization *BodyTextToSpeechFullWithTimestampsApplyTextNormalization) {
	b.ApplyTextNormalization = applyTextNormalization
}

func (b *BodyTextToSpeechFullWithTimestamps) SetApplyLanguageTextNormalization(applyLanguageTextNormalization *bool) {
	b.ApplyLanguageTextNormalization = applyLanguageTextNormalization
}

func (b *BodyTextToSpeechFullWithTimestamps) SetAudioEffects(audioEffects *EffectsSpecInput) {
	b.AudioEffects = audioEffects
}

func (b *BodyTextToSpeechFullWithTimestamps) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyTextToSpeechFullWithTimestamps) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyTextToSpeechFullWithTimestamps to string"
	}
	return string(jsonData)
}

func (b *BodyTextToSpeechFullWithTimestamps) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyTextToSpeechFullWithTimestamps) MarshalJSON() ([]byte, error) {
	type embed BodyTextToSpeechFullWithTimestamps
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
