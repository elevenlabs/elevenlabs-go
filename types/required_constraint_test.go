package types

import (
	"testing"
)

func TestRequiredConstraintRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "required": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(RequiredConstraint))
}
