package types

import (
	"testing"
)

func TestProjectExternalAudioResponseModelSourceContextRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source_type": "song",
  "song_id": "value",
  "chat_id": "value",
  "title": "value",
  "description": "value",
  "genres": [
    "value"
  ],
  "languages": [
    "value"
  ],
  "is_explicit": true,
  "bpm": 123,
  "generation_settings": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(ProjectExternalAudioResponseModelSourceContext))
}
