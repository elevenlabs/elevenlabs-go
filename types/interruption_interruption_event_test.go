package types

import (
	"testing"
)

func TestInterruptionInterruptionEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "event_id": 123
}`
	assertRoundTrip(t, wire, new(InterruptionInterruptionEvent))
}
