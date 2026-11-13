package types

import (
	"testing"
)

func TestSegmentDeleteResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version": 123
}`
	assertRoundTrip(t, wire, new(SegmentDeleteResponse))
}
