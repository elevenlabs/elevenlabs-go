package types

import (
	"testing"
)

func TestSpeakerSegmentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "start_time": 9.99,
  "end_time": 9.99,
  "text": "value",
  "subtitles": [
    {
      "start_time": 9.99,
      "end_time": 9.99,
      "lines": [
        "value"
      ]
    }
  ],
  "dubs": {
    "key": {
      "start_time": 9.99,
      "end_time": 9.99,
      "text": "value",
      "subtitles": [
        {
          "start_time": 9.99,
          "end_time": 9.99,
          "lines": [
            "value"
          ]
        }
      ],
      "audio_stale": true,
      "media_ref": {
        "src": "value",
        "content_type": "value",
        "bucket_name": "value",
        "random_path_slug": "value",
        "duration_secs": 9.99,
        "is_audio": true,
        "url": "value"
      }
    }
  }
}`
	assertRoundTrip(t, wire, new(SpeakerSegment))
}
