package types

import (
	"testing"
)

func TestRagRetrievalInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chunks": [
    {
      "document_id": "value",
      "chunk_id": "value",
      "vector_distance": 9.99
    }
  ],
  "embedding_model": "e5_mistral_7b_instruct",
  "retrieval_query": "value",
  "rag_latency_secs": 9.99,
  "used_chunk_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(RagRetrievalInfo))
}

func TestRagRetrievalInfoRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chunks": [
    {
      "document_id": "value",
      "chunk_id": "value",
      "vector_distance": 9.99
    }
  ],
  "embedding_model": "e5_mistral_7b_instruct",
  "retrieval_query": "value",
  "rag_latency_secs": 9.99,
  "used_chunk_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(RagRetrievalInfo))
}

func TestRagRetrievalInfoRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chunks": [
    {
      "document_id": "value",
      "chunk_id": "value",
      "vector_distance": 9.99
    }
  ],
  "embedding_model": "e5_mistral_7b_instruct",
  "retrieval_query": "value",
  "rag_latency_secs": 9.99,
  "used_chunk_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(RagRetrievalInfo))
}
