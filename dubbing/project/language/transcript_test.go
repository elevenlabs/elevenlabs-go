package language

import (
	"testing"
)

func TestDubbingBulkTargetSegmentUpdateRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "segments": {
    "key": {
      "translation": "value"
    }
  }
}`
	assertRoundTrip(t, wire, new(DubbingBulkTargetSegmentUpdateRequest))
}
