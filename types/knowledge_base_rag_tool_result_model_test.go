package types

import (
	"testing"
)

func TestKnowledgeBaseRagToolResultModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "knowledge_base_rag_success",
  "status": "success",
  "chunk_count": 123,
  "message": "value",
  "chunks": [
    {
      "chunk_id": "value",
      "document_id": "value",
      "content": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseRagToolResultModel))
}

func TestKnowledgeBaseRagToolResultModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "knowledge_base_rag_success",
  "status": "success",
  "chunk_count": 123,
  "message": "value",
  "chunks": [
    {
      "chunk_id": "value",
      "document_id": "value",
      "content": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseRagToolResultModel))
}
