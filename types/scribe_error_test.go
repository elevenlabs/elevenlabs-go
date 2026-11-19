package types

import (
	"testing"
)

func TestScribeErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "error",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(ScribeError))
}
