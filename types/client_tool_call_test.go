package types

import (
	"testing"
)

func TestClientToolCallRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "client_tool_call": {
    "tool_name": "value",
    "tool_call_id": "value",
    "parameters": {},
    "event_id": 123,
    "expects_response": true
  },
  "type": "client_tool_call"
}`
	assertRoundTrip(t, wire, new(ClientToolCall))
}
