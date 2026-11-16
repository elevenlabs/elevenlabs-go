package types

import (
	"testing"
)

func TestManualSourceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "manual",
  "created_by_user_id": "value",
  "notes": "value"
}`
	assertRoundTrip(t, wire, new(ManualSource))
}
