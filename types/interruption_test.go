package types

import (
	"testing"
)

func TestInterruptionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "interruption_event": {
    "event_id": 123
  },
  "type": "interruption"
}`
	assertRoundTrip(t, wire, new(Interruption))
}
