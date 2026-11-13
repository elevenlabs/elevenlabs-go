package types

import (
	"testing"
)

func TestDiarizedTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "diarized_transcript",
  "segment_index": 123,
  "text": "value",
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
	assertRoundTrip(t, wire, new(DiarizedTranscript))
}
