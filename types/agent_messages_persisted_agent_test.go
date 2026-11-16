package types

import (
	"testing"
)

func TestAgentMessagesPersistedAgentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "persisted",
  "agent_id": "value",
  "branch_id": "value",
  "overrides": {
    "prompt": "value",
    "llm": "gpt-4o-mini",
    "language": "value",
    "tool_ids": [
      "value"
    ],
    "native_mcp_server_ids": [
      "value"
    ],
    "knowledge_base_ids": [
      "value"
    ],
    "procedure_ids": [
      "value"
    ],
    "custom_llm_extra_body": {
      "key": {}
    }
  }
}`
	assertRoundTrip(t, wire, new(AgentMessagesPersistedAgent))
}
