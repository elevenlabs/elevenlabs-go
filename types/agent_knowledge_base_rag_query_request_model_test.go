package types

import (
	"testing"
)

func TestAgentKnowledgeBaseRagQueryRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "strategy": "cat",
  "query": "value",
  "documents": [
    "value"
  ],
  "use_agent_defaults": true,
  "max_documents_length": 123,
  "max_retrieved_rag_chunks_count": 20
}`
	assertRoundTrip(t, wire, new(AgentKnowledgeBaseRagQueryRequestModel))
}
