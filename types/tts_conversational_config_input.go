package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TtsConversationalConfigInput struct {
	ModelId *TtsConversationalModel `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// The voice ID to use for TTS
	VoiceId *string `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	// Additional supported voices for the agent
	SupportedVoices []*SupportedVoice `json:"supported_voices,omitempty" xml:"supported_voices,omitempty"`
	// When enabled, applies expressive audio tags prompt. Automatically disabled for non-v3 models.
	ExpressiveMode *bool `json:"expressive_mode,omitempty" xml:"expressive_mode,omitempty"`
	// Suggested audio tags to boost expressive speech (for eleven_v3 and eleven_v3_conversational models). The agent can still use other tags not listed here.
	SuggestedAudioTags       []*SuggestedAudioTag `json:"suggested_audio_tags,omitempty" xml:"suggested_audio_tags,omitempty" maxItems:"20"`
	AgentOutputAudioFormat   *TtsOutputFormat     `json:"agent_output_audio_format,omitempty" xml:"agent_output_audio_format,omitempty"`
	OptimizeStreamingLatency *int                 `json:"optimize_streaming_latency,omitempty" xml:"optimize_streaming_latency,omitempty"`
	// The stability of generated speech
	Stability *float64 `json:"stability,omitempty" xml:"stability,omitempty" min:"0" max:"1"`
	// The speed of generated speech
	Speed *float64 `json:"speed,omitempty" xml:"speed,omitempty" min:"0.7" max:"1.2"`
	// The similarity boost for generated speech
	SimilarityBoost *float64 `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty" min:"0" max:"1"`
	// The clone strength for generated speech
	CloneStrength *float64 `json:"clone_strength,omitempty" xml:"clone_strength,omitempty" min:"0" max:"1"`
	// Method for converting numbers to words before sending to TTS
	TextNormalisationType *TextNormalisationType `json:"text_normalisation_type,omitempty" xml:"text_normalisation_type,omitempty"`
	// The pronunciation dictionary locators
	PronunciationDictionaryLocators []*PydanticPronunciationDictionaryVersionLocator `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	// Opt-in to SSML phoneme tag handling for V3 models. When enabled, phoneme tags (inline and from pronunciation dictionaries) are parsed into inline IPA before being sent to the model.
	EnablePhonemeTags *bool `json:"enable_phoneme_tags,omitempty" xml:"enable_phoneme_tags,omitempty"`
	// Optional TTS effects spec: filter preset, distance (proximity EQ), and environment (convolution reverb).
	AudioEffects    *EffectsSpecInput `json:"audio_effects,omitempty" xml:"audio_effects,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TtsConversationalConfigInput) GetModelId() (value *TtsConversationalModel) {
	if t == nil {
		return
	}
	return t.ModelId
}

func (t *TtsConversationalConfigInput) GetVoiceId() (value *string) {
	if t == nil {
		return
	}
	return t.VoiceId
}

func (t *TtsConversationalConfigInput) GetSupportedVoices() (value []*SupportedVoice) {
	if t == nil {
		return
	}
	return t.SupportedVoices
}

func (t *TtsConversationalConfigInput) GetExpressiveMode() (value *bool) {
	if t == nil {
		return
	}
	return t.ExpressiveMode
}

func (t *TtsConversationalConfigInput) GetSuggestedAudioTags() (value []*SuggestedAudioTag) {
	if t == nil {
		return
	}
	return t.SuggestedAudioTags
}

func (t *TtsConversationalConfigInput) GetAgentOutputAudioFormat() (value *TtsOutputFormat) {
	if t == nil {
		return
	}
	return t.AgentOutputAudioFormat
}

func (t *TtsConversationalConfigInput) GetOptimizeStreamingLatency() (value *int) {
	if t == nil {
		return
	}
	return t.OptimizeStreamingLatency
}

func (t *TtsConversationalConfigInput) GetStability() (value *float64) {
	if t == nil {
		return
	}
	return t.Stability
}

