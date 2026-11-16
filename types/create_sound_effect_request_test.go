package types

import (
	"testing"
)

func TestCreateSoundEffectRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "loop": true,
  "duration_seconds": 9.99,
  "prompt_influence": 9.99,
  "model_id": "eleven_text_to_sound_v2"
}`
	assertRoundTrip(t, wire, new(CreateSoundEffectRequest))
}
