package types

import (
	"testing"
)

func TestBackgroundSoundConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_type": "preset",
  "source_id": "office2",
  "volume": 0.5,
  "crossfade_loop": true
}`
	assertRoundTrip(t, wire, new(BackgroundSoundConfig))
}

func TestBackgroundSoundConfigRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_type": "preset",
  "source_id": "office2",
  "volume": 0.5,
  "crossfade_loop": true
}`
	assertRoundTrip(t, wire, new(BackgroundSoundConfig))
}
