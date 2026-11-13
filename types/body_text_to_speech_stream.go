package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyTextToSpeechStream struct {
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
	ApplyTextNormalization *BodyTextToSpeechStreamApplyTextNormalization `json:"apply_text_normalization,omitempty" xml:"apply_text_normalization,omitempty"`
	// This parameter controls language text normalization. This helps with proper pronunciation of text in some supported languages. WARNING: This parameter can heavily increase the latency of the request. Currently only supported for Japanese.
	ApplyLanguageTextNormalization *bool `json:"apply_language_text_normalization,omitempty" xml:"apply_language_text_normalization,omitempty"`
	// Audio effects applied to the generated audio. Unknown IDs return 404 with `param` naming the field. Reverb (`environment_id` with `send_level` above 0) and a non-zero `pan` produce stereo audio and are rejected with 400 for `ulaw_*`, `alaw_*`, and `pcm_*` output formats. Effects don't change the cost of the request.
	AudioEffects *EffectsSpecInput `json:"audio_effects,omitempty" xml:"audio_effects,omitempty"`
}

func (b *BodyTextToSpeechStream) GetText() (value string) {
	if b == nil {
		return
	}
	return b.Text
}

func (b *BodyTextToSpeechStream) GetModelId() (value *string) {
	if b == nil {
		return
	}
	return b.ModelId
}

func (b *BodyTextToSpeechStream) GetLanguageCode() (value *string) {
	if b == nil {
		return
	}
	return b.LanguageCode
}

func (b *BodyTextToSpeechStream) GetVoiceSettings() (value VoiceSettings) {
	if b == nil {
		return
	}
	return b.VoiceSettings
}

func (b *BodyTextToSpeechStream) GetPronunciationDictionaryLocators() (value []*PronunciationDictionaryVersionLocatorRequestModel) {
	if b == nil {
		return
	}
	return b.PronunciationDictionaryLocators
}

func (b *BodyTextToSpeechStream) GetSeed() (value *int) {
	if b == nil {
		return
	}
	return b.Seed
}

func (b *BodyTextToSpeechStream) GetPreviousText() (value *string) {
	if b == nil {
		return
	}
	return b.PreviousText
}

func (b *BodyTextToSpeechStream) GetNextText() (value *string) {
	if b == nil {
		return
	}
	return b.NextText
}

func (b *BodyTextToSpeechStream) GetPreviousRequestIds() (value []string) {
	if b == nil {
		return
	}
	return b.PreviousRequestIds
}

func (b *BodyTextToSpeechStream) GetNextRequestIds() (value []string) {
	if b == nil {
		return
	}
	return b.NextRequestIds
}

func (b *BodyTextToSpeechStream) GetUsePvcAsIvc() (value *bool) {
	if b == nil {
		return
	}
	return b.UsePvcAsIvc
}

func (b *BodyTextToSpeechStream) GetApplyTextNormalization() (value *BodyTextToSpeechStreamApplyTextNormalization) {
	if b == nil {
		return
	}
	return b.ApplyTextNormalization
}

func (b *BodyTextToSpeechStream) GetApplyLanguageTextNormalization() (value *bool) {
	if b == nil {
		return
	}
	return b.ApplyLanguageTextNormalization
}

func (b *BodyTextToSpeechStream) GetAudioEffects() (value *EffectsSpecInput) {
	if b == nil {
		return
	}
	return b.AudioEffects
}

func (b *BodyTextToSpeechStream) SetText(text string) {
	b.Text = text
}

func (b *BodyTextToSpeechStream) SetModelId(modelId *string) {
	b.ModelId = modelId
}

func (b *BodyTextToSpeechStream) SetLanguageCode(languageCode *string) {
	b.LanguageCode = languageCode
}

func (b *BodyTextToSpeechStream) SetVoiceSettings(voiceSettings VoiceSettings) {
	b.VoiceSettings = voiceSettings
}

func (b *BodyTextToSpeechStream) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*PronunciationDictionaryVersionLocatorRequestModel) {
	b.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (b *BodyTextToSpeechStream) SetSeed(seed *int) {
	b.Seed = seed
}

func (b *BodyTextToSpeechStream) SetPreviousText(previousText *string) {
	b.PreviousText = previousText
}

func (b *BodyTextToSpeechStream) SetNextText(nextText *string) {
	b.NextText = nextText
}

func (b *BodyTextToSpeechStream) SetPreviousRequestIds(previousRequestIds []string) {
	b.PreviousRequestIds = previousRequestIds
}

func (b *BodyTextToSpeechStream) SetNextRequestIds(nextRequestIds []string) {
	b.NextRequestIds = nextRequestIds
}

func (b *BodyTextToSpeechStream) SetUsePvcAsIvc(usePvcAsIvc *bool) {
	b.UsePvcAsIvc = usePvcAsIvc
}

func (b *BodyTextToSpeechStream) SetApplyTextNormalization(applyTextNormalization *BodyTextToSpeechStreamApplyTextNormalization) {
	b.ApplyTextNormalization = applyTextNormalization
}

func (b *BodyTextToSpeechStream) SetApplyLanguageTextNormalization(applyLanguageTextNormalization *bool) {
	b.ApplyLanguageTextNormalization = applyLanguageTextNormalization
}

func (b *BodyTextToSpeechStream) SetAudioEffects(audioEffects *EffectsSpecInput) {
	b.AudioEffects = audioEffects
}

func (b *BodyTextToSpeechStream) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyTextToSpeechStream) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyTextToSpeechStream to string"
	}
	return string(jsonData)
}

func (b *BodyTextToSpeechStream) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyTextToSpeechStream) MarshalJSON() ([]byte, error) {
	type embed BodyTextToSpeechStream
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
