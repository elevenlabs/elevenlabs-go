package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Audio effects applied to generated speech: a filter preset, perceived distance, a reverb
// environment, a background noise bed, and stereo panning. Leave every field at its default
// to get unprocessed audio.
type EffectsSpecOutput struct {
	// ID of a filter preset that colors the voice, such as `phone` or `old_radio`. List valid IDs with `GET /v1/audio-effects/catalog`.
	FilterPresetId *string `json:"filter_preset_id" xml:"filter_preset_id" required:"true"`
	// How far the speaker sounds from the microphone, from 0 (close, unchanged) to 1 (far). Higher values roll off low and high frequencies and, when an environment is set, shift the balance from the dry voice toward the reverb.
	Distance float64 `json:"distance" xml:"distance" required:"true" min:"0" max:"1"`
	// ID of a reverb environment that places the voice in a space, such as `small_room` or `hall`. List valid IDs with `GET /v1/audio-effects/catalog`. Reverb is rendered in stereo and adds a tail of up to several seconds after the speech ends.
	EnvironmentId *string `json:"environment_id" xml:"environment_id" required:"true"`
	// ID of a background noise bed looped under the voice, such as `cafe` or `keyboard`. List valid IDs with `GET /v1/audio-effects/catalog`.
	BackgroundNoiseId *string `json:"background_noise_id" xml:"background_noise_id" required:"true"`
	// How much of the voice feeds the reverb, from 0 (none) to 1 (full). Only applies when `environment_id` is set; 0 turns the reverb off.
	SendLevel float64 `json:"send_level" xml:"send_level" required:"true" min:"0" max:"1"`
	// Stereo position of the voice, from -1 (hard left) through 0 (center) to 1 (hard right). Only the voice is panned; reverb and background noise stay centered.
	Pan             float64 `json:"pan" xml:"pan" required:"true" min:"-1" max:"1"`
	extraProperties map[string]interface{}
}

func (e *EffectsSpecOutput) GetFilterPresetId() (value *string) {
	if e == nil {
		return
	}
	return e.FilterPresetId
}

func (e *EffectsSpecOutput) GetDistance() (value float64) {
	if e == nil {
		return
	}
	return e.Distance
}

func (e *EffectsSpecOutput) GetEnvironmentId() (value *string) {
	if e == nil {
		return
	}
	return e.EnvironmentId
}

func (e *EffectsSpecOutput) GetBackgroundNoiseId() (value *string) {
	if e == nil {
		return
	}
	return e.BackgroundNoiseId
}

func (e *EffectsSpecOutput) GetSendLevel() (value float64) {
	if e == nil {
		return
	}
	return e.SendLevel
}

func (e *EffectsSpecOutput) GetPan() (value float64) {
	if e == nil {
		return
	}
	return e.Pan
}

func (e *EffectsSpecOutput) SetFilterPresetId(filterPresetId *string) {
	e.FilterPresetId = filterPresetId
}

func (e *EffectsSpecOutput) SetDistance(distance float64) {
	e.Distance = distance
}

func (e *EffectsSpecOutput) SetEnvironmentId(environmentId *string) {
	e.EnvironmentId = environmentId
}

func (e *EffectsSpecOutput) SetBackgroundNoiseId(backgroundNoiseId *string) {
	e.BackgroundNoiseId = backgroundNoiseId
}

func (e *EffectsSpecOutput) SetSendLevel(sendLevel float64) {
	e.SendLevel = sendLevel
}

func (e *EffectsSpecOutput) SetPan(pan float64) {
	e.Pan = pan
}

func (e *EffectsSpecOutput) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EffectsSpecOutput) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EffectsSpecOutput to string"
	}
	return string(jsonData)
}

func (e *EffectsSpecOutput) UnmarshalJSON(data []byte) error {
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

func (e *EffectsSpecOutput) MarshalJSON() ([]byte, error) {
	type embed EffectsSpecOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
