package types

import (
	"testing"
)

func TestBodyCreatesAudioNativeEnabledProjectV1AudioNativePostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "image": "value",
  "author": "value",
  "title": "value",
  "small": true,
  "text_color": "value",
  "background_color": "value",
  "sessionization": 123,
  "voice_id": "value",
  "model_id": "value",
  "auto_convert": true,
  "apply_text_normalization": "auto",
  "pronunciation_dictionary_locators": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(BodyCreatesAudioNativeEnabledProjectV1AudioNativePost))
}
