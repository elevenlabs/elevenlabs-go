package types

import (
	"testing"
)

func TestBodyUpdateMetadataForASpeakerV1DubbingResourceDubbingIdSpeakerSpeakerIdPatchRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "speaker_name": "value",
  "voice_id": "value",
  "voice_stability": 9.99,
  "voice_similarity": 9.99,
  "voice_style": 9.99,
  "languages": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(BodyUpdateMetadataForASpeakerV1DubbingResourceDubbingIdSpeakerSpeakerIdPatch))
}
