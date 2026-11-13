package types

import (
	"testing"
)

func TestKnowledgeBaseRagChunkModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chunk_id": "value",
  "document_id": "value",
  "content": "value"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseRagChunkModel))
}
