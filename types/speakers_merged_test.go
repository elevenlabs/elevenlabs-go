package types

import (
	"testing"
)

func TestSpeakersMergedRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "speakers_merged",
  "from_speaker_id": "value",
  "to_speaker_id": "value"
}`
	assertRoundTrip(t, wire, new(SpeakersMerged))
}
