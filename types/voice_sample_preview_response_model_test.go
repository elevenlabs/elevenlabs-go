package types

import (
	"testing"
)

func TestVoiceSamplePreviewResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "audio_base_64": "value",
  "voice_id": "value",
  "sample_id": "value",
  "media_type": "value",
  "duration_secs": 9.99
}`
	assertRoundTrip(t, wire, new(VoiceSamplePreviewResponseModel))
}
