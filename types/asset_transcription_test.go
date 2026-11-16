package types

import (
	"testing"
)

func TestAssetTranscriptionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "processing",
  "data": {
    "language_code": "value",
    "text": "value",
    "words": [
      "value"
    ],
    "word_start_times_ms": [
      123
    ],
    "word_end_times_ms": [
      123
    ],
    "word_speaker_ids": [
      "value"
    ]
  },
  "updated_at_ms": 123
}`
	assertRoundTrip(t, wire, new(AssetTranscription))
}
