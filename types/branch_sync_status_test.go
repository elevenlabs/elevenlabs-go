package types

import (
	"testing"
)

func TestBranchSyncStatusRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "last_attempt": {
    "outcome": "succeeded",
    "attempted_at": 123
  }
}`
	assertRoundTrip(t, wire, new(BranchSyncStatus))
}
