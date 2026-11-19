package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationVoiceUsageModel struct {
	VoiceId            string   `json:"voice_id" xml:"voice_id" required:"true"`
	AudioOutputSeconds *float64 `json:"audio_output_seconds,omitempty" xml:"audio_output_seconds,omitempty"`
	extraProperties    map[string]interface{}
}

func (c *ConversationVoiceUsageModel) GetVoiceId() (value string) {
	if c == nil {
		return
	}
	return c.VoiceId
}

func (c *ConversationVoiceUsageModel) GetAudioOutputSeconds() (value *float64) {
	if c == nil {
		return
	}
	return c.AudioOutputSeconds
}

func (c *ConversationVoiceUsageModel) SetVoiceId(voiceId string) {
	c.VoiceId = voiceId
}

func (c *ConversationVoiceUsageModel) SetAudioOutputSeconds(audioOutputSeconds *float64) {
	c.AudioOutputSeconds = audioOutputSeconds
}

func (c *ConversationVoiceUsageModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationVoiceUsageModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationVoiceUsageModel to string"
	}
	return string(jsonData)
}

func (c *ConversationVoiceUsageModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationVoiceUsageModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationVoiceUsageModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationVoiceUsageModel) MarshalJSON() ([]byte, error) {
	type embed ConversationVoiceUsageModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
