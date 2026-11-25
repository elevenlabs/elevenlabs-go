package types

import (
	"testing"
)

func TestSubtitlingInputErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "input_error",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(SubtitlingInputError))
}
