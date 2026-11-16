package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Context for sound effect clips.
type SfxSourceContext struct {
	SoundGenerationHistoryItemId *string        `json:"sound_generation_history_item_id,omitempty" xml:"sound_generation_history_item_id,omitempty"`
	Text                         *string        `json:"text,omitempty" xml:"text,omitempty"`
	GenerationConfig             map[string]any `json:"generation_config,omitempty"`
	extraProperties              map[string]interface{}
}

func (s *SfxSourceContext) GetSoundGenerationHistoryItemId() (value *string) {
	if s == nil {
		return
	}
	return s.SoundGenerationHistoryItemId
}

func (s *SfxSourceContext) GetText() (value *string) {
	if s == nil {
		return
	}
	return s.Text
}

func (s *SfxSourceContext) GetGenerationConfig() (value map[string]any) {
	if s == nil {
		return
	}
	return s.GenerationConfig
}

func (s *SfxSourceContext) SetSoundGenerationHistoryItemId(soundGenerationHistoryItemId *string) {
	s.SoundGenerationHistoryItemId = soundGenerationHistoryItemId
}

func (s *SfxSourceContext) SetText(text *string) {
	s.Text = text
}

func (s *SfxSourceContext) SetGenerationConfig(generationConfig map[string]any) {
	s.GenerationConfig = generationConfig
}

func (s *SfxSourceContext) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SfxSourceContext) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SfxSourceContext to string"
	}
	return string(jsonData)
}

func (s *SfxSourceContext) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, s); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "source_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SfxSourceContext) MarshalJSON() ([]byte, error) {
	type embed SfxSourceContext
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
