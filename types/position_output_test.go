package types

import (
	"testing"
)

func TestPositionOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "x": 9.99,
  "y": 9.99
}`
	assertRoundTrip(t, wire, new(PositionOutput))
}
