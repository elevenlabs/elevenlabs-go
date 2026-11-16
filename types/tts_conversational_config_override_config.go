package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type TtsConversationalConfigOverrideConfig struct {
	// Whether to allow overriding the model_id field.
	ModelId *bool `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// Whether to allow overriding the voice_id field.
	VoiceId *bool `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	// Whether to allow overriding the supported_voices field.
	SupportedVoices *bool `json:"supported_voices,omitempty" xml:"supported_voices,omitempty"`
	// Whether to allow overriding the stability field.
	Stability *bool `json:"stability,omitempty" xml:"stability,omitempty"`
	// Whether to allow overriding the speed field.
	Speed *bool `json:"speed,omitempty" xml:"speed,omitempty"`
	// Whether to allow overriding the similarity_boost field.
	SimilarityBoost *bool `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty"`
	// Whether to allow overriding the clone_strength field.
	CloneStrength *bool `json:"clone_strength,omitempty" xml:"clone_strength,omitempty"`
	// Whether to allow overriding the pronunciation_dictionary_locators field.
	PronunciationDictionaryLocators *bool `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	extraProperties                 map[string]interface{}
}

func (t *TtsConversationalConfigOverrideConfig) GetModelId() (value *bool) {
	if t == nil {
		return
	}
	return t.ModelId
}

func (t *TtsConversationalConfigOverrideConfig) GetVoiceId() (value *bool) {
	if t == nil {
		return
	}
	return t.VoiceId
}

func (t *TtsConversationalConfigOverrideConfig) GetSupportedVoices() (value *bool) {
	if t == nil {
		return
	}
	return t.SupportedVoices
}

func (t *TtsConversationalConfigOverrideConfig) GetStability() (value *bool) {
	if t == nil {
		return
	}
	return t.Stability
}

func (t *TtsConversationalConfigOverrideConfig) GetSpeed() (value *bool) {
	if t == nil {
		return
	}
	return t.Speed
}

func (t *TtsConversationalConfigOverrideConfig) GetSimilarityBoost() (value *bool) {
	if t == nil {
		return
	}
	return t.SimilarityBoost
}

func (t *TtsConversationalConfigOverrideConfig) GetCloneStrength() (value *bool) {
	if t == nil {
		return
	}
	return t.CloneStrength
}

func (t *TtsConversationalConfigOverrideConfig) GetPronunciationDictionaryLocators() (value *bool) {
	if t == nil {
		return
	}
	return t.PronunciationDictionaryLocators
}

func (t *TtsConversationalConfigOverrideConfig) SetModelId(modelId *bool) {
	t.ModelId = modelId
}

func (t *TtsConversationalConfigOverrideConfig) SetVoiceId(voiceId *bool) {
	t.VoiceId = voiceId
}

func (t *TtsConversationalConfigOverrideConfig) SetSupportedVoices(supportedVoices *bool) {
	t.SupportedVoices = supportedVoices
}

func (t *TtsConversationalConfigOverrideConfig) SetStability(stability *bool) {
	t.Stability = stability
}

func (t *TtsConversationalConfigOverrideConfig) SetSpeed(speed *bool) {
	t.Speed = speed
}

func (t *TtsConversationalConfigOverrideConfig) SetSimilarityBoost(similarityBoost *bool) {
	t.SimilarityBoost = similarityBoost
}

func (t *TtsConversationalConfigOverrideConfig) SetCloneStrength(cloneStrength *bool) {
	t.CloneStrength = cloneStrength
}

func (t *TtsConversationalConfigOverrideConfig) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators *bool) {
	t.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (t *TtsConversationalConfigOverrideConfig) GetExtraProperties() map[string]interface{} {
	if t == nil {
		return nil
	}
	return t.extraProperties
}

func (t TtsConversationalConfigOverrideConfig) String() string {
	jsonData, err := json.MarshalIndent(t, "", "  ")
	if err != nil {
		return "error converting struct: TtsConversationalConfigOverrideConfig to string"
	}
	return string(jsonData)
}

func (t *TtsConversationalConfigOverrideConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler TtsConversationalConfigOverrideConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TtsConversationalConfigOverrideConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *t)
	if err != nil {
		return err
	}
	t.extraProperties = extraProperties
	return nil
}

func (t *TtsConversationalConfigOverrideConfig) MarshalJSON() ([]byte, error) {
	type embed TtsConversationalConfigOverrideConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*t),
	}
	return json.Marshal(marshaler)
}
