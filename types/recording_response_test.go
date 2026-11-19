package types

import (
	"testing"
)

func TestRecordingResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "recording_id": "value",
  "mime_type": "value",
  "size_bytes": 123,
  "upload_date_unix": 123,
  "transcription": "value"
}`
	assertRoundTrip(t, wire, new(RecordingResponse))
}
