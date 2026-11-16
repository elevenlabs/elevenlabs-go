package types

import (
	"testing"
)

func TestDubbingProjectResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "project_id": "value",
  "status": "queued",
  "reference": "value",
  "source_language": "value",
  "model_id": "value",
  "media": {
    "filename": "value",
    "duration_s": 9.99,
    "has_video": true,
    "mime_type": "value"
  },
  "language_ids": [
    "value"
  ],
  "webhook_ids": [
    "value"
  ],
  "revision": 123,
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
	assertRoundTrip(t, wire, new(DubbingProjectResponse))
}
