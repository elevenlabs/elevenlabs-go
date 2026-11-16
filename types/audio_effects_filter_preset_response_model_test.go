package types

import (
	"testing"
)

func TestAudioEffectsFilterPresetResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "filter_preset_id": "value",
  "name": "value",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(AudioEffectsFilterPresetResponseModel))
}
