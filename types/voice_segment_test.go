package types

import (
	"testing"
)

func TestVoiceSegmentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_id": "value",
  "start_time_seconds": 9.99,
  "end_time_seconds": 9.99,
  "character_start_index": 123,
  "character_end_index": 123,
  "dialogue_input_index": 123
}`
	assertRoundTrip(t, wire, new(VoiceSegment))
}
