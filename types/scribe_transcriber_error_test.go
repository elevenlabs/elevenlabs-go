package types

import (
	"testing"
)

func TestScribeTranscriberErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "transcriber_error",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(ScribeTranscriberError))
}
