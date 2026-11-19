package types

import (
	"testing"
)

func TestDubbingSourceSegmentUpdateResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "segment": {
    "id": "value",
    "text": "value",
    "speaker_id": "value",
    "start_s": 9.99,
    "end_s": 9.99,
    "external_id": "value"
  },
  "revision": 123
}`
	assertRoundTrip(t, wire, new(DubbingSourceSegmentUpdateResponse))
}
