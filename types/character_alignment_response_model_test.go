package types

import (
	"testing"
)

func TestCharacterAlignmentResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "characters": [
    "value"
  ],
  "character_start_times_seconds": [
    9.99
  ],
  "character_end_times_seconds": [
    9.99
  ]
}`
	assertRoundTrip(t, wire, new(CharacterAlignmentResponseModel))
}

func TestCharacterAlignmentResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "characters": [
    "value"
  ],
  "character_start_times_seconds": [
    9.99
  ],
  "character_end_times_seconds": [
    9.99
  ]
}`
	assertRoundTrip(t, wire, new(CharacterAlignmentResponseModel))
}
