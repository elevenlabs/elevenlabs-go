package types

import (
	"testing"
)

func TestAudioEffectsCatalogResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "filter_presets": [
    {
      "filter_preset_id": "value",
      "name": "value",
      "description": "value"
    }
  ],
  "environments": [
    {
      "environment_id": "value",
      "name": "value",
      "description": "value"
    }
  ],
  "background_noises": [
    {
      "background_noise_id": "value",
      "name": "value",
      "description": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(AudioEffectsCatalogResponseModel))
}
