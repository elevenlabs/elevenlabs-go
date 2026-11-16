package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioEffectsFilterPresetResponseModel struct {
	// Value for `audio_effects.filter_preset_id`.
	FilterPresetId string `json:"filter_preset_id" xml:"filter_preset_id" required:"true"`
	// Display name.
	Name string `json:"name" xml:"name" required:"true"`
	// What the preset sounds like.
	Description     *string `json:"description" xml:"description" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AudioEffectsFilterPresetResponseModel) GetFilterPresetId() (value string) {
	if a == nil {
		return
	}
	return a.FilterPresetId
}

func (a *AudioEffectsFilterPresetResponseModel) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AudioEffectsFilterPresetResponseModel) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AudioEffectsFilterPresetResponseModel) SetFilterPresetId(filterPresetId string) {
	a.FilterPresetId = filterPresetId
}

func (a *AudioEffectsFilterPresetResponseModel) SetName(name string) {
	a.Name = name
}

func (a *AudioEffectsFilterPresetResponseModel) SetDescription(description *string) {
	a.Description = description
}

func (a *AudioEffectsFilterPresetResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioEffectsFilterPresetResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioEffectsFilterPresetResponseModel to string"
	}
	return string(jsonData)
}

func (a *AudioEffectsFilterPresetResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AudioEffectsFilterPresetResponseModel) MarshalJSON() ([]byte, error) {
	type embed AudioEffectsFilterPresetResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
