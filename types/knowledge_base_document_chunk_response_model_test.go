package types

import (
	"testing"
)

func TestKnowledgeBaseDocumentChunkResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "content": "value"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseDocumentChunkResponseModel))
}

func TestKnowledgeBaseDocumentChunkResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "content": "value"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseDocumentChunkResponseModel))
}
