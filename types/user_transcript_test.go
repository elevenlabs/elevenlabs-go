package types

import (
	"testing"
)

func TestUserTranscriptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "user_transcript",
  "event_id": 123,
  "user_transcript": [
    {
      "role": "value",
      "content": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(UserTranscript))
}
