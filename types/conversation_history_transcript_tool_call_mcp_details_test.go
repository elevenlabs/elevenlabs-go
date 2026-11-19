package types

import (
	"testing"
)

func TestConversationHistoryTranscriptToolCallMcpDetailsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "mcp",
  "mcp_server_id": "value",
  "mcp_server_name": "value",
  "integration_type": "value",
  "parameters": {
    "key": "value"
  },
  "approval_policy": "value",
  "requires_approval": true,
  "mcp_tool_name": "value",
  "mcp_tool_description": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptToolCallMcpDetails))
}

func TestConversationHistoryTranscriptToolCallMcpDetailsRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "mcp",
  "mcp_server_id": "value",
  "mcp_server_name": "value",
  "integration_type": "value",
  "parameters": {
    "key": "value"
  },
  "approval_policy": "value",
  "requires_approval": true,
  "mcp_tool_name": "value",
  "mcp_tool_description": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptToolCallMcpDetails))
}

func TestConversationHistoryTranscriptToolCallMcpDetailsRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "mcp",
  "mcp_server_id": "value",
  "mcp_server_name": "value",
  "integration_type": "value",
  "parameters": {
    "key": "value"
  },
  "approval_policy": "value",
  "requires_approval": true,
  "mcp_tool_name": "value",
  "mcp_tool_description": "value"
}`
	assertRoundTrip(t, wire, new(ConversationHistoryTranscriptToolCallMcpDetails))
}
