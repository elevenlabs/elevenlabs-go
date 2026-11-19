package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AudioEffectsEnvironmentResponseModel struct {
	// Value for `audio_effects.environment_id`.
	EnvironmentId string `json:"environment_id" xml:"environment_id" required:"true"`
	// Display name.
	Name string `json:"name" xml:"name" required:"true"`
	// What the space sounds like.
	Description     *string `json:"description" xml:"description" required:"true"`
	extraProperties map[string]interface{}
}

func (a *AudioEffectsEnvironmentResponseModel) GetEnvironmentId() (value string) {
	if a == nil {
		return
	}
	return a.EnvironmentId
}

func (a *AudioEffectsEnvironmentResponseModel) GetName() (value string) {
	if a == nil {
		return
	}
	return a.Name
}

func (a *AudioEffectsEnvironmentResponseModel) GetDescription() (value *string) {
	if a == nil {
		return
	}
	return a.Description
}

func (a *AudioEffectsEnvironmentResponseModel) SetEnvironmentId(environmentId string) {
	a.EnvironmentId = environmentId
}

func (a *AudioEffectsEnvironmentResponseModel) SetName(name string) {
	a.Name = name
}

func (a *AudioEffectsEnvironmentResponseModel) SetDescription(description *string) {
	a.Description = description
}

func (a *AudioEffectsEnvironmentResponseModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AudioEffectsEnvironmentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AudioEffectsEnvironmentResponseModel to string"
	}
	return string(jsonData)
}

func (a *AudioEffectsEnvironmentResponseModel) UnmarshalJSON(data []byte) error {
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

func (a *AudioEffectsEnvironmentResponseModel) MarshalJSON() ([]byte, error) {
	type embed AudioEffectsEnvironmentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
