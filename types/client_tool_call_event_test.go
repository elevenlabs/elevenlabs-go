package types

import (
	"testing"
)

func TestClientToolCallEventRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "client_tool_call",
  "client_tool_call": {
    "tool_name": "value",
    "tool_call_id": "value",
    "parameters": {
      "key": {}
    },
    "event_id": 123,
    "expects_response": true
  }
}`
	assertRoundTrip(t, wire, new(ClientToolCallEvent))
}
