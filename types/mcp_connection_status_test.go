package types

import (
	"testing"
)

func TestMcpConnectionStatusRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "mcp_connection_status": {
    "integrations": [
      {
        "integration_id": "value",
        "integration_type": "mcp_server",
        "is_connected": true,
        "tool_count": 123
      }
    ]
  },
  "type": "mcp_connection_status"
}`
	assertRoundTrip(t, wire, new(McpConnectionStatus))
}
