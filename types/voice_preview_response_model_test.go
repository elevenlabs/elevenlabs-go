package types

import (
	"testing"
)

func TestVoicePreviewResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "audio_base_64": "value",
  "generated_voice_id": "value",
  "media_type": "value",
  "duration_secs": 9.99,
  "language": "value"
}`
	assertRoundTrip(t, wire, new(VoicePreviewResponseModel))
}
