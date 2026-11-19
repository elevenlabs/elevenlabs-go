package types

import (
	"testing"
)

func TestAgentKnowledgeBaseRagChunkResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "document_id": "value",
  "document_name": "value",
  "source_url": "value",
  "chunk_id": "value",
  "text": "value",
  "vector_distance": 9.99,
  "content_format": "html",
  "document_type": "file"
}`
	assertRoundTrip(t, wire, new(AgentKnowledgeBaseRagChunkResponseModel))
}
