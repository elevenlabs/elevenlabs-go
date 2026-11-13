package types

import (
	"testing"
)

func TestAudioEffectsBackgroundNoiseResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "background_noise_id": "value",
  "name": "value",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(AudioEffectsBackgroundNoiseResponseModel))
}
