package types

import (
	"testing"
)

func TestClientToolResultRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "client_tool_result",
  "tool_call_id": "value",
  "result": "value",
  "is_error": true,
  "error_type": "user_rejected"
}`
	assertRoundTrip(t, wire, new(ClientToolResult))
}
