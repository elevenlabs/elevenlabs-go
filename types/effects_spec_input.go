package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Audio effects applied to generated speech: a filter preset, perceived distance, a reverb
// environment, a background noise bed, and stereo panning. Leave every field at its default
// to get unprocessed audio.
type EffectsSpecInput struct {
	// ID of a filter preset that colors the voice, such as `phone` or `old_radio`. List valid IDs with `GET /v1/audio-effects/catalog`.
	FilterPresetId *string `json:"filter_preset_id,omitempty" xml:"filter_preset_id,omitempty"`
	// How far the speaker sounds from the microphone, from 0 (close, unchanged) to 1 (far). Higher values roll off low and high frequencies and, when an environment is set, shift the balance from the dry voice toward the reverb.
	Distance *float64 `json:"distance,omitempty" xml:"distance,omitempty" min:"0" max:"1"`
	// ID of a reverb environment that places the voice in a space, such as `small_room` or `hall`. List valid IDs with `GET /v1/audio-effects/catalog`. Reverb is rendered in stereo and adds a tail of up to several seconds after the speech ends.
	EnvironmentId *string `json:"environment_id,omitempty" xml:"environment_id,omitempty"`
	// ID of a background noise bed looped under the voice, such as `cafe` or `keyboard`. List valid IDs with `GET /v1/audio-effects/catalog`.
	BackgroundNoiseId *string `json:"background_noise_id,omitempty" xml:"background_noise_id,omitempty"`
	// How much of the voice feeds the reverb, from 0 (none) to 1 (full). Only applies when `environment_id` is set; 0 turns the reverb off.
	SendLevel *float64 `json:"send_level,omitempty" xml:"send_level,omitempty" min:"0" max:"1"`
	// Stereo position of the voice, from -1 (hard left) through 0 (center) to 1 (hard right). Only the voice is panned; reverb and background noise stay centered.
	Pan             *float64 `json:"pan,omitempty" xml:"pan,omitempty" min:"-1" max:"1"`
	extraProperties map[string]interface{}
}

func (e *EffectsSpecInput) GetFilterPresetId() (value *string) {
	if e == nil {
		return
	}
	return e.FilterPresetId
}

func (e *EffectsSpecInput) GetDistance() (value *float64) {
	if e == nil {
		return
	}
	return e.Distance
}

func (e *EffectsSpecInput) GetEnvironmentId() (value *string) {
	if e == nil {
		return
	}
	return e.EnvironmentId
}

func (e *EffectsSpecInput) GetBackgroundNoiseId() (value *string) {
	if e == nil {
		return
	}
	return e.BackgroundNoiseId
}

func (e *EffectsSpecInput) GetSendLevel() (value *float64) {
	if e == nil {
		return
	}
	return e.SendLevel
}

func (e *EffectsSpecInput) GetPan() (value *float64) {
	if e == nil {
		return
	}
	return e.Pan
}

func (e *EffectsSpecInput) SetFilterPresetId(filterPresetId *string) {
	e.FilterPresetId = filterPresetId
}

func (e *EffectsSpecInput) SetDistance(distance *float64) {
	e.Distance = distance
}

func (e *EffectsSpecInput) SetEnvironmentId(environmentId *string) {
	e.EnvironmentId = environmentId
}

func (e *EffectsSpecInput) SetBackgroundNoiseId(backgroundNoiseId *string) {
	e.BackgroundNoiseId = backgroundNoiseId
}

func (e *EffectsSpecInput) SetSendLevel(sendLevel *float64) {
	e.SendLevel = sendLevel
}

func (e *EffectsSpecInput) SetPan(pan *float64) {
	e.Pan = pan
}

func (e *EffectsSpecInput) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EffectsSpecInput) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EffectsSpecInput to string"
	}
	return string(jsonData)
}

func (e *EffectsSpecInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, e); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *e)
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EffectsSpecInput) MarshalJSON() ([]byte, error) {
	type embed EffectsSpecInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
