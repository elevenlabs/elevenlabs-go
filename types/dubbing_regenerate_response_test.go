package types

import (
	"testing"
)

func TestDubbingRegenerateResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "regenerated_segment_ids": [
    "value"
  ],
  "regenerated_seconds": 9.99,
  "charged_seconds": 9.99,
  "free_regeneration_seconds_remaining": 9.99
}`
	assertRoundTrip(t, wire, new(DubbingRegenerateResponse))
}
