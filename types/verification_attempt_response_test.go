package types

import (
	"testing"
)

func TestVerificationAttemptResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(VerificationAttemptResponse))
}
