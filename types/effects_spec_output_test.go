package types

import (
	"testing"
)

func TestEffectsSpecOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "filter_preset_id": "value",
  "distance": 0.5,
  "environment_id": "value",
  "background_noise_id": "value",
  "send_level": 0.5,
  "pan": 0.5
}`
	assertRoundTrip(t, wire, new(EffectsSpecOutput))
}

func TestEffectsSpecOutputRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "filter_preset_id": "value",
  "distance": 0.5,
  "environment_id": "value",
  "background_noise_id": "value",
  "send_level": 0.5,
  "pan": 0.5
}`
	assertRoundTrip(t, wire, new(EffectsSpecOutput))
}
