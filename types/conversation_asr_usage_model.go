package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Aggregated ASR usage for a conversation (analytics-only, not billing).
type ConversationAsrUsageModel struct {
	AsrModel                *string  `json:"asr_model,omitempty" xml:"asr_model,omitempty"`
	TotalTranscriptionCalls *int     `json:"total_transcription_calls,omitempty" xml:"total_transcription_calls,omitempty"`
	TotalAudioInputSeconds  *float64 `json:"total_audio_input_seconds,omitempty" xml:"total_audio_input_seconds,omitempty"`
	extraProperties         map[string]interface{}
}

func (c *ConversationAsrUsageModel) GetAsrModel() (value *string) {
	if c == nil {
		return
	}
	return c.AsrModel
}

func (c *ConversationAsrUsageModel) GetTotalTranscriptionCalls() (value *int) {
	if c == nil {
		return
	}
	return c.TotalTranscriptionCalls
}

func (c *ConversationAsrUsageModel) GetTotalAudioInputSeconds() (value *float64) {
	if c == nil {
		return
	}
	return c.TotalAudioInputSeconds
}

func (c *ConversationAsrUsageModel) SetAsrModel(asrModel *string) {
	c.AsrModel = asrModel
}

func (c *ConversationAsrUsageModel) SetTotalTranscriptionCalls(totalTranscriptionCalls *int) {
	c.TotalTranscriptionCalls = totalTranscriptionCalls
}

func (c *ConversationAsrUsageModel) SetTotalAudioInputSeconds(totalAudioInputSeconds *float64) {
	c.TotalAudioInputSeconds = totalAudioInputSeconds
}

func (c *ConversationAsrUsageModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationAsrUsageModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationAsrUsageModel to string"
	}
	return string(jsonData)
}

func (c *ConversationAsrUsageModel) UnmarshalJSON(data []byte) error {
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

func (c *ConversationAsrUsageModel) MarshalJSON() ([]byte, error) {
	type embed ConversationAsrUsageModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
