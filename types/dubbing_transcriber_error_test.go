package types

import (
	"testing"
)

func TestDubbingTranscriberErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "transcriber_error",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(DubbingTranscriberError))
}
