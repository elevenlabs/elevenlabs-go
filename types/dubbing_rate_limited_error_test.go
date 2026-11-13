package types

import (
	"testing"
)

func TestDubbingRateLimitedErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "rate_limited",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(DubbingRateLimitedError))
}
