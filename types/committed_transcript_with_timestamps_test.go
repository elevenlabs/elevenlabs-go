package types

import (
	"testing"
)

func TestCommittedTranscriptWithTimestampsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "committed_transcript_with_timestamps",
  "text": "value",
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
	assertRoundTrip(t, wire, new(CommittedTranscriptWithTimestamps))
}
