package types

import (
	"testing"
)

func TestKnowledgeBaseRagToolConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "system_tool_type": "knowledge_base_rag"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseRagToolConfig))
}
