package types

import (
	"testing"
)

func TestBinaryScoringRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mode": "binary"
}`
	assertRoundTrip(t, wire, new(BinaryScoring))
}
