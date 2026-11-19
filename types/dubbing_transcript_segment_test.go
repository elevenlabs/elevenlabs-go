package types

import (
	"testing"
)

func TestDubbingTranscriptSegmentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "text": "value",
  "speaker_id": "value",
  "start_s": 9.99,
  "end_s": 9.99,
  "external_id": "value"
}`
	assertRoundTrip(t, wire, new(DubbingTranscriptSegment))
}
