package types

import (
	"testing"
)

func TestSpeakerAudioResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "audio_base_64": "value",
  "media_type": "value",
  "duration_secs": 9.99
}`
	assertRoundTrip(t, wire, new(SpeakerAudioResponseModel))
}
