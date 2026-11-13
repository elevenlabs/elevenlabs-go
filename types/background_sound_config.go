package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BackgroundSoundConfig struct {
	// The type of background sound source.
	SourceType *BackgroundSoundSourceType `json:"source_type,omitempty" xml:"source_type,omitempty"`
	// Identifier for the sound source.
	SourceId *BackgroundSoundPresetId `json:"source_id,omitempty" xml:"source_id,omitempty"`
	// Volume level for background sound (0.01 to 1.0).
	Volume *float64 `json:"volume,omitempty" xml:"volume,omitempty" min:"0.01" max:"1"`
	// Apply a crossfade at the loop boundary to avoid audible pops when the sound loops.
	CrossfadeLoop   *bool `json:"crossfade_loop,omitempty" xml:"crossfade_loop,omitempty"`
	extraProperties map[string]interface{}
}

func (b *BackgroundSoundConfig) GetSourceType() (value *BackgroundSoundSourceType) {
	if b == nil {
		return
	}
	return b.SourceType
}

func (b *BackgroundSoundConfig) GetSourceId() (value *BackgroundSoundPresetId) {
	if b == nil {
		return
	}
	return b.SourceId
}

func (b *BackgroundSoundConfig) GetVolume() (value *float64) {
	if b == nil {
		return
	}
	return b.Volume
}

func (b *BackgroundSoundConfig) GetCrossfadeLoop() (value *bool) {
	if b == nil {
		return
	}
	return b.CrossfadeLoop
}

func (b *BackgroundSoundConfig) SetSourceType(sourceType *BackgroundSoundSourceType) {
	b.SourceType = sourceType
}

func (b *BackgroundSoundConfig) SetSourceId(sourceId *BackgroundSoundPresetId) {
	b.SourceId = sourceId
}

func (b *BackgroundSoundConfig) SetVolume(volume *float64) {
	b.Volume = volume
}

func (b *BackgroundSoundConfig) SetCrossfadeLoop(crossfadeLoop *bool) {
	b.CrossfadeLoop = crossfadeLoop
}

func (b *BackgroundSoundConfig) GetExtraProperties() map[string]interface{} {
	if b == nil {
		return nil
	}
	return b.extraProperties
}

func (b BackgroundSoundConfig) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BackgroundSoundConfig to string"
	}
	return string(jsonData)
}

func (b *BackgroundSoundConfig) UnmarshalJSON(data []byte) error {
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

func (b *BackgroundSoundConfig) MarshalJSON() ([]byte, error) {
	type embed BackgroundSoundConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
