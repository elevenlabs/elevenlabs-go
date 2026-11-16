package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Aggregated TTS usage for a conversation (analytics-only, not billing).
type ConversationTtsUsageModel struct {
	PrimaryTtsModel         *string                        `json:"primary_tts_model,omitempty" xml:"primary_tts_model,omitempty"`
	TotalAudioOutputSeconds *float64                       `json:"total_audio_output_seconds,omitempty" xml:"total_audio_output_seconds,omitempty"`
	TotalCharacters         *int                           `json:"total_characters,omitempty" xml:"total_characters,omitempty"`
	PerVoiceUsage           []*ConversationVoiceUsageModel `json:"per_voice_usage,omitempty" xml:"per_voice_usage,omitempty"`
	extraProperties         map[string]interface{}
}

func (c *ConversationTtsUsageModel) GetPrimaryTtsModel() (value *string) {
	if c == nil {
		return
	}
	return c.PrimaryTtsModel
}

func (c *ConversationTtsUsageModel) GetTotalAudioOutputSeconds() (value *float64) {
	if c == nil {
		return
	}
	return c.TotalAudioOutputSeconds
}

func (c *ConversationTtsUsageModel) GetTotalCharacters() (value *int) {
	if c == nil {
		return
	}
	return c.TotalCharacters
}

func (c *ConversationTtsUsageModel) GetPerVoiceUsage() (value []*ConversationVoiceUsageModel) {
	if c == nil {
		return
	}
	return c.PerVoiceUsage
}

func (c *ConversationTtsUsageModel) SetPrimaryTtsModel(primaryTtsModel *string) {
	c.PrimaryTtsModel = primaryTtsModel
}

func (c *ConversationTtsUsageModel) SetTotalAudioOutputSeconds(totalAudioOutputSeconds *float64) {
	c.TotalAudioOutputSeconds = totalAudioOutputSeconds
}

func (c *ConversationTtsUsageModel) SetTotalCharacters(totalCharacters *int) {
	c.TotalCharacters = totalCharacters
}

func (c *ConversationTtsUsageModel) SetPerVoiceUsage(perVoiceUsage []*ConversationVoiceUsageModel) {
	c.PerVoiceUsage = perVoiceUsage
}

func (c *ConversationTtsUsageModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationTtsUsageModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationTtsUsageModel to string"
	}
	return string(jsonData)
}

func (c *ConversationTtsUsageModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationTtsUsageModel) MarshalJSON() ([]byte, error) {
	type embed ConversationTtsUsageModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
