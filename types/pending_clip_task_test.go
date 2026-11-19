package types

import (
	"testing"
)

func TestPendingClipTaskRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "preprocessing",
  "progress": 9.99,
  "started_at_ms": 123,
  "updated_at_ms": 123,
  "metadata": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(PendingClipTask))
}
