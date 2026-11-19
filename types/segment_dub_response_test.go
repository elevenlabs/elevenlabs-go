package types

import (
	"testing"
)

func TestSegmentDubResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version": 123
}`
	assertRoundTrip(t, wire, new(SegmentDubResponse))
}
