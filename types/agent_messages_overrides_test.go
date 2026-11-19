package types

import (
	"testing"
)

func TestAgentMessagesOverridesRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(AgentMessagesOverrides))
}
