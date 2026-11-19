package types

import (
	"testing"
)

func TestRagConfigWorkflowOverrideInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "embedding_model": "e5_mistral_7b_instruct",
  "max_vector_distance": 9.99,
  "max_documents_length": 123,
  "max_retrieved_rag_chunks_count": 123,
  "num_candidates": 123,
  "query_rewrite_prompt_override": "value",
  "knowledge_base_tool_info": {
    "enabled_strategies": [
      "cat"
    ]
  }
}`
	assertRoundTrip(t, wire, new(RagConfigWorkflowOverrideInput))
}
