package types

import (
	"testing"
)

func TestPreviewAudioDbModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_id": "value",
  "text": "value",
  "audio_url": "value",
  "hls_manifest_url": "value",
  "dash_manifest_url": "value",
  "is_auto_generated": true,
  "generated_at_unix": 123
}`
	assertRoundTrip(t, wire, new(PreviewAudioDbModel))
}
