package agents

import (
	"testing"
)

func TestGetOrCreateRagIndexesKnowledgeBaseResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "success",
  "data": {
    "id": "value",
    "model": "e5_mistral_7b_instruct",
    "status": "new",
    "progress_percentage": 9.99,
    "document_model_index_usage": {
      "used_bytes": 123
    }
  }
}`
	assertRoundTrip(t, wire, new(GetOrCreateRagIndexesKnowledgeBaseResponse))
}

func TestBodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "items": [
    {
      "document_id": "value",
      "create_if_missing": true,
      "model": "e5_mistral_7b_instruct"
    }
  ]
}`
	assertRoundTrip(t, wire, new(BodyComputeRagIndexesInBatchV1ConvaiKnowledgeBaseRagIndexPost))
}
