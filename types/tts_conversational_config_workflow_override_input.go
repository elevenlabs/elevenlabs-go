package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TtsConversationalConfigWorkflowOverrideInput struct {
	// The model to use for TTS
	ModelId *TtsConversationalModel `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// The voice ID to use for TTS
	VoiceId *string `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	// Additional supported voices for the agent
	SupportedVoices []*SupportedVoice `json:"supported_voices,omitempty" xml:"supported_voices,omitempty"`
	// When enabled, applies expressive audio tags prompt. Automatically disabled for non-v3 models.
	ExpressiveMode *bool `json:"expressive_mode,omitempty" xml:"expressive_mode,omitempty"`
	// Suggested audio tags to boost expressive speech (for eleven_v3 and eleven_v3_conversational models). The agent can still use other tags not listed here.
	SuggestedAudioTags []*SuggestedAudioTag `json:"suggested_audio_tags,omitempty" xml:"suggested_audio_tags,omitempty"`
	// The audio format to use for TTS
	AgentOutputAudioFormat *TtsOutputFormat `json:"agent_output_audio_format,omitempty" xml:"agent_output_audio_format,omitempty"`
	// Deprecated: this field is a no-op and is ignored.
	OptimizeStreamingLatency *int `json:"optimize_streaming_latency,omitempty" xml:"optimize_streaming_latency,omitempty"`
	// The stability of generated speech
	Stability *float64 `json:"stability,omitempty" xml:"stability,omitempty"`
	// The speed of generated speech
	Speed *float64 `json:"speed,omitempty" xml:"speed,omitempty"`
	// The similarity boost for generated speech
	SimilarityBoost *float64 `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty"`
	// The clone strength for generated speech
	CloneStrength *float64 `json:"clone_strength,omitempty" xml:"clone_strength,omitempty"`
	// Method for converting numbers to words before converting text to speech. If set to SYSTEM_PROMPT, the system prompt will be updated to include normalization instructions. If set to ELEVENLABS, the text will be normalized after generation, incurring slight additional latency.
	TextNormalisationType *TextNormalisationType `json:"text_normalisation_type,omitempty" xml:"text_normalisation_type,omitempty"`
	// The pronunciation dictionary locators
	PronunciationDictionaryLocators []*PydanticPronunciationDictionaryVersionLocator `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	// Opt-in to SSML phoneme tag handling for V3 models. When enabled, phoneme tags (inline and from pronunciation dictionaries) are parsed into inline IPA before being sent to the model.
	EnablePhonemeTags *bool `json:"enable_phoneme_tags,omitempty" xml:"enable_phoneme_tags,omitempty"`
	// Optional TTS effects spec: filter preset, distance (proximity EQ), and environment (convolution reverb).
	AudioEffects    *EffectsSpecInput `json:"audio_effects,omitempty" xml:"audio_effects,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetModelId() (value *TtsConversationalModel) {
	if t == nil {
		return
	}
	return t.ModelId
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetVoiceId() (value *string) {
	if t == nil {
		return
	}
	return t.VoiceId
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetSupportedVoices() (value []*SupportedVoice) {
	if t == nil {
		return
	}
	return t.SupportedVoices
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetExpressiveMode() (value *bool) {
	if t == nil {
		return
	}
	return t.ExpressiveMode
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetSuggestedAudioTags() (value []*SuggestedAudioTag) {
	if t == nil {
		return
	}
	return t.SuggestedAudioTags
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetAgentOutputAudioFormat() (value *TtsOutputFormat) {
	if t == nil {
		return
	}
	return t.AgentOutputAudioFormat
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetOptimizeStreamingLatency() (value *int) {
	if t == nil {
		return
	}
	return t.OptimizeStreamingLatency
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetStability() (value *float64) {
	if t == nil {
		return
	}
	return t.Stability
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetSpeed() (value *float64) {
	if t == nil {
		return
	}
	return t.Speed
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetSimilarityBoost() (value *float64) {
	if t == nil {
		return
	}
	return t.SimilarityBoost
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetCloneStrength() (value *float64) {
	if t == nil {
		return
	}
	return t.CloneStrength
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetTextNormalisationType() (value *TextNormalisationType) {
	if t == nil {
		return
	}
	return t.TextNormalisationType
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetPronunciationDictionaryLocators() (value []*PydanticPronunciationDictionaryVersionLocator) {
	if t == nil {
		return
	}
	return t.PronunciationDictionaryLocators
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetEnablePhonemeTags() (value *bool) {
	if t == nil {
		return
	}
	return t.EnablePhonemeTags
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetAudioEffects() (value *EffectsSpecInput) {
	if t == nil {
		return
	}
	return t.AudioEffects
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetModelId(modelId *TtsConversationalModel) {
	t.ModelId = modelId
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetVoiceId(voiceId *string) {
	t.VoiceId = voiceId
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetSupportedVoices(supportedVoices []*SupportedVoice) {
	t.SupportedVoices = supportedVoices
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetExpressiveMode(expressiveMode *bool) {
	t.ExpressiveMode = expressiveMode
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetSuggestedAudioTags(suggestedAudioTags []*SuggestedAudioTag) {
	t.SuggestedAudioTags = suggestedAudioTags
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetAgentOutputAudioFormat(agentOutputAudioFormat *TtsOutputFormat) {
	t.AgentOutputAudioFormat = agentOutputAudioFormat
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetOptimizeStreamingLatency(optimizeStreamingLatency *int) {
	t.OptimizeStreamingLatency = optimizeStreamingLatency
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetStability(stability *float64) {
	t.Stability = stability
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetSpeed(speed *float64) {
	t.Speed = speed
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetSimilarityBoost(similarityBoost *float64) {
	t.SimilarityBoost = similarityBoost
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetCloneStrength(cloneStrength *float64) {
	t.CloneStrength = cloneStrength
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetTextNormalisationType(textNormalisationType *TextNormalisationType) {
	t.TextNormalisationType = textNormalisationType
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*PydanticPronunciationDictionaryVersionLocator) {
	t.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetEnablePhonemeTags(enablePhonemeTags *bool) {
	t.EnablePhonemeTags = enablePhonemeTags
}

func (t *TtsConversationalConfigWorkflowOverrideInput) SetAudioEffects(audioEffects *EffectsSpecInput) {
	t.AudioEffects = audioEffects
}

func (t *TtsConversationalConfigWorkflowOverrideInput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TtsConversationalConfigWorkflowOverrideInput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TtsConversationalConfigWorkflowOverrideInput to string"
	}
	return string(jsonData)
}

func (t *TtsConversationalConfigWorkflowOverrideInput) UnmarshalJSON(data []byte) error {
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

func (t *TtsConversationalConfigWorkflowOverrideInput) MarshalJSON() ([]byte, error) {
	type embed TtsConversationalConfigWorkflowOverrideInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
