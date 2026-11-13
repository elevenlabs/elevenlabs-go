package types

import (
	"testing"
)

func TestSpeakerTrackRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "media_ref": {
    "src": "value",
    "content_type": "value",
    "bucket_name": "value",
    "random_path_slug": "value",
    "duration_secs": 9.99,
    "is_audio": true,
    "url": "value"
  },
  "speaker_name": "value",
  "voices": {
    "key": "value"
  },
  "segments": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(SpeakerTrack))
}
