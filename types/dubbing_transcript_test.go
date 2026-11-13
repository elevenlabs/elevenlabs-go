package types

import (
	"testing"
)

func TestDubbingTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "language": "value",
  "utterances": [
    {
      "text": "value",
      "speaker_id": "value",
      "start_s": 9.99,
      "end_s": 9.99,
      "words": [
        {
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
        }
      ]
    }
  ]
}`
	assertRoundTrip(t, wire, new(DubbingTranscript))
}
