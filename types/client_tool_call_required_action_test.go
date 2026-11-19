package types

import (
	"testing"
)

func TestClientToolCallRequiredActionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "client_tool_call",
  "tool_call_id": "value",
  "tool_name": "value",
  "parameters": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(ClientToolCallRequiredAction))
}
