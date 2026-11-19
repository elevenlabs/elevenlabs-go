package types

import (
	"testing"
)

func TestDubbingTargetTranscriptSegmentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "speaker_id": "value",
  "start_s": 9.99,
  "end_s": 9.99,
  "source_text": "value",
  "translation": "value"
}`
	assertRoundTrip(t, wire, new(DubbingTargetTranscriptSegment))
}
