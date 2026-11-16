package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SendText struct {
	Text                 string                    `json:"text" xml:"text" required:"true"`
	TryTriggerGeneration *bool                     `json:"try_trigger_generation,omitempty" xml:"try_trigger_generation,omitempty"`
	VoiceSettings        *SendTextVoiceSettings    `json:"voice_settings,omitempty" xml:"voice_settings,omitempty"`
	GenerationConfig     *SendTextGenerationConfig `json:"generation_config,omitempty" xml:"generation_config,omitempty"`
	Flush                *bool                     `json:"flush,omitempty" xml:"flush,omitempty"`
	extraProperties      map[string]interface{}
}

func (s *SendText) GetText() (value string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SendText) GetTryTriggerGeneration() (value *bool) {
	if s == nil {
		return
	}
	return s.TryTriggerGeneration
}

func (s *SendText) GetVoiceSettings() (value *SendTextVoiceSettings) {
	if s == nil {
		return
	}
	return s.VoiceSettings
}

func (s *SendText) GetGenerationConfig() (value *SendTextGenerationConfig) {
	if s == nil {
		return
	}
	return s.GenerationConfig
}

func (s *SendText) GetFlush() (value *bool) {
	if s == nil {
		return
	}
	return s.Flush
}

func (s *SendText) SetText(text string) {
	s.Text = text
}

func (s *SendText) SetTryTriggerGeneration(tryTriggerGeneration *bool) {
	s.TryTriggerGeneration = tryTriggerGeneration
}

func (s *SendText) SetVoiceSettings(voiceSettings *SendTextVoiceSettings) {
	s.VoiceSettings = voiceSettings
}

func (s *SendText) SetGenerationConfig(generationConfig *SendTextGenerationConfig) {
	s.GenerationConfig = generationConfig
}

func (s *SendText) SetFlush(flush *bool) {
	s.Flush = flush
}

func (s *SendText) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SendText) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SendText to string"
	}
	return string(jsonData)
}

func (s *SendText) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SendText
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SendText(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SendText) MarshalJSON() ([]byte, error) {
	type embed SendText
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
