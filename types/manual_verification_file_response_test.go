package types

import (
	"testing"
)

func TestManualVerificationFileResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "file_id": "value",
  "file_name": "value",
  "mime_type": "value",
  "size_bytes": 123,
  "upload_date_unix": 123
}`
	assertRoundTrip(t, wire, new(ManualVerificationFileResponse))
}
