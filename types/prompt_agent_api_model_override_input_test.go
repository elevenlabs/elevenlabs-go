package types

import (
	"testing"
)

func TestPromptAgentApiModelOverrideInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "prompt": "value",
  "llm": "gpt-4o-mini",
  "reasoning_effort": "none",
  "tool_ids": [
    "value"
  ],
  "tool_ids_strict": true,
  "native_mcp_server_ids": [
    "value"
  ],
  "knowledge_base": [
    {
      "type": "file",
      "name": "value",
      "id": "value",
      "usage_mode": "prompt"
    }
  ]
}`
	assertRoundTrip(t, wire, new(PromptAgentApiModelOverrideInput))
}
