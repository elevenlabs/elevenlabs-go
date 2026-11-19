package types

import (
	"testing"
)

func TestPartialTranscriptWordsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(PartialTranscriptWordsItem))
}
