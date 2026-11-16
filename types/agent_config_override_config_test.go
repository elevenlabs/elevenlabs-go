package types

import (
	"testing"
)

func TestAgentConfigOverrideConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "first_message": true,
  "language": true,
  "max_conversation_duration_message": true,
  "prompt": {
    "prompt": true,
    "llm": true,
    "reasoning_effort": true,
    "tool_ids": true,
    "tool_ids_strict": true,
    "native_mcp_server_ids": true,
    "knowledge_base": true
  }
}`
	assertRoundTrip(t, wire, new(AgentConfigOverrideConfig))
}
