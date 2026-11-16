package types

import (
	"testing"
)

func TestClientToolCallClientEventDataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tool_name": "value",
  "tool_call_id": "value",
  "parameters": {
    "key": {}
  },
  "event_id": 123,
  "expects_response": true
}`
	assertRoundTrip(t, wire, new(ClientToolCallClientEventData))
}
