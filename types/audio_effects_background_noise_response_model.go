package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioEffectsBackgroundNoiseResponseModel struct {
	// Value for `audio_effects.background_noise_id`.
	BackgroundNoiseId string `json:"background_noise_id" xml:"background_noise_id" required:"true"`
	// Display name.
	Name string `json:"name" xml:"name" required:"true"`
	// What the noise bed sounds like.
	Description     *string `json:"description" xml:"description" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AudioEffectsBackgroundNoiseResponseModel) GetBackgroundNoiseId() (value string) {
	if a == nil {
		return
	}
	return a.BackgroundNoiseId
}

func (a *AudioEffectsBackgroundNoiseResponseModel) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AudioEffectsBackgroundNoiseResponseModel) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AudioEffectsBackgroundNoiseResponseModel) SetBackgroundNoiseId(backgroundNoiseId string) {
	a.BackgroundNoiseId = backgroundNoiseId
}

func (a *AudioEffectsBackgroundNoiseResponseModel) SetName(name string) {
	a.Name = name
}

func (a *AudioEffectsBackgroundNoiseResponseModel) SetDescription(description *string) {
	a.Description = description
}

func (a *AudioEffectsBackgroundNoiseResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioEffectsBackgroundNoiseResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioEffectsBackgroundNoiseResponseModel to string"
	}
	return string(jsonData)
}

func (a *AudioEffectsBackgroundNoiseResponseModel) UnmarshalJSON(data []byte) error {
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

func (a *AudioEffectsBackgroundNoiseResponseModel) MarshalJSON() ([]byte, error) {
	type embed AudioEffectsBackgroundNoiseResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
