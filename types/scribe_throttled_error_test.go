package types

import (
	"testing"
)

func TestScribeThrottledErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "commit_throttled",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(ScribeThrottledError))
}
