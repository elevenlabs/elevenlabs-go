package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyTextToSpeechStreamWithTimestamps struct {
	// The text that will get converted into speech.
	Text string `json:"text" xml:"text" required:"true"`
	// Identifier of the model that will be used, you can query them using GET /v1/models. The model needs to have support for text to speech, you can check this using the can_do_text_to_speech property.
	ModelId *string `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// Language code (ISO 639-1) used to enforce a language for the model and text normalization. If the model does not support the provided language code, it will be ignored. This parameter is not supported for multilingual_v2 models.
	LanguageCode *string `json:"language_code,omitempty" xml:"language_code,omitempty"`
	// Voice settings overriding stored settings for the given voice. They are applied only on the given request.
	VoiceSettings *VoiceSettings `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
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
	ApplyTextNormalization *BodyTextToSpeechStreamWithTimestampsApplyTextNormalization `json:"apply_text_normalization,omitempty" xml:"apply_text_normalization,omitempty"`
	// This parameter controls language text normalization. This helps with proper pronunciation of text in some supported languages. WARNING: This parameter can heavily increase the latency of the request. Currently only supported for Japanese.
	ApplyLanguageTextNormalization *bool `json:"apply_language_text_normalization,omitempty" xml:"apply_language_text_normalization,omitempty"`
	// Audio effects applied to the generated audio. Unknown IDs return 404 with `param` naming the field. Reverb (`environment_id` with `send_level` above 0) and a non-zero `pan` produce stereo audio and are rejected with 400 for `ulaw_*`, `alaw_*`, and `pcm_*` output formats. Effects don't change the cost of the request.
	AudioEffects *EffectsSpecInput `json:"audio_effects,omitempty" xml:"audio_effects,omitempty"`
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetText() (value string) {
	if b == nil {
		return
	}
	return b.Text
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetModelId() (value *string) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetLanguageCode() (value *string) {
	if b == nil {
		return
	}
	return b.LanguageCode
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetVoiceSettings() (value *VoiceSettings) {
	if b == nil {
		return
	}
	return b.VoiceSettings
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetPronunciationDictionaryLocators() (value []*PronunciationDictionaryVersionLocatorRequestModel) {
	if b == nil {
		return
	}
	return b.PronunciationDictionaryLocators
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetSeed() (value *int) {
	if b == nil {
		return
	}
	return b.Seed
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetPreviousText() (value *string) {
	if b == nil {
		return
	}
	return b.PreviousText
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetNextText() (value *string) {
	if b == nil {
		return
	}
	return b.NextText
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetPreviousRequestIds() (value []string) {
	if b == nil {
		return
	}
	return b.PreviousRequestIds
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetNextRequestIds() (value []string) {
	if b == nil {
		return
	}
	return b.NextRequestIds
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetUsePvcAsIvc() (value *bool) {
	if b == nil {
		return
	}
	return b.UsePvcAsIvc
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetApplyTextNormalization() (value *BodyTextToSpeechStreamWithTimestampsApplyTextNormalization) {
	if b == nil {
		return
	}
	return b.ApplyTextNormalization
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetApplyLanguageTextNormalization() (value *bool) {
	if b == nil {
		return
	}
	return b.ApplyLanguageTextNormalization
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetAudioEffects() (value *EffectsSpecInput) {
	if b == nil {
		return
	}
	return b.AudioEffects
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetText(text string) {
	b.Text = text
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetModelId(modelId *string) {
	b.ModelId = modelId
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetLanguageCode(languageCode *string) {
	b.LanguageCode = languageCode
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetVoiceSettings(voiceSettings *VoiceSettings) {
	b.VoiceSettings = voiceSettings
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*PronunciationDictionaryVersionLocatorRequestModel) {
	b.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetSeed(seed *int) {
	b.Seed = seed
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetPreviousText(previousText *string) {
	b.PreviousText = previousText
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetNextText(nextText *string) {
	b.NextText = nextText
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetPreviousRequestIds(previousRequestIds []string) {
	b.PreviousRequestIds = previousRequestIds
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetNextRequestIds(nextRequestIds []string) {
	b.NextRequestIds = nextRequestIds
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetUsePvcAsIvc(usePvcAsIvc *bool) {
	b.UsePvcAsIvc = usePvcAsIvc
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetApplyTextNormalization(applyTextNormalization *BodyTextToSpeechStreamWithTimestampsApplyTextNormalization) {
	b.ApplyTextNormalization = applyTextNormalization
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetApplyLanguageTextNormalization(applyLanguageTextNormalization *bool) {
	b.ApplyLanguageTextNormalization = applyLanguageTextNormalization
}

func (b *BodyTextToSpeechStreamWithTimestamps) SetAudioEffects(audioEffects *EffectsSpecInput) {
	b.AudioEffects = audioEffects
}

func (b *BodyTextToSpeechStreamWithTimestamps) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyTextToSpeechStreamWithTimestamps) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyTextToSpeechStreamWithTimestamps to string"
	}
	return string(jsonData)
}

func (b *BodyTextToSpeechStreamWithTimestamps) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyTextToSpeechStreamWithTimestamps) MarshalJSON() ([]byte, error) {
	type embed BodyTextToSpeechStreamWithTimestamps
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
