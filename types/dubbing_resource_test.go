package types

import (
	"testing"
)

func TestDubbingResourceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "version": 123,
  "source_language": "value",
  "target_languages": [
    "value"
  ],
  "input": {
    "src": "value",
    "content_type": "value",
    "bucket_name": "value",
    "random_path_slug": "value",
    "duration_secs": 9.99,
    "is_audio": true,
    "url": "value"
  },
  "background": {
    "src": "value",
    "content_type": "value",
    "bucket_name": "value",
    "random_path_slug": "value",
    "duration_secs": 9.99,
    "is_audio": true,
    "url": "value"
  },
  "foreground": {
    "src": "value",
    "content_type": "value",
    "bucket_name": "value",
    "random_path_slug": "value",
    "duration_secs": 9.99,
    "is_audio": true,
    "url": "value"
  },
  "speaker_tracks": {
    "key": {
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
    }
  },
  "speaker_segments": {
    "key": {
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
    }
  },
  "renders": {
    "key": {
      "id": "value",
      "version": 123,
      "language": "value",
      "type": "mp4",
      "media_ref": {
        "src": "value",
        "content_type": "value",
        "bucket_name": "value",
        "random_path_slug": "value",
        "duration_secs": 9.99,
        "is_audio": true,
        "url": "value"
      },
      "status": "complete"
    }
  }
}`
	assertRoundTrip(t, wire, new(DubbingResource))
}
