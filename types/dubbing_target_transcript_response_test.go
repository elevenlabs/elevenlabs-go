package types

import (
	"testing"
)

func TestDubbingTargetTranscriptResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_language": "value",
  "target_language": "value",
  "segments": [
    {
      "id": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "source_text": "value",
      "translation": "value"
    }
  ],
  "revision": 123
}`
	assertRoundTrip(t, wire, new(DubbingTargetTranscriptResponse))
}
