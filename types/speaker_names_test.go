package types

import (
	"testing"
)

func TestSpeakerNamesRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "speaker_names",
  "speaker_names": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(SpeakerNames))
}
