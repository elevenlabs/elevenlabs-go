package types

import (
	"testing"
)

func TestAudioEffectsEnvironmentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "environment_id": "value",
  "name": "value",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(AudioEffectsEnvironmentResponseModel))
}
