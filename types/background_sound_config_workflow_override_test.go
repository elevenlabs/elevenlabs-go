package types

import (
	"testing"
)

func TestBackgroundSoundConfigWorkflowOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_type": "preset",
  "source_id": "office2",
  "volume": 9.99,
  "crossfade_loop": true
}`
	assertRoundTrip(t, wire, new(BackgroundSoundConfigWorkflowOverride))
}
