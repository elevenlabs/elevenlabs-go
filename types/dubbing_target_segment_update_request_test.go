package types

import (
	"testing"
)

func TestDubbingTargetSegmentUpdateRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "translation": "value"
}`
	assertRoundTrip(t, wire, new(DubbingTargetSegmentUpdateRequest))
}
