package types

import (
	"testing"
)

func TestSendUpstreamMessageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "agent_response",
  "event_id": 123,
  "content": "value",
  "is_final": true
}`
	assertRoundTrip(t, wire, new(SendUpstreamMessage))
}