func (t *TtsConversationalConfigInput) GetSpeed() (value *float64) {
	if t == nil {
		return
	}
	return t.Speed
}

func (t *TtsConversationalConfigInput) GetSimilarityBoost() (value *float64) {
	if t == nil {
		return
	}
	return t.SimilarityBoost
}

func (t *TtsConversationalConfigInput) GetCloneStrength() (value *float64) {
	if t == nil {
		return
	}
	return t.CloneStrength
}

func (t *TtsConversationalConfigInput) GetTextNormalisationType() (value *TextNormalisationType) {
	if t == nil {
		return
	}
	return t.TextNormalisationType
}

func (t *TtsConversationalConfigInput) GetPronunciationDictionaryLocators() (value []*PydanticPronunciationDictionaryVersionLocator) {
	if t == nil {
		return
	}
	return t.PronunciationDictionaryLocators
}

func (t *TtsConversationalConfigInput) GetEnablePhonemeTags() (value *bool) {
	if t == nil {
		return
	}
	return t.EnablePhonemeTags
}

func (t *TtsConversationalConfigInput) GetAudioEffects() (value *EffectsSpecInput) {
	if t == nil {
		return
	}
	return t.AudioEffects
}

func (t *TtsConversationalConfigInput) SetModelId(modelId *TtsConversationalModel) {
	t.ModelId = modelId
}

func (t *TtsConversationalConfigInput) SetVoiceId(voiceId *string) {
	t.VoiceId = voiceId
}

func (t *TtsConversationalConfigInput) SetSupportedVoices(supportedVoices []*SupportedVoice) {
	t.SupportedVoices = supportedVoices
}

func (t *TtsConversationalConfigInput) SetExpressiveMode(expressiveMode *bool) {
	t.ExpressiveMode = expressiveMode
}

func (t *TtsConversationalConfigInput) SetSuggestedAudioTags(suggestedAudioTags []*SuggestedAudioTag) {
	t.SuggestedAudioTags = suggestedAudioTags
}

func (t *TtsConversationalConfigInput) SetAgentOutputAudioFormat(agentOutputAudioFormat *TtsOutputFormat) {
	t.AgentOutputAudioFormat = agentOutputAudioFormat
}

func (t *TtsConversationalConfigInput) SetOptimizeStreamingLatency(optimizeStreamingLatency *int) {
	t.OptimizeStreamingLatency = optimizeStreamingLatency
}

func (t *TtsConversationalConfigInput) SetStability(stability *float64) {
	t.Stability = stability
}

func (t *TtsConversationalConfigInput) SetSpeed(speed *float64) {
	t.Speed = speed
}

func (t *TtsConversationalConfigInput) SetSimilarityBoost(similarityBoost *float64) {
	t.SimilarityBoost = similarityBoost
}

func (t *TtsConversationalConfigInput) SetCloneStrength(cloneStrength *float64) {
	t.CloneStrength = cloneStrength
}

func (t *TtsConversationalConfigInput) SetTextNormalisationType(textNormalisationType *TextNormalisationType) {
	t.TextNormalisationType = textNormalisationType
}

func (t *TtsConversationalConfigInput) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*PydanticPronunciationDictionaryVersionLocator) {
	t.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (t *TtsConversationalConfigInput) SetEnablePhonemeTags(enablePhonemeTags *bool) {
	t.EnablePhonemeTags = enablePhonemeTags
}

func (t *TtsConversationalConfigInput) SetAudioEffects(audioEffects *EffectsSpecInput) {
	t.AudioEffects = audioEffects
}

func (t *TtsConversationalConfigInput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TtsConversationalConfigInput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TtsConversationalConfigInput to string"
	}
	return string(jsonData)
}

func (t *TtsConversationalConfigInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, t); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TtsConversationalConfigInput) MarshalJSON() ([]byte, error) {
	type embed TtsConversationalConfigInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
