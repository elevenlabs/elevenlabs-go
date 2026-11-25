package types

import (
	"testing"
)

func TestMcpConnectionStatusMcpConnectionStatusIntegrationsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "integration_id": "value",
  "integration_type": "mcp_server",
  "is_connected": true,
  "tool_count": 123
}`
	assertRoundTrip(t, wire, new(McpConnectionStatusMcpConnectionStatusIntegrationsItem))
}
