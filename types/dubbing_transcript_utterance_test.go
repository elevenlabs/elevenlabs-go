package types

import (
	"testing"
)

func TestDubbingTranscriptUtteranceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(DubbingTranscriptUtterance))
}

func TestDubbingTranscriptUtteranceRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(DubbingTranscriptUtterance))
}
