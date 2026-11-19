package types

import (
	"testing"
)

func TestScribeInputErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "input_error",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(ScribeInputError))
}
