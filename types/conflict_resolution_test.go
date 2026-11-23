package types

import (
	"testing"
)

func TestConflictResolutionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "conflict_id": "value",
  "choice": "source",
  "custom_value": "value"
}`
	assertRoundTrip(t, wire, new(ConflictResolution))
}
