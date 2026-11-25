package types

import (
	"testing"
)

func TestMcpConnectionStatusMcpConnectionStatusRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "integrations": [
    {
      "integration_id": "value",
      "integration_type": "mcp_server",
      "is_connected": true,
      "tool_count": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(McpConnectionStatusMcpConnectionStatus))
}
