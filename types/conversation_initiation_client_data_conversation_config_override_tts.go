package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

// Configuration for conversational text to speech
type ConversationInitiationClientDataConversationConfigOverrideTts struct {
	// The model to use for TTS
	ModelId *ConversationInitiationClientDataConversationConfigOverrideTtsModelId `json:"model_id,omitempty" xml:"model_id,omitempty"`
	// The voice ID to use for TTS
	VoiceId *string `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	// Additional supported voices for the agent
	SupportedVoices []*ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem `json:"supported_voices,omitempty" xml:"supported_voices,omitempty"`
	// The stability of generated speech
	Stability *float64 `json:"stability,omitempty" xml:"stability,omitempty"`
	// The speed of generated speech
	Speed *float64 `json:"speed,omitempty" xml:"speed,omitempty"`
	// The similarity boost for generated speech
	SimilarityBoost *float64 `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty"`
	// The clone strength for generated speech
	CloneStrength *float64 `json:"clone_strength,omitempty" xml:"clone_strength,omitempty"`
	// The pronunciation dictionary locators
	PronunciationDictionaryLocators []*InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem `json:"pronunciation_dictionary_locators,omitempty" xml:"pronunciation_dictionary_locators,omitempty"`
	extraProperties                 map[string]interface{}
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) GetModelId() (value *ConversationInitiationClientDataConversationConfigOverrideTtsModelId) {
	if c == nil {
		return
	}
	return c.ModelId
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) GetVoiceId() (value *string) {
	if c == nil {
		return
	}
	return c.VoiceId
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) GetSupportedVoices() (value []*ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) {
	if c == nil {
		return
	}
	return c.SupportedVoices
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) GetStability() (value *float64) {
	if c == nil {
		return
	}
	return c.Stability
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) GetSpeed() (value *float64) {
	if c == nil {
		return
	}
	return c.Speed
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) GetSimilarityBoost() (value *float64) {
	if c == nil {
		return
	}
	return c.SimilarityBoost
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) GetCloneStrength() (value *float64) {
	if c == nil {
		return
	}
	return c.CloneStrength
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) GetPronunciationDictionaryLocators() (value []*InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) {
	if c == nil {
		return
	}
	return c.PronunciationDictionaryLocators
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) SetModelId(modelId *ConversationInitiationClientDataConversationConfigOverrideTtsModelId) {
	c.ModelId = modelId
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) SetVoiceId(voiceId *string) {
	c.VoiceId = voiceId
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) SetSupportedVoices(supportedVoices []*ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) {
	c.SupportedVoices = supportedVoices
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) SetStability(stability *float64) {
	c.Stability = stability
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) SetSpeed(speed *float64) {
	c.Speed = speed
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) SetSimilarityBoost(similarityBoost *float64) {
	c.SimilarityBoost = similarityBoost
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) SetCloneStrength(cloneStrength *float64) {
	c.CloneStrength = cloneStrength
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) SetPronunciationDictionaryLocators(pronunciationDictionaryLocators []*InitiationClientDataConversationConfigOverrideTtsPronunciationDictionaryLocatorsItem) {
	c.PronunciationDictionaryLocators = pronunciationDictionaryLocators
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataConversationConfigOverrideTts) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataConversationConfigOverrideTts to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationInitiationClientDataConversationConfigOverrideTts
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataConversationConfigOverrideTts(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTts) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataConversationConfigOverrideTts
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
