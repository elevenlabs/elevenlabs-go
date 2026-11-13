package types

import (
	"testing"
)

func TestOpenOutcomeRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "open"
}`
	assertRoundTrip(t, wire, new(OpenOutcome))
}
