package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type TtsConversationalConfigOverride struct {
	// The model to use for TTS
	ModelId *TtsConversationalModel `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// The voice ID to use for TTS
	VoiceId *string `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	// Additional supported voices for the agent
	SupportedVoices []*SupportedVoice `json:"supported_voices,omitempty" xml:"supported_voices,omitempty"`
	// The stability of generated speech
	Stability *float64 `json:"stability,omitempty" xml:"stability,omitempty"`
	// The speed of generated speech
	Speed *float64 `json:"speed,omitempty" xml:"speed,omitempty"`
	// The similarity boost for generated speech
	SimilarityBoost *float64 `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty"`
	// The clone strength for generated speech
	CloneStrength *float64 `json:"clone_strength,omitempty" xml:"clone_strength,omitempty"`
	// The pronunciation dictionary locators
	PronunciationDictionaryLocators []*PydanticPronunciationDictionaryVersionLocator `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	extraProperties                 map[string]interface{}
}

func (t *TtsConversationalConfigOverride) GetModelId() (value *TtsConversationalModel) {
	if t == nil {
		return
	}
	return t.ModelId
}

func (t *TtsConversationalConfigOverride) GetVoiceId() (value *string) {
	if t == nil {
		return
	}
	return t.VoiceId
}

func (t *TtsConversationalConfigOverride) GetSupportedVoices() (value []*SupportedVoice) {
	if t == nil {
		return
	}
	return t.SupportedVoices
}

func (t *TtsConversationalConfigOverride) GetStability() (value *float64) {
	if t == nil {
		return
	}
	return t.Stability
}

func (t *TtsConversationalConfigOverride) GetSpeed() (value *float64) {
	if t == nil {
		return
	}
	return t.Speed
}

func (t *TtsConversationalConfigOverride) GetSimilarityBoost() (value *float64) {
	if t == nil {
		return
	}
	return t.SimilarityBoost
}

func (t *TtsConversationalConfigOverride) GetCloneStrength() (value *float64) {
	if t == nil {
		return
	}
	return t.CloneStrength
}

func (t *TtsConversationalConfigOverride) GetPronunciationDictionaryLocators() (value []*PydanticPronunciationDictionaryVersionLocator) {
	if t == nil {
		return
	}
	return t.PronunciationDictionaryLocators
}

func (t *TtsConversationalConfigOverride) SetModelId(modelId *TtsConversationalModel) {
	t.ModelId = modelId
}

func (t *TtsConversationalConfigOverride) SetVoiceId(voiceId *string) {
	t.VoiceId = voiceId
}

func (t *TtsConversationalConfigOverride) SetSupportedVoices(supportedVoices []*SupportedVoice) {
	t.SupportedVoices = supportedVoices
}

func (t *TtsConversationalConfigOverride) SetStability(stability *float64) {
	t.Stability = stability
}

func (t *TtsConversationalConfigOverride) SetSpeed(speed *float64) {
	t.Speed = speed
}

func (t *TtsConversationalConfigOverride) SetSimilarityBoost(similarityBoost *float64) {
	t.SimilarityBoost = similarityBoost
}

func (t *TtsConversationalConfigOverride) SetCloneStrength(cloneStrength *float64) {
	t.CloneStrength = cloneStrength
}

func (t *TtsConversationalConfigOverride) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*PydanticPronunciationDictionaryVersionLocator) {
	t.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (t *TtsConversationalConfigOverride) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TtsConversationalConfigOverride) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TtsConversationalConfigOverride to string"
	}
	return string(jsonData)
}

func (t *TtsConversationalConfigOverride) UnmarshalJSON(data []byte) error {
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

func (t *TtsConversationalConfigOverride) MarshalJSON() ([]byte, error) {
	type embed TtsConversationalConfigOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
