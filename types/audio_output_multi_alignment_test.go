package types

import (
	"testing"
)

func TestAudioOutputMultiAlignmentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chars": [
    "value"
  ],
  "char_start_times_ms": [
    123
  ],
  "char_durations_ms": [
    123
  ]
}`
	assertRoundTrip(t, wire, new(AudioOutputMultiAlignment))
}
