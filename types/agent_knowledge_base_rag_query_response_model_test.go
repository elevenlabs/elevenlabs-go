package types

import (
	"testing"
)

func TestAgentKnowledgeBaseRagQueryResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "strategy": "cat",
  "retrieval_query": "value",
  "chunks": [
    {
      "document_id": "value",
      "document_name": "value",
      "source_url": "value",
      "chunk_id": "value",
      "text": "value",
      "vector_distance": 9.99,
      "content_format": "html",
      "document_type": "file"
    }
  ],
  "document": {
    "document_id": "value",
    "document_name": "value",
    "source_url": "value",
    "content": "value",
    "truncated": true,
    "content_format": "html",
    "document_type": "file"
  },
  "listing": {
    "entries": [
      {
        "id": "value",
        "name": "value",
        "is_folder": true
      }
    ],
    "total_count": 123
  }
}`
	assertRoundTrip(t, wire, new(AgentKnowledgeBaseRagQueryResponseModel))
}
