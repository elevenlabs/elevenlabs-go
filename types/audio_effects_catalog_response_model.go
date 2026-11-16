package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type AudioEffectsCatalogResponseModel struct {
	FilterPresets    []*AudioEffectsFilterPresetResponseModel    `json:"filter_presets" xml:"filter_presets" required:"true"`
	Environments     []*AudioEffectsEnvironmentResponseModel     `json:"environments" xml:"environments" required:"true"`
	BackgroundNoises []*AudioEffectsBackgroundNoiseResponseModel `json:"background_noises" xml:"background_noises" required:"true"`
	extraProperties  map[string]interface{}
}

func (a *AudioEffectsCatalogResponseModel) GetFilterPresets() (value []*AudioEffectsFilterPresetResponseModel) {
	if a == nil {
		return
	}
	return a.FilterPresets
}

func (a *AudioEffectsCatalogResponseModel) GetEnvironments() (value []*AudioEffectsEnvironmentResponseModel) {
	if a == nil {
		return
	}
	return a.Environments
}

func (a *AudioEffectsCatalogResponseModel) GetBackgroundNoises() (value []*AudioEffectsBackgroundNoiseResponseModel) {
	if a == nil {
		return
	}
	return a.BackgroundNoises
}

func (a *AudioEffectsCatalogResponseModel) SetFilterPresets(filterPresets []*AudioEffectsFilterPresetResponseModel) {
	a.FilterPresets = filterPresets
}

func (a *AudioEffectsCatalogResponseModel) SetEnvironments(environments []*AudioEffectsEnvironmentResponseModel) {
	a.Environments = environments
}

func (a *AudioEffectsCatalogResponseModel) SetBackgroundNoises(backgroundNoises []*AudioEffectsBackgroundNoiseResponseModel) {
	a.BackgroundNoises = backgroundNoises
}

func (a *AudioEffectsCatalogResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioEffectsCatalogResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioEffectsCatalogResponseModel to string"
	}
	return string(jsonData)
}

func (a *AudioEffectsCatalogResponseModel) UnmarshalJSON(data []byte) error {
	type unmarshaler AudioEffectsCatalogResponseModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*a = AudioEffectsCatalogResponseModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioEffectsCatalogResponseModel) MarshalJSON() ([]byte, error) {
	type embed AudioEffectsCatalogResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
