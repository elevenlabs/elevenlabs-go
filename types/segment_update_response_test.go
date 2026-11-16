package types

import (
	"testing"
)

func TestSegmentUpdateResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version": 123
}`
	assertRoundTrip(t, wire, new(SegmentUpdateResponse))
}
