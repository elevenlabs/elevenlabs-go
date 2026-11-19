package types

import (
	"testing"
)

func TestDubbingStatusRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "status",
  "status": "started"
}`
	assertRoundTrip(t, wire, new(DubbingStatus))
}
