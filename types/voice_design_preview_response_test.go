package types

import (
	"testing"
)

func TestVoiceDesignPreviewResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "previews": [
    {
      "audio_base_64": "value",
      "generated_voice_id": "value",
      "media_type": "value",
      "duration_secs": 9.99,
      "language": "value"
    }
  ],
  "text": "value"
}`
	assertRoundTrip(t, wire, new(VoiceDesignPreviewResponse))
}
