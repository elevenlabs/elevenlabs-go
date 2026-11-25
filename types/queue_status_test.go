package types

import (
	"testing"
)

func TestQueueStatusRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "queue_status_event": {
    "status": "waiting"
  },
  "type": "queue_status"
}`
	assertRoundTrip(t, wire, new(QueueStatus))
}
