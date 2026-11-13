package types

import (
	"testing"
)

func TestAutoRebaseAttemptRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "outcome": "succeeded",
  "attempted_at": 123
}`
	assertRoundTrip(t, wire, new(AutoRebaseAttempt))
}
