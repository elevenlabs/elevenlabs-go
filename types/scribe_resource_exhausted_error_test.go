package types

import (
	"testing"
)

func TestScribeResourceExhaustedErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "resource_exhausted",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(ScribeResourceExhaustedError))
}
