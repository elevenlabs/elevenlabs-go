package types

import (
	"testing"
)

func TestPartialTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "partial_transcript",
  "text": "value",
  "turn_probability": 9.99,
  "language_code": "value",
  "words": [
    {
      "text": "value",
      "start": 9.99,
      "end": 9.99,
      "type": "word",
      "speaker_id": "value",
      "logprob": 9.99,
      "characters": [
        {
          "text": "value",
          "start": 9.99,
          "end": 9.99
        }
      ],
      "channel_index": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(PartialTranscript))
}
