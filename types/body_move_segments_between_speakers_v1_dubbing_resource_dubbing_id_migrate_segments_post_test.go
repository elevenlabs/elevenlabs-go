package types

import (
	"testing"
)

func TestBodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "segment_ids": [
    "value"
  ],
  "speaker_id": "value"
}`
	assertRoundTrip(t, wire, new(BodyMoveSegmentsBetweenSpeakersV1DubbingResourceDubbingIdMigrateSegmentsPost))
}
