package types

import (
	"testing"
)

func TestAudioNativeProjectSettingsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "image": "value",
  "author": "value",
  "small": true,
  "text_color": "value",
  "background_color": "value",
  "sessionization": 123,
  "audio_path": "value",
  "audio_url": "value",
  "status": "processing"
}`
	assertRoundTrip(t, wire, new(AudioNativeProjectSettingsResponseModel))
}
