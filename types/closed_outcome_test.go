package types

import (
	"testing"
)

func TestClosedOutcomeRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "closed",
  "reason": "withdrawn",
  "at": 123,
  "by_user_id": "value"
}`
	assertRoundTrip(t, wire, new(ClosedOutcome))
}
