package types

import (
	"testing"
)

func TestKnowledgeBaseToolConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "system_tool_type": "knowledge_base",
  "enabled_strategies": [
    "cat"
  ]
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseToolConfig))
}
