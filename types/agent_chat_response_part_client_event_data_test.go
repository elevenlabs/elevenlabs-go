package types

import (
	"testing"
)

func TestAgentChatResponsePartClientEventDataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "text": "value",
  "type": "start",
  "event_id": 123,
  "response_id": "value"
}`
	assertRoundTrip(t, wire, new(AgentChatResponsePartClientEventData))
}
