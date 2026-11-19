package types

import (
	"testing"
)

func TestBodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "speaker_name": "value",
  "voice_id": "value",
  "voice_stability": 9.99,
  "voice_similarity": 9.99,
  "voice_style": 9.99
}`
	assertRoundTrip(t, wire, new(BodyCreateANewSpeakerV1DubbingResourceDubbingIdSpeakerPost))
}
