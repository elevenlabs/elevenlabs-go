package types

import (
	"testing"
)

func TestScribeSessionTimeLimitExceededErrorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "session_time_limit_exceeded",
  "error": "value"
}`
	assertRoundTrip(t, wire, new(ScribeSessionTimeLimitExceededError))
}
