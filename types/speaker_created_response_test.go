package types

import (
	"testing"
)

func TestSpeakerCreatedResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version": 123,
  "speaker_id": "value"
}`
	assertRoundTrip(t, wire, new(SpeakerCreatedResponse))
}
