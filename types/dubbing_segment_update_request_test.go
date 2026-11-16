package types

import (
	"testing"
)

func TestDubbingSegmentUpdateRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "speaker_id": "value",
  "start_s": 9.99,
  "end_s": 9.99
}`
	assertRoundTrip(t, wire, new(DubbingSegmentUpdateRequest))
}
