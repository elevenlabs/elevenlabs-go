package types

import (
	"testing"
)

func TestTranscriptEditErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "kind": "error",
  "error_type": "edit_failed",
  "message": "value"
}`
	assertRoundTrip(t, wire, new(TranscriptEditError))
}
