package types

import (
	"testing"
)

func TestDubbingTranscriptWordRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "word_type": "value",
  "start_s": 9.99,
  "end_s": 9.99,
  "characters": [
    {
      "text": "value",
      "start_s": 9.99,
      "end_s": 9.99
    }
  ]
}`
	assertRoundTrip(t, wire, new(DubbingTranscriptWord))
}
