package types

import (
	"testing"
)

func TestManualVerificationResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(ManualVerificationResponse))
}
