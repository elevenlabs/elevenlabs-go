package types

import (
	"testing"
)

func TestCueOptionsRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "min_duration_ms": 123,
  "max_duration_ms": 4000,
  "max_lines_per_cue": 3,
  "max_chars_per_line": 50,
  "max_chars_per_s": 30,
  "min_gap_between_cues_frames": 6
}`
	assertRoundTrip(t, wire, new(CueOptionsRequest))
}
