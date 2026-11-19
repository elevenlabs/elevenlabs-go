package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BackgroundSoundConfigWorkflowOverride struct {
	// The type of background sound source.
	SourceType *BackgroundSoundSourceType `json:"source_type,omitempty" xml:"source_type,omitempty"`
	// Identifier for the sound source.
	SourceId *BackgroundSoundPresetId `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// Volume level for background sound (0.01 to 1.0).
	Volume *float64 `json:"volume,omitempty" xml:"volume,omitempty"`
	// Apply a crossfade at the loop boundary to avoid audible pops when the sound loops.
	CrossfadeLoop   *bool `json:"crossfade_loop,omitempty" xml:"crossfade_loop,omitempty"`
	extraProperties map[string]interface{}
}

func (b *BackgroundSoundConfigWorkflowOverride) GetSourceType() (value *BackgroundSoundSourceType) {
	if b == nil {
		return
	}
	return b.SourceType
}

func (b *BackgroundSoundConfigWorkflowOverride) GetSourceId() (value *BackgroundSoundPresetId) {
	if b == nil {
		return
	}
	return b.SourceId
}

func (b *BackgroundSoundConfigWorkflowOverride) GetVolume() (value *float64) {
	if b == nil {
		return
	}
	return b.Volume
}

func (b *BackgroundSoundConfigWorkflowOverride) GetCrossfadeLoop() (value *bool) {
	if b == nil {
		return
	}
	return b.CrossfadeLoop
}

func (b *BackgroundSoundConfigWorkflowOverride) SetSourceType(sourceType *BackgroundSoundSourceType) {
	b.SourceType = sourceType
}

func (b *BackgroundSoundConfigWorkflowOverride) SetSourceId(sourceId *BackgroundSoundPresetId) {
	b.SourceId = sourceId
}

func (b *BackgroundSoundConfigWorkflowOverride) SetVolume(volume *float64) {
	b.Volume = volume
}

func (b *BackgroundSoundConfigWorkflowOverride) SetCrossfadeLoop(crossfadeLoop *bool) {
	b.CrossfadeLoop = crossfadeLoop
}

func (b *BackgroundSoundConfigWorkflowOverride) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BackgroundSoundConfigWorkflowOverride) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BackgroundSoundConfigWorkflowOverride to string"
	}
	return string(jsonData)
}

func (b *BackgroundSoundConfigWorkflowOverride) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, b); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	return nil
}

func (b *BackgroundSoundConfigWorkflowOverride) MarshalJSON() ([]byte, error) {
	type embed BackgroundSoundConfigWorkflowOverride
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
