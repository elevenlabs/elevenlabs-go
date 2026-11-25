package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem struct {
	Label                    string                                                                                       `json:"label" xml:"label" required:"true" minLength:"1"`
	VoiceId                  string                                                                                       `json:"voice_id" xml:"voice_id" required:"true" minLength:"1"`
	Description              *string                                                                                      `json:"description,omitempty" xml:"description,omitempty"`
	Language                 *string                                                                                      `json:"language,omitempty" xml:"language,omitempty"`
	ModelFamily              *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily `json:"model_family,omitempty" xml:"model_family,omitempty"`
	OptimizeStreamingLatency *int                                                                                         `json:"optimize_streaming_latency,omitempty" xml:"optimize_streaming_latency,omitempty"`
	Stability                *float64                                                                                     `json:"stability,omitempty" xml:"stability,omitempty" min:"0" max:"1"`
	Speed                    *float64                                                                                     `json:"speed,omitempty" xml:"speed,omitempty" min:"0.7" max:"1.2"`
	SimilarityBoost          *float64                                                                                     `json:"similarity_boost,omitempty" xml:"similarity_boost,omitempty" min:"0" max:"1"`
	CloneStrength            *float64                                                                                     `json:"clone_strength,omitempty" xml:"clone_strength,omitempty" min:"0" max:"1"`
	extraProperties          map[string]interface{}
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) GetLabel() (value string) {
	if c == nil {
		return
	}
	return c.Label
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) GetVoiceId() (value string) {
	if c == nil {
		return
	}
	return c.VoiceId
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) GetDescription() (value *string) {
	if c == nil {
		return
	}
	return c.Description
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) GetLanguage() (value *string) {
	if c == nil {
		return
	}
	return c.Language
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) GetModelFamily() (value *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily) {
	if c == nil {
		return
	}
	return c.ModelFamily
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) GetOptimizeStreamingLatency() (value *int) {
	if c == nil {
		return
	}
	return c.OptimizeStreamingLatency
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) GetStability() (value *float64) {
	if c == nil {
		return
	}
	return c.Stability
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) GetSpeed() (value *float64) {
	if c == nil {
		return
	}
	return c.Speed
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) GetSimilarityBoost() (value *float64) {
	if c == nil {
		return
	}
	return c.SimilarityBoost
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) GetCloneStrength() (value *float64) {
	if c == nil {
		return
	}
	return c.CloneStrength
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) SetLabel(label string) {
	c.Label = label
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) SetVoiceId(voiceId string) {
	c.VoiceId = voiceId
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) SetDescription(description *string) {
	c.Description = description
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) SetLanguage(language *string) {
	c.Language = language
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) SetModelFamily(modelFamily *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItemModelFamily) {
	c.ModelFamily = modelFamily
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) SetOptimizeStreamingLatency(optimizeStreamingLatency *int) {
	c.OptimizeStreamingLatency = optimizeStreamingLatency
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) SetStability(stability *float64) {
	c.Stability = stability
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) SetSpeed(speed *float64) {
	c.Speed = speed
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) SetSimilarityBoost(similarityBoost *float64) {
	c.SimilarityBoost = similarityBoost
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) SetCloneStrength(cloneStrength *float64) {
	c.CloneStrength = cloneStrength
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataConversationConfigOverrideTtsSupportedVoicesItem
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
