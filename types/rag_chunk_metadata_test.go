package types

import (
	"testing"
)

func TestRagChunkMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "document_id": "value",
  "chunk_id": "value",
  "vector_distance": 9.99
}`
	assertRoundTrip(t, wire, new(RagChunkMetadata))
}
