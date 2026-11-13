package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TtsConversationalConfigWorkflowOverrideOutput struct {
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
	AudioEffects    *EffectsSpecOutput `json:"audio_effects,omitempty" xml:"audio_effects,omitempty"`
	extraProperties map[string]interface{}
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetModelId() (value *TtsConversationalModel) {
	if t == nil {
		return
	}
	return t.ModelId
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetVoiceId() (value *string) {
	if t == nil {
		return
	}
	return t.VoiceId
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetSupportedVoices() (value []*SupportedVoice) {
	if t == nil {
		return
	}
	return t.SupportedVoices
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetExpressiveMode() (value *bool) {
	if t == nil {
		return
	}
	return t.ExpressiveMode
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetSuggestedAudioTags() (value []*SuggestedAudioTag) {
	if t == nil {
		return
	}
	return t.SuggestedAudioTags
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetAgentOutputAudioFormat() (value *TtsOutputFormat) {
	if t == nil {
		return
	}
	return t.AgentOutputAudioFormat
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetOptimizeStreamingLatency() (value *int) {
	if t == nil {
		return
	}
	return t.OptimizeStreamingLatency
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetStability() (value *float64) {
	if t == nil {
		return
	}
	return t.Stability
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetSpeed() (value *float64) {
	if t == nil {
		return
	}
	return t.Speed
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetSimilarityBoost() (value *float64) {
	if t == nil {
		return
	}
	return t.SimilarityBoost
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetCloneStrength() (value *float64) {
	if t == nil {
		return
	}
	return t.CloneStrength
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetTextNormalisationType() (value *TextNormalisationType) {
	if t == nil {
		return
	}
	return t.TextNormalisationType
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetPronunciationDictionaryLocators() (value []*PydanticPronunciationDictionaryVersionLocator) {
	if t == nil {
		return
	}
	return t.PronunciationDictionaryLocators
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetEnablePhonemeTags() (value *bool) {
	if t == nil {
		return
	}
	return t.EnablePhonemeTags
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetAudioEffects() (value *EffectsSpecOutput) {
	if t == nil {
		return
	}
	return t.AudioEffects
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetModelId(modelId *TtsConversationalModel) {
	t.ModelId = modelId
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetVoiceId(voiceId *string) {
	t.VoiceId = voiceId
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetSupportedVoices(supportedVoices []*SupportedVoice) {
	t.SupportedVoices = supportedVoices
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetExpressiveMode(expressiveMode *bool) {
	t.ExpressiveMode = expressiveMode
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetSuggestedAudioTags(suggestedAudioTags []*SuggestedAudioTag) {
	t.SuggestedAudioTags = suggestedAudioTags
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetAgentOutputAudioFormat(agentOutputAudioFormat *TtsOutputFormat) {
	t.AgentOutputAudioFormat = agentOutputAudioFormat
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetOptimizeStreamingLatency(optimizeStreamingLatency *int) {
	t.OptimizeStreamingLatency = optimizeStreamingLatency
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetStability(stability *float64) {
	t.Stability = stability
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetSpeed(speed *float64) {
	t.Speed = speed
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetSimilarityBoost(similarityBoost *float64) {
	t.SimilarityBoost = similarityBoost
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetCloneStrength(cloneStrength *float64) {
	t.CloneStrength = cloneStrength
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetTextNormalisationType(textNormalisationType *TextNormalisationType) {
	t.TextNormalisationType = textNormalisationType
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*PydanticPronunciationDictionaryVersionLocator) {
	t.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetEnablePhonemeTags(enablePhonemeTags *bool) {
	t.EnablePhonemeTags = enablePhonemeTags
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) SetAudioEffects(audioEffects *EffectsSpecOutput) {
	t.AudioEffects = audioEffects
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TtsConversationalConfigWorkflowOverrideOutput) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TtsConversationalConfigWorkflowOverrideOutput to string"
	}
	return string(jsonData)
}

func (t *TtsConversationalConfigWorkflowOverrideOutput) UnmarshalJSON(data []byte) error {
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

func (t *TtsConversationalConfigWorkflowOverrideOutput) MarshalJSON() ([]byte, error) {
	type embed TtsConversationalConfigWorkflowOverrideOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
