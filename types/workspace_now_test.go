package types

import (
	"testing"
)

func TestWorkspaceNowRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "today": {
    "year": 2026,
    "month": 12,
    "day": 31
  },
  "hour": 123,
  "minute": 123,
  "fetched_at_ms": 123
}`
	assertRoundTrip(t, wire, new(WorkspaceNow))
}
