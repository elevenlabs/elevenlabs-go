package types

import (
	"testing"
)

func TestVoiceVerificationResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "requires_verification": true,
  "is_verified": true,
  "verification_failures": [
    "value"
  ],
  "verification_attempts_count": 123,
  "language": "value",
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
  ]
}`
	assertRoundTrip(t, wire, new(VoiceVerificationResponse))
}
