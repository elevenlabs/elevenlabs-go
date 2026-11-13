package types

import (
	"testing"
)

func TestPromptAgentApiModelOverrideConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "prompt": true,
  "llm": true,
  "reasoning_effort": true,
  "tool_ids": true,
  "tool_ids_strict": true,
  "native_mcp_server_ids": true,
  "knowledge_base": true
}`
	assertRoundTrip(t, wire, new(PromptAgentApiModelOverrideConfig))
}
