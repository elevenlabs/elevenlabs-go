package types

import (
	"testing"
)

func TestProjectVoiceResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "project_voice_ref_id": "value",
  "voice_id": "value",
  "alias": "value",
  "stability": 9.99,
  "similarity_boost": 9.99,
  "style": 9.99,
  "is_pinned": true,
  "use_speaker_boost": true,
  "volume_gain": 9.99,
  "speed": 9.99
}`
	assertRoundTrip(t, wire, new(ProjectVoiceResponseModel))
}
