package types

import (
	"testing"
)

func TestDubbingLanguageResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "language_id": "value",
  "project_id": "value",
  "target_language": "value",
  "status": "queued",
  "model_id": "value",
  "voice_settings": {
    "cloning_strength": 10
  },
  "outputs": {
    "lossless_audio": "value",
    "subtitles": "value"
  },
  "revision": 123,
  "output_revision": 123,
  "error": {
    "code": "value",
    "message": "value",
    "retryable": true
  },
  "warnings": [
    {
      "type": "voices_not_permitted",
      "speaker_ids": [
        "value"
      ],
      "message": "value"
    }
  ],
  "created_at": "2024-01-01T00:00:00Z",
  "updated_at": "2024-01-01T00:00:00Z"
}`
	assertRoundTrip(t, wire, new(DubbingLanguageResponse))
}
