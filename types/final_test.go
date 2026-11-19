package types

import (
	"testing"
)

func TestFinalRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_final": true
}`
	assertRoundTrip(t, wire, new(Final))
}
