package types

import (
	"testing"
)

func TestRagDocumentIndexResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "model": "e5_mistral_7b_instruct",
  "status": "new",
  "progress_percentage": 9.99,
  "document_model_index_usage": {
    "used_bytes": 123
  }
}`
	assertRoundTrip(t, wire, new(RagDocumentIndexResponseModel))
}

func TestRagDocumentIndexResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "model": "e5_mistral_7b_instruct",
  "status": "new",
  "progress_percentage": 9.99,
  "document_model_index_usage": {
    "used_bytes": 123
  }
}`
	assertRoundTrip(t, wire, new(RagDocumentIndexResponseModel))
}

func TestRagDocumentIndexResponseModelRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "model": "e5_mistral_7b_instruct",
  "status": "new",
  "progress_percentage": 9.99,
  "document_model_index_usage": {
    "used_bytes": 123
  }
}`
	assertRoundTrip(t, wire, new(RagDocumentIndexResponseModel))
}
