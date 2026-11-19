package types

import (
	"testing"
)

func TestSpeakerResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "speaker_id": "value",
  "duration_secs": 9.99,
  "utterances": [
    {
      "start": 9.99,
      "end": 9.99
    }
  ]
}`
	assertRoundTrip(t, wire, new(SpeakerResponseModel))
}
