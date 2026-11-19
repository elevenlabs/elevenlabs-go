package types

import (
	"testing"
)

func TestKnowledgeBaseDocumentChunksResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chunks": [
    {
      "id": "value",
      "name": "value",
      "content": "value"
    }
  ],
  "next_cursor": "value"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseDocumentChunksResponseModel))
}
