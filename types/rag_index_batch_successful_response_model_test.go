package types

import (
	"testing"
)

func TestRagIndexBatchSuccessfulResponseModelRoundTrip(t *testing.T) {
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
	assertRoundTrip(t, wire, new(RagIndexBatchSuccessfulResponseModel))
}
