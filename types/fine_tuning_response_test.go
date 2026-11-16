package types

import (
	"testing"
)

func TestFineTuningResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_allowed_to_fine_tune": true,
  "state": {
    "key": "not_started"
  },
  "verification_failures": [
    "value"
  ],
  "verification_attempts_count": 123,
  "manual_verification_requested": true,
  "language": "value",
  "progress": {
    "key": 9.99
  },
  "message": {
    "key": "value"
  },
  "dataset_duration_seconds": 9.99,
  "verification_attempts": [
    {
      "text": "value",
      "date_unix": 123,
      "accepted": true,
      "similarity": 9.99,
      "levenshtein_distance": 9.99,
      "recording": {
        "recording_id": "value",
        "mime_type": "value",
        "size_bytes": 123,
        "upload_date_unix": 123,
        "transcription": "value"
      }
    }
  ],
  "slice_ids": [
    "value"
  ],
  "manual_verification": {
    "extra_text": "value",
    "request_time_unix": 123,
    "files": [
      {
        "file_id": "value",
        "file_name": "value",
        "mime_type": "value",
        "size_bytes": 123,
        "upload_date_unix": 123
      }
    ]
  },
  "max_verification_attempts": 123,
  "next_max_verification_attempts_reset_unix_ms": 123,
  "finetuning_state": {}
}`
	assertRoundTrip(t, wire, new(FineTuningResponse))
}
