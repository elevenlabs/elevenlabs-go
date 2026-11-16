package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationTurnMetrics struct {
	Metrics           map[string]*MetricRecord `json:"metrics,omitempty"`
	ConvaiAsrProvider *string                  `json:"convai_asr_provider,omitempty" xml:"convai_asr_provider,omitempty"`
	ConvaiTtsModel    *string                  `json:"convai_tts_model,omitempty" xml:"convai_tts_model,omitempty"`
	ConvaiTtsCascade  *string                  `json:"convai_tts_cascade,omitempty" xml:"convai_tts_cascade,omitempty"`
	VoiceId           *string                  `json:"voice_id,omitempty" xml:"voice_id,omitempty"`
	extraProperties   map[string]interface{}
}

func (c *ConversationTurnMetrics) GetMetrics() (value map[string]*MetricRecord) {
	if c == nil {
		return
	}
	return c.Metrics
}

func (c *ConversationTurnMetrics) GetConvaiAsrProvider() (value *string) {
	if c == nil {
		return
	}
	return c.ConvaiAsrProvider
}

func (c *ConversationTurnMetrics) GetConvaiTtsModel() (value *string) {
	if c == nil {
		return
	}
	return c.ConvaiTtsModel
}

func (c *ConversationTurnMetrics) GetConvaiTtsCascade() (value *string) {
	if c == nil {
		return
	}
	return c.ConvaiTtsCascade
}

func (c *ConversationTurnMetrics) GetVoiceId() (value *string) {
	if c == nil {
		return
	}
	return c.VoiceId
}

func (c *ConversationTurnMetrics) SetMetrics(metrics map[string]*MetricRecord) {
	c.Metrics = metrics
}

func (c *ConversationTurnMetrics) SetConvaiAsrProvider(convaiAsrProvider *string) {
	c.ConvaiAsrProvider = convaiAsrProvider
}

func (c *ConversationTurnMetrics) SetConvaiTtsModel(convaiTtsModel *string) {
	c.ConvaiTtsModel = convaiTtsModel
}

func (c *ConversationTurnMetrics) SetConvaiTtsCascade(convaiTtsCascade *string) {
	c.ConvaiTtsCascade = convaiTtsCascade
}

func (c *ConversationTurnMetrics) SetVoiceId(voiceId *string) {
	c.VoiceId = voiceId
}

func (c *ConversationTurnMetrics) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationTurnMetrics) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationTurnMetrics to string"
	}
	return string(jsonData)
}

func (c *ConversationTurnMetrics) UnmarshalJSON(data []byte) error {
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

func (c *ConversationTurnMetrics) MarshalJSON() ([]byte, error) {
	type embed ConversationTurnMetrics
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
