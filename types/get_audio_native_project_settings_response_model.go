package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetAudioNativeProjectSettingsResponseModel struct {
	// Whether the project is enabled.
	Enabled bool `json:"enabled" xml:"enabled" required:"true"`
	// The ID of the latest snapshot of the project.
	SnapshotId *string `json:"snapshot_id,omitempty" xml:"snapshot_id,omitempty"`
	// The settings of the project.
	Settings        *AudioNativeProjectSettingsResponseModel `json:"settings,omitempty" xml:"settings,omitempty"`
	extraProperties map[string]interface{}
}

func (g *GetAudioNativeProjectSettingsResponseModel) GetEnabled() (value bool) {
	if g == nil {
		return
	}
	return g.Enabled
}

func (g *GetAudioNativeProjectSettingsResponseModel) GetSnapshotId() (value *string) {
	if g == nil {
		return
	}
	return g.SnapshotId
}

func (g *GetAudioNativeProjectSettingsResponseModel) GetSettings() (value *AudioNativeProjectSettingsResponseModel) {
	if g == nil {
		return
	}
	return g.Settings
}

func (g *GetAudioNativeProjectSettingsResponseModel) SetEnabled(enabled bool) {
	g.Enabled = enabled
}

func (g *GetAudioNativeProjectSettingsResponseModel) SetSnapshotId(snapshotId *string) {
	g.SnapshotId = snapshotId
}

func (g *GetAudioNativeProjectSettingsResponseModel) SetSettings(settings *AudioNativeProjectSettingsResponseModel) {
	g.Settings = settings
}

func (g *GetAudioNativeProjectSettingsResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAudioNativeProjectSettingsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAudioNativeProjectSettingsResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetAudioNativeProjectSettingsResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetAudioNativeProjectSettingsResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetAudioNativeProjectSettingsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
