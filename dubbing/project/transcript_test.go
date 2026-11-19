package project

import (
	"testing"
)

func TestDubbingBulkSegmentUpdateRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "segments": {
    "key": {
      "text": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99
    }
  }
}`
	assertRoundTrip(t, wire, new(DubbingBulkSegmentUpdateRequest))
}

func TestDubbingSegmentCreateRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "speaker_id": "value",
  "start_s": 9.99,
  "end_s": 9.99
}`
	assertRoundTrip(t, wire, new(DubbingSegmentCreateRequest))
}
