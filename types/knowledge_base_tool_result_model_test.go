package types

import (
	"testing"
)

func TestKnowledgeBaseToolResultModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "knowledge_base_success",
  "status": "success",
  "chunk_count": 123,
  "message": "value"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseToolResultModel))
}

func TestKnowledgeBaseToolResultModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "result_type": "knowledge_base_success",
  "status": "success",
  "chunk_count": 123,
  "message": "value"
}`
	assertRoundTrip(t, wire, new(KnowledgeBaseToolResultModel))
}
