package types

import (
	"testing"
)

func TestMergedOutcomeRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "merged",
  "version_id": "value",
  "at": 123,
  "by_user_id": "value"
}`
	assertRoundTrip(t, wire, new(MergedOutcome))
}
