package types

import (
	"testing"
)

func TestMcpToolCallMcpToolCallLoadingRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "service_id": "value",
  "tool_call_id": "value",
  "tool_name": "value",
  "tool_description": "value",
  "parameters": {
    "key": {}
  },
  "timestamp": "value",
  "state": "loading"
}`
	assertRoundTrip(t, wire, new(McpToolCallMcpToolCallLoading))
}
