package types

import (
	"testing"
)

func TestAssetTranscriptionDataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(AssetTranscriptionData))
}
