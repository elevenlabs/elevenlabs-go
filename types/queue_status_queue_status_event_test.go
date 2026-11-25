package types

import (
	"testing"
)

func TestQueueStatusQueueStatusEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "waiting"
}`
	assertRoundTrip(t, wire, new(QueueStatusQueueStatusEvent))
}
