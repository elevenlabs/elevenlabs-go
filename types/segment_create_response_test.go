package types

import (
	"testing"
)

func TestSegmentCreateResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version": 123,
  "new_segment": "value"
}`
	assertRoundTrip(t, wire, new(SegmentCreateResponse))
}
